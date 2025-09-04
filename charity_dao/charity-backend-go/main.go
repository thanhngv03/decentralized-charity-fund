package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// Replace with the correct path to the generated Go bindings for your CharityVault contract, e.g.:
	"github.com/thanhngv03/decentralized-charity-fund/charity-backend-go/contract"
)

func main() {
	// 1. Kết nối tới Hardhat node
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("Lỗi kết nối Ethereum client:", err)
	}
	fmt.Println("✅ Đã kết nối Hardhat node")

	// 2. Địa chỉ contract (copy từ bước deploy)
	contractAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	vault, err := contract.NewCharityVault(contractAddress, client)
	if err != nil {
		log.Fatal("Không load contract:", err)
	}

	// 3. Private key (copy từ Hardhat node log, bỏ "0x")
	privateKey, err := crypto.HexToECDSA("59c6995e998f97a5a004497e5da...") // thay bằng private key thật
	if err != nil {
		log.Fatal("Không load private key:", err)
	}

	// 4. Tạo transaction auth
	chainID := big.NewInt(31337) // Hardhat chainId
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Không tạo được auth:", err)
	}

	// Gửi kèm 0.1 ETH khi gọi donate()
	auth.Value = big.NewInt(1e17) // 0.1 ETH in wei

	// 5. Gọi hàm donate()
	tx, err := vault.Donate(auth)
	if err != nil {
		log.Fatal("Lỗi donate:", err)
	}
	fmt.Println("💸 Donate thành công, Tx hash:", tx.Hash().Hex())

	// 6. Đọc tổng donate
	total, err := vault.TotalDonated(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		log.Fatal("Lỗi đọc tổng donate:", err)
	}
	fmt.Println("💰 Tổng donate trong contract:", total)
}
