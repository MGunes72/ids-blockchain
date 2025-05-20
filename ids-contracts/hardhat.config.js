require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

const { RPC_URL, PRIVATE_KEY } = process.env;

module.exports = {
  solidity: "0.8.20",
  networks: {
    ganache: {
      url: RPC_URL,
      accounts: [PRIVATE_KEY]
    }
  }
};