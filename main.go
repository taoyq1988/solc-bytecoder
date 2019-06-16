package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"github.com/tomtao/solc-bytecoder/common"
	"github.com/tomtao/solc-bytecoder/crypto"
)

var (
	version = "0.0.1"
)

type argOp struct {
	cli.Helper
	Version bool   `cli:"v,version" usage:"get tool version"`
	BinPath string `cli:"bin" usage:"Solidity bin file path"`
	Sha3    bool   `cli:"s,sha3" usage:"get sha3 of the string, result is hex"`
	Arg     string `cli:"arg" usage:"sha3 string arg"`
}

func main() {
	cli.Run(new(argOp), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argOp)

		switch {
		case argv.Version:
			fmt.Println(version)
		case argv.Sha3:
			data, err := crypto.Sha3([]byte(argv.Arg))
			if err != nil {
				return err
			}
			hex := common.Bytes2Hex(data[:4])
			fmt.Println(hex)
		default:
			fmt.Println(argv.BinPath)
		}
		return nil
	})
}
