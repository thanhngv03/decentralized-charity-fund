// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.20;

/// @title CharityVault 
/// @notice Hợp đồng cho phép nhận donation ETH và rút donation ETH 

contract CharityVault {
    address public owner;
    uint256 public totalDonated;

    mapping(address => uint256) public donations; 

    event Donated(address indexed donor, uint256 amount); /// hàm donate
    event Withdraw(address indexed to, uint256 amount); /// sự kiện rút tiền

    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    constructor(){
        owner = msg.sender;
    }

    ///@notice Nhận donation ETH   
    function donate() external payable {
        require(msg.value > 0,"Must send ETH");
        donations[msg.sender] += msg.value;
        totalDonated += msg.value;

        emit Donated(msg.sender, msg.value);
    }

    /// @notice Rút tiền về ví của owner 
    /// @param amount Số ETH muốn rút (tính theo wei)
    function withdraw(uint256 amount) external onlyOwner{
        require(amount <= address(this).balance, "Insufficient balance");

        //Gửi ETH về ví của admin 
        (bool sent, )= payable(owner).call{value: amount}("");
        require(sent, "Withdraw failed");

        emit Withdraw(owner, amount);
    }

    /// @notice Xem số tiền một địa chỉ đã donate 
    function getDonationOf(address donor) external view returns (uint256) {
        return donations[donor];
    }

    /// @dev Tự động xử lý nếu nhận ETH trực tiếp
    receive() external payable{
        // Gọi lại donate() để ghi nhận event và cập nhật dữ liệu
        this.donate();
    }
}