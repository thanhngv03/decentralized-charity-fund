const hre = require ("hardhat");

async function main(){
    const [deployer] = await hre.eithers.getSigners();

    const CharityVault = await hre.ethers.getContractFactory("CharityVault");
    const vault = await CharityVault.deploy(deployer.address);

    await vault.deployed();

    console.log(`CharityVault deployed to: ${vault.address}`);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});

