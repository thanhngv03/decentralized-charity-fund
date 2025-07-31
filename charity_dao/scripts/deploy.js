const hre = require("hardhat");

async function main(){
    const CharityVault = await hre.ethers.getContractFactory("CharityVault");
    const vault = await CharityVault.deploy();
    // await vault.deployed();

    console.log(`CharityVault deployed to: ${vault.target || vault.address}`);
}

main().catch((error) =>{
    console.error(error);
    process.exitCode = 1;
});

