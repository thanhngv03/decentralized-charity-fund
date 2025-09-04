// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

/**
 * @title CharityVault
 * @dev Hợp đồng thông minh cho phép quyên góp ETH vào quỹ từ thiện minh bạch.
 * - Ai cũng có thể donate.
 * - Chỉ admin/DAO mới được quyền rút.
 * - Mọi giao dịch đều phát event công khai.
 */
contract CharityVault is AccessControl, ReentrancyGuard {
    /// @notice Role cho admin (DAO/multi-sig)
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /// @notice mapping địa chỉ → tổng số tiền đã donate
    mapping(address => uint256) public donations;

    /// @notice tổng số tiền donate
    uint256 public totalDonated;

    /// @notice event khi có donate
    event DonationReceived(address indexed donor, uint256 amount);

    /// @notice event khi có rút tiền
    event FundsWithdrawn(address indexed to, uint256 amount);

    /**
     * @dev constructor: deployer mặc định là admin
     */
    constructor(address admin) {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(ADMIN_ROLE, admin);
    }

    /**
     * @notice Donate ETH vào quỹ
     * @dev anyone có thể donate
     */
    function donate() external payable nonReentrant {
        require(msg.value > 0, "Phai gui ETH > 0");
        donations[msg.sender] += msg.value;
        totalDonated += msg.value;
        emit DonationReceived(msg.sender, msg.value);
    }

    /**
     * @notice Admin rút tiền từ quỹ
     * @dev chỉ ADMIN_ROLE mới gọi được
     * @param to địa chỉ nhận tiền
     * @param amount số wei rút ra
     */
    function withdraw(address payable to, uint256 amount)
        external
        onlyRole(ADMIN_ROLE)
        nonReentrant
    {
        require(amount > 0, "So tien phai > 0");
        require(address(this).balance >= amount, "Khong du so du");

        (bool success, ) = to.call{value: amount}("");
        require(success, "Chuyen tien that bai");

        emit FundsWithdrawn(to, amount);
    }

    /**
     * @notice Xem số tiền một donor đã donate
     */
    function getDonationOf(address donor) external view returns (uint256) {
        return donations[donor];
    }

    /**
     * @notice Nhận ETH trực tiếp
     */
    receive() external payable {
        donate();
    }
}
