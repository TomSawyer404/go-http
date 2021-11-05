package main

import (
	"go_http/operator"
	"go_http/parser"
	"log"
	"os"
)

const (
	NO_DATA = iota
	JSON_DATA
	FORM_DATA
)

func main() {
	argv := os.Args
	if len(argv) <= 1 {
		log.Fatalln("Usage: go-http <ip>:<port>")
	} else if len(argv) == 2 {
		operator.DoGET(argv[1], nil)
	} else {
		cli, err := parser.ParseAgrv(argv[2:])
		if err != nil {
			log.Fatalln(`parser.ParseAgrv() ->`, err)
		}

		if 0 == len(cli.BodyTable) {
			operator.DoGET(argv[1], cli.HeaderTable)
		} else {
			operator.DoPOST(argv[1], cli)
		}
	}
}
