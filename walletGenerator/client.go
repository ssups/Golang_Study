package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// =======================================================

	// generate prviateKey
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("privateKey: ", privateKey)

	// convert to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)

	// encodes to hex string
	encodedPrivateKey := hexutil.Encode(privateKeyBytes)[2:]
	fmt.Println("encoded privateKey: ", encodedPrivateKey)

	// =======================================================

	// extract publicKey from privateKey
	publicKey := privateKey.Public()
	fmt.Println("publickKey: ", publicKey)

	// convert using ECDSA
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ESDSA")
	}
	fmt.Println("publicKeyECDSA: ", publicKeyECDSA)

	// convert to bytes
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("encoded publicKey: ", hexutil.Encode(publicKeyBytes)[4:])

	// convert to hex string
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)

	// =======================================================
}
