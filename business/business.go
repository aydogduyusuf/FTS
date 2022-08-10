package business

import (
	"fmt"
	"FTS/blockchain"
)

func main() {
	//setup client
	avalancheFuji := "https://api.avax-test.network/ext/bc/C/rpc"
	client := blockchain.InitNetwork(avalancheFuji)

	fmt.Println("client", client)
}
