package main

import (
	"os"
	"github.com/IkezoeMakoto/zcrypt/zcrypt"
)

func main() {
	cli := &zcrypt.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}