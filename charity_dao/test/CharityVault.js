const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("CharityVault", function () {
    let CharityVault, vault, owner, addr1, addr2;

    beforeEach(async function () {
        [owner, addr1, addr2] = await ethers.getSigners();
        
        CharityVault = await ethers.getContractFactory("CharityVault");
        vault = await CharityVault.deploy(owner.address);
        await vault.waitForDeployment();
    });

    it ("Owner should be deployer", async function () {
        expect(await vault.owner()).to.equal(owner.address);    
    });

    it ("Should accept donations", async function () {
        await vault.connect(addr1).donate({ value: ethers.parseEther("1") });

        const donation = await vault.getDonationOf(addr1.address);
        expect(donation).to.equal(ethers.parseEther("1"));

        const total = await vault.totalDonated();
        expect(total).to.equal(ethers.parseEther("1"));
    });
    
    it ("Should only allow owner to withdraw", async function () {
        await vault.connect(addr1).donate({ value: ethers.parseEther("2") });

        // non-owner (addr1) trying to withdraw -> revert
        await expect(
            vault.connect(addr1).withdraw(address, ethers.parseEther("1"))
        ).to.be.revertedWith("Not Owner");

        // owner can withdraw
        await expect(vault.connect(owner).withdraw(ethers.parseEther("1")))
            .to.emit(vault, "Withdraw")
            .withArgs(owner.address, ethers.parseEther("1"));
    });

    it ("Should fail if withdraw > balance", async function () {
        await vault.connect(addr1).donate({ value: ethers.parseEther("1") });

        await expect(
            vault.connect(owner).withdraw(ethers.parseEther("5"))
        ).to.be.revertedWith("Insufficient balance");
    });

});