package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bithyve/bithyve-wrapper/electrs"

	//"strings"

	erpc "github.com/Varunram/essentials/rpc"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Mainnet bool `short:"m" description:"Connect to mainnet"`
	Test    bool `short:"t" description:"Use for testing"`
}

func startHandlers() {
	MultiData()
	MultiBalTxs()
	MultiUtxoTxs()
	MultiUtxos()
	MultiBalances()
	MultiTxs()

	erpc.SetupPingHandler()
	GetFees()
	PostTx()
	RelayTxid()
	RelayGetRequest()
}

func main() {
	startHandlers()

	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if opts.Mainnet {
		log.Println("connecting to electrs mainnet")
		electrs.SetMainnet()
	}

	if opts.Test {
		err = http.ListenAndServe("localhost:8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	} else {
		err = http.ListenAndServeTLS("localhost:445", "ssl/server.crt", "ssl/server.key", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}
