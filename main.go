package main

import (
	"Proj/HomeWork/golang_package/blockchain"
	"fmt"
)

func main() {

	myBlockchain := blockchain.CreateBlockchain(3)
	myBlockchain.AddBlock("Alice", "Bob", 5)
	fmt.Println(myBlockchain.IsValid())

}
