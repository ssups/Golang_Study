package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("privateKey: ", privateKey)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("encoded privateKey: ", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ESDSA")
	}
	fmt.Println("publicKey: ", publicKey)
	fmt.Println("publicKeyESDSA: ", publicKeyECDSA)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("encoded publicKey: ", hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)

}
