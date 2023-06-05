// docs https://goethereumbook.org/transfer-eth/

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind" // https://pkg.go.dev/github.com/ethereum/go-ethereum/accounts/abi/bind
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"go_ether_tutorial/transactions/util"
)

func main() {
	// check env isValid
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	// key setting
	private_key_account_1 := os.Getenv("PRIVATE_KEY_ACCOUNT_1")
	address_account_2 := os.Getenv("ADDRESS_ACCOUNT_2")

	privateKeyECDSA, err := crypto.HexToECDSA(private_key_account_1)
	if err != nil {
		log.Fatal(err)
	}

	// connect client
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// get chainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainID: ", chainID)

	// convertAddress
	toAddress := common.HexToAddress(address_account_2)
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("erro casting public key to ESDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// get before balance
	beforeBalance, err := client.BalanceAt(context.Background(), toAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	// get nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce: ", nonce)

	// set gas
	gasLimit := uint64(21000)
	gasPrice := gweiToWei(800)
	// get suggested gas price
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// make tx
	tx := types.NewTransaction(nonce, toAddress, etherToWei(1000), gasLimit, gasPrice, nil)
	// nonce, toAddress, value, gasLimit, gasPrice, data

	// sign tx
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKeyECDSA)
	if err != nil {
		log.Fatal(err)
	}

	// send singed tx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

	// get after balance
	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx's block number: ", receipt.BlockNumber)
	afterBalance, err := client.BalanceAt(context.Background(), toAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("beforeBalance: ", beforeBalance)
	fmt.Println("afterBalance: ", afterBalance)

}

func etherToWei(eth int64) *big.Int {
	return util.ToWei(eth, 18)
}

func gweiToWei(gwei int64) *big.Int {
	return util.ToWei(gwei, 9)
}
