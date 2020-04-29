package cli

import (
	"block"
	"fmt"
	"flag"
	"os"
)

type CLI struct {
	Bc *block.Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println(" createwallet - Generates a new key-pair and saves it into the wallet file")
	fmt.Println(" getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println(" listaddresses - Lists all addresses from the wallet file")
	fmt.Println(" printchain - Print all the blocks of the blockchain")
	fmt.Println(" reindexutxo - Rebuilds the UTXO set")
	fmt.Println(" send -from FROM -to TO -amount AMOUNT -mine - Send AMOUNT of coins from FROM address to TO. Mine on the same node, when -mine is set.")
	fmt.Println(" startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining")
}

func (cli *CLI) validateArgs() { 
	if len(os.Args) < 2 { 
		cli.printUsage() 
		os.Exit(1) 
	} 
} 

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		addBlockCmd.Parse(os.Args[2:])
	case "printchain":
		printChainCmd.Parse(os.Args[2:])
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}