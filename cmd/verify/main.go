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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	port, _ := strconv.Atoi(dbPort)

	// Connect to Ethereum
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	contract, err := snortlogger.NewSnortlogger(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, port, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("PostgreSQL connection failed: %v", err)
	}
	defer db.Close()

	// Input timestamp
	fmt.Print("Enter timestamp to verify (UNIX): ")
	var tsInput uint64
	fmt.Scan(&tsInput)

	if err := verifyAlert(db, contract, tsInput); err != nil {
		log.Println(err)
	}
}

func verifyAlert(db *sql.DB, contract *snortlogger.Snortlogger, timestamp uint64) error {
	ctx := context.Background()

	var alertText string
	err := db.QueryRow(`
		SELECT alert_text FROM snort_alerts
		WHERE EXTRACT(EPOCH FROM timestamp)::bigint = $1
	`, timestamp).Scan(&alertText)
	if err != nil {
		return fmt.Errorf("❌ PostgreSQL alert not found: %w", err)
	}

	localHash := sha256.Sum256([]byte(alertText))
	localHashHex := hex.EncodeToString(localHash[:])

	sender, alertHash, _, err := contract.GetAlert(&bind.CallOpts{Context: ctx}, big.NewInt(int64(timestamp)))
	if err != nil {
		return fmt.Errorf("❌ Failed to read from contract: %w", err)
	}

	fmt.Printf("👨‍💼 Sender: %s\n", sender.Hex())
	fmt.Printf("🔍 Verifying timestamp: %d\n", timestamp)
	fmt.Printf("📝 PostgreSQL hash: %s\n", localHashHex)
	fmt.Printf("🧾 On-chain hash:     %s\n", alertHash)

	if localHashHex == alertHash {
		fmt.Println("✅ VERIFIED: Hash matches")
	} else {
		fmt.Println("❌ MISMATCH: Alert may have been tampered with!")
	}

	return nil
}
