const { ethers } = require("hardhat");
const fs = require("fs");
const path = require ("path");

async function main(){
    const [deployer] = await ethers.getSigners();

    console.log("Deploying contracts with account:", deployer.address);

    const CharityVault = await ethers.getContractFactory("CharityVault");

    //deploy và truyền admin (dùng deployer Làm admin)
    const vault = await CharityVault.deploy(deployer.address);  

    //với ethers v6
    await vault.waitForDeployment();

    const contractAddress = await vault.getAddress();
    console.log("CharityVault deployed to: ", contractAddress);
}
    
    //Ghi địa chỉ contract vào file JSON
    const deploymentsDir = path.join (__dirname, "..", "deployments");
    if (!fs.existsSync(deploymentsDir)) {
        fs.mkdirSync(deploymentsDir);
    }

    const deploymentInfo = {
        CharityVault: contractAddress,
        network: "localhost",
        deployer: deployer.address,
    };

    fs.writeFileSync(
        path.join(deploymentsDir, "deployed-address.json"),
        JSON.stringify(deploymentInfo, null, 2),
        "utf-8"
    );

    console.log("Deployment info saved to deployments/deployed-address.json");
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});

