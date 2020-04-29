package main

import blockchain "block"
import command "cli"
//import "fmt"
//import "strconv"

func main() {
	/*
	bc := blockchain.NewBlockchain()

	// 블록 생성
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		// 블록 내용 출력
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		// 블록 검증
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
	*/

	bc := blockchain.NewBlockchain()
	defer bc.Db.Close()

	cli := command.CLI{bc}
	cli.Run()
}
