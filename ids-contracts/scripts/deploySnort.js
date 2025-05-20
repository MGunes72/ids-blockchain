const hre = require("hardhat");

async function main() {
  await hre.run('compile');

  const SnortLogger = await hre.ethers.getContractFactory("SnortLogger");
  const snortLogger = await SnortLogger.deploy();

  await snortLogger.waitForDeployment();

  console.log("SnortLoggger deployed to:", snortLogger.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});