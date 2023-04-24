// docs https://goethereumbook.org/smart-contract-read-erc20/

package main

import (
	"context"
	"fmt"
	"log"

	// "math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	token "go_ether_tutorial/interact_contract/contracts"
)

const TOKEN_ADDRESS = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
const ACCOUNT0 = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
const ACCOUNT1 = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// account's Eth balance
	account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ToDecimal(balance, 18))

	// connect to token contract
	token_address := common.HexToAddress(TOKEN_ADDRESS)
	instance, err := token.NewToken(token_address, client)
	if err != nil {
		log.Fatal(err)
	}

	// token name
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)

	// token symbol
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("symbol: %s\n", symbol)

	// token decimal
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decimals: %v\n", decimals)

	// account's token balance
	account0 := common.HexToAddress(ACCOUNT0)
	balOfAcc0, err := instance.BalanceOf(&bind.CallOpts{}, account0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balOfAcc0)
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}
