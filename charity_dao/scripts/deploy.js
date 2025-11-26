const hre = require("hardhat");
const fs = require("fs");
const path = require("path");

async function main() {
  const [deployer] = await hre.ethers.getSigners();

  console.log("Deploying contracts with account:", deployer.address);

  // Lấy contract factory (ethers v6)
  const Vault = await hre.ethers.getContractFactory("CharityVault");

  // Deploy contract, truyền địa chỉ admin vào constructor
  const vault = await Vault.deploy(deployer.address);

  // Ethers v6: chờ deploy
  await vault.waitForDeployment();

  const contractAddress = await vault.getAddress();
  console.log("CharityVault deployed to:", contractAddress);

  // -----------------------------
  // Lưu vào JSON để backend Go sử dụng
  // -----------------------------

  const outputDir = path.join(__dirname, "../deployments");
  if (!fs.existsSync(outputDir)) {
    fs.mkdirSync(outputDir);
  }

  const data = {
    network: "localhost",
    CharityVault: contractAddress,
    deployer: deployer.address,
  };

  fs.writeFileSync(
    path.join(outputDir, "deployed-address.json"),
    JSON.stringify(data, null, 2)
  );

  console.log("Saved deployment info to deployments/deployed-address.json");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
