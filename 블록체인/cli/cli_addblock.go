package cli

import "fmt"

func (cli *CLI) addBlock(data string) {
	cli.Bc.AddBlock(data)
	fmt.Println("Success!")
}
