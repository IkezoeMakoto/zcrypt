package main

import (
	"os"
	"github.com/IkezoeMakoto/etc-trial/golang/app/service/zcrypt"
)

func main() {
	cli := &zcrypt.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}