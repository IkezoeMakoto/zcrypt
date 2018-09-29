package zcrypt

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const Version string = "v0.1.0"

const (
	ExitCodeOK = 0
	ExitCodeNG = 1

	Name = "zcrypt"
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var (
		enc     bool
		dec     bool
		in      string
		out     string
		key     string
		version bool
	)

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.SetOutput(cli.ErrStream)
	flags.Usage = cli.printHelp
	flags.BoolVar(&enc, "enc", false, "Encrypt")
	flags.BoolVar(&dec, "dec", false, "Decrypt")
	flags.StringVar(&in, "in", "", "Input file path")
	flags.StringVar(&out, "out", "", "Output file path")
	flags.StringVar(&key, "key", "", "Encryption key string")
	flags.BoolVar(&version, "version", false, "Print version")
	flags.BoolVar(&version, "v", false, "Print version")

	if len(args) < 2 {
		cli.printHelp()
		return ExitCodeNG
	}

	mode := args[1]
	err := flags.Parse(args[2:])
	cli.chkErr(err)

	if version {
		fmt.Fprintln(cli.OutStream, Version)
		return ExitCodeOK
	}

	var Zcrypt Zcrypt

	keyStr, err := ioutil.ReadFile(key)
	cli.chkErr(err)

	switch mode {
	case "enc":
		Zcrypt, err = NewEncryptor(string(keyStr))
	case "dec":
		Zcrypt, err = NewDecryptor(string(keyStr))
	default:
		fmt.Fprintln(cli.ErrStream, "Invalid command")
		return ExitCodeNG
	}
	cli.chkErr(err)

	input, err := ioutil.ReadFile(in)
	cli.chkErr(err)

	result, err := Zcrypt.Exec(input)
	cli.chkErr(err)

	err = ioutil.WriteFile(out, result, 0644)
	cli.chkErr(err)

	return ExitCodeOK
}

func (cli *CLI) chkErr(err error) {
	if err != nil {
		fmt.Fprintln(cli.ErrStream, err)
		os.Exit(ExitCodeNG)
	}
}

func (cli *CLI) printHelp() {
	fmt.Fprintln(cli.ErrStream, usage)
}

var usage = `Usage: zcrypt (enc|dec) -in input_path -out output_path -key encrypt_key_path`