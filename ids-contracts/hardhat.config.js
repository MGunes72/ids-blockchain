require("@nomicfoundation/hardhat-toolbox");

module.exports = {
  solidity: "0.8.20",
  networks: {
    ganache: {
      url: "http://127.0.0.1:7545",
      accounts: [
        "0x45ed158fb3960245e59d81a963f61ae33bdd0c73796e7e6cef0a2daada6ef31c"
      ]
    }
  }
};