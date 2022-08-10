package main

import (
	"FTS/blockchain"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	avalancheFuji := "https://api.avax-test.network/ext/bc/C/rpc"
	client := blockchain.InitNetwork(avalancheFuji)

	address, privateKey := blockchain.ImportWallet("ec9b67551089ca3416f3355d95be190ab522d9cf12d2ff3aaa322b3909cd1405")

	txOps := blockchain.NewTransactOpts(context.Background(), privateKey)
	blockchain.SetTransactOpts(client, address, context.Background(), txOps, big.NewInt(0), uint64(100000000), big.NewInt(0))

	contractAddress := common.HexToAddress("0xC798460bfB5145E20DBC7dE7c28e5cCba304D7D9")
	instance, err := blockchain.NewBlockchain(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Contract")
	fmt.Println("Contract address: ", contractAddress.Hex())

	owner, err := instance.Owner(blockchain.NewCallOpts(true, address, nil, nil))
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Wallet Owner: ", owner.Hex())

	value, _ := client.BalanceAt(context.Background(), address, nil)
    fmt.Println("wallet: ", value)
}