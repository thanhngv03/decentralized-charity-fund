const { ethers } = require("hardhat");

async function main(){
    const [deployer] = await ethers.getSigners();

    console.log("Deploying contracts with account:", deployer.address);

    const CharityVault = await ethers.getContractFactory("CharityVault");

    //deploy và truyền admin (dùng deployer Làm admin)
    const vault = await CharityVault.deploy(deployer.address);  

    //với ethers v6
    await vault.waitForDeployed();

    console.log("CharityVault deployed to: ", await vault.getAddress());
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});

