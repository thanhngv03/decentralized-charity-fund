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

  it("Deployer should have ADMIN_ROLE", async function () {
    const ADMIN_ROLE = ethers.keccak256(ethers.toUtf8Bytes("ADMIN_ROLE"));
    expect(await vault.hasRole(ADMIN_ROLE, owner.address)).to.equal(true);
  });

  it("Should accept donations", async function () {
    await vault.connect(addr1).donate({ value: ethers.parseEther("1") });

    const donation = await vault.getDonationOf(addr1.address);
    expect(donation).to.equal(ethers.parseEther("1"));

    const total = await vault.totalDonated();
    expect(total).to.equal(ethers.parseEther("1"));
  });

  it("Should only allow ADMIN_ROLE to withdraw", async function () {
    await vault.connect(addr1).donate({ value: ethers.parseEther("2") });

    // non-admin (addr1) trying to withdraw -> revert
    await expect(
      vault.connect(addr1).withdraw(addr1.address, ethers.parseEther("1"))
    ).to.be.reverted;

    // admin (owner) can withdraw
    await expect(
      vault.connect(owner).withdraw(owner.address, ethers.parseEther("1"))
    )
      .to.emit(vault, "FundsWithdrawn")
      .withArgs(owner.address, ethers.parseEther("1"));
  });

  it("Should fail if withdraw > balance", async function () {
    await vault.connect(addr1).donate({ value: ethers.parseEther("1") });

    await expect(
      vault
        .connect(owner)
        .withdraw(owner.address, ethers.parseEther("5"))
    ).to.be.revertedWith("Khong du so du");
  });
});
