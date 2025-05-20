package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/MGunes72/ids-blockchain/snortlogger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rpcURL := os.Getenv("RPC_URL")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	privateKeyHex := os.Getenv("PRIVATE_KEY")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	port, _ := strconv.Atoi(dbPort)

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	contract, err := snortlogger.NewSnortlogger(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, port, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("PostgreSQL connection failed: %v", err)
	}
	defer db.Close()

	lastSeenID := 0
	fmt.Println("✅ Monitoring database for new Snort alerts...")

	for {
		rows, err := db.Query(`
			SELECT id, timestamp, alert_text FROM snort_alerts
			WHERE id > $1 ORDER BY id ASC
		`, lastSeenID)
		if err != nil {
			log.Printf("Failed to query alerts: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for rows.Next() {
			var id int
			var ts time.Time
			var alert string
			if err := rows.Scan(&id, &ts, &alert); err != nil {
				log.Printf("Scan failed: %v", err)
				continue
			}

			lastSeenID = id

			data := fmt.Sprintf("%d|%s", ts.Unix(), alert)
			hash := sha256.Sum256([]byte(data))
			hashStr := hex.EncodeToString(hash[:])
			fmt.Printf("[→] Sending alert ID %d: %s\n", id, hashStr)

			auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
			if err != nil {
				log.Fatal(err)
			}
			nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
			gasPrice, _ := client.SuggestGasPrice(context.Background())

			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)
			auth.GasLimit = uint64(300000)
			auth.GasPrice = gasPrice

			tx, err := contract.LogAlert(auth, big.NewInt(int64(id)), hashStr)
			if err != nil {
				log.Printf("❌ Failed to send alert %d: %v", id, err)
			} else {
				log.Printf("✅ Alert %d written to blockchain. TX: %s", id, tx.Hash().Hex())
			}

			time.Sleep(3 * time.Second)
		}

		rows.Close()
		time.Sleep(5 * time.Second)
	}
}
