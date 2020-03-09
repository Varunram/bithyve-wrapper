package electrs

import (
	"encoding/json"
	"log"

	"github.com/bithyve/bithyve-wrapper/format"

	erpc "github.com/Varunram/essentials/rpc"
	"github.com/Varunram/essentials/utils"
)

// ElectrsURL is the URL of a running electrs instance
var ElectrsURL = "http://testapi.bithyve.com"

// SetMainnet is the URL of a running mainnet electrs instance
func SetMainnet() {
	ElectrsURL = "http://api.bithyve.com"
}

// CurrentBlockHeight gets the current block height from the blockchain
func CurrentBlockHeight() (float64, error) {
	body := ElectrsURL + "/blocks/tip/height"
	data, err := erpc.GetRequest(body)
	if err != nil {
		log.Println("did not get response", err)
		return -1, err
	}

	return utils.ToFloat(data)
}

// GetBalanceCount gets the total incoming balance
func GetBalanceCount(addr string) (float64, float64) {
	body := ElectrsURL + "/address/" + addr
	data, err := erpc.GetRequest(body)
	if err != nil {
		log.Println("did not get response", err)
		return 0, 0
	}
	// now data is in byte, we need the other structure now
	var x format.Balance
	err = json.Unmarshal(data, &x)
	if err != nil {
		log.Println("did not unmarshal json", err)
		return 0, 0
	}

	return x.ChainStats.FundedTxoCount, x.MempoolStats.FundedTxoCount
}

// GetBalanceAddress gets the net balance of an address
func GetBalanceAddress(addr string) (float64, float64) {
	body := ElectrsURL + "/address/" + addr
	data, err := erpc.GetRequest(body)
	if err != nil {
		log.Println("did not get response", err)
		return 0, 0
	}
	// now data is in byte, we need the other structure now
	var x format.Balance
	err = json.Unmarshal(data, &x)
	if err != nil {
		log.Println("did not unmarshal json", err)
		return 0, 0
	}

	return x.ChainStats.FundedTxoSum - x.ChainStats.SpentTxoSum,
		x.MempoolStats.FundedTxoSum - x.MempoolStats.SpentTxoSum
}

// GetTxsAddress gets the transactions associated with a given address
func GetTxsAddress(addr string) ([]format.Tx, error) {
	var x []format.Tx
	body := ElectrsURL + "/address/" + addr + "/txs"
	data, err := erpc.GetRequest(body)
	if err != nil {
		log.Println("did not get response", err)
		return x, err
	}
	// now data is in byte, we need the other structure now
	err = json.Unmarshal(data, &x)
	if err != nil {
		log.Println("did not unmarshal json", err)
		return x, err
	}

	return x, nil
}

// GetUtxosAddress gets the utxos associated with a given address
func GetUtxosAddress(addr string) ([]format.Utxo, error) {
	var x []format.Utxo
	body := ElectrsURL + "/address/" + addr + "/utxo"
	data, err := erpc.GetRequest(body)
	if err != nil {
		log.Println("did not get response", err)
		return nil, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		log.Println("did not unmarshal json", err)
		return nil, err
	}

	for i := range x {
		x[i].Address = addr
	}
	return x, nil
}
