package multicoinwallet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
)

var btcRpcClient *rpcclient.Client

func BTCInit() {
	var err error

	fmt.Printf("btc init start:\n")

	rpcCertificate, err := ioutil.ReadFile("rpc.cert")
	if err != nil {
		log.Printf("error read cert file: %v", err)
		return
	}
	fmt.Printf("btc init:get rpc.cert ok\n")
	// create new client instance
	btcRpcClient, err = rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		Certificates: rpcCertificate,
		Host:         "192.168.19.130:8332",
		User:         "root",
		Pass:         "root",
	}, nil)
	if err != nil {
		log.Printf("error creating new btc client: %v", err)
		return
	}
	fmt.Printf("btc init:get create client instance ok\n")
	fmt.Printf("btc init end\n")
}

var bchRpcClient *rpcclient.Client

func BCHInit() {
	var err error

	fmt.Printf("bch init start:\n")
	//	rpcCertificate, err := ioutil.ReadFile("/home/zhutian/.btcwallet/rpc.cert")
	//	if err != nil {
	//		log.Printf("error read cert file: %v", err)
	//		return
	//	}

	// create new client instance
	bchRpcClient, err = rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "127.0.0.1:19001",
		User:         "admin1",
		Pass:         "123",
	}, nil)
	if err != nil {
		log.Printf("error creating new btc client: %v", err)
		return
	}
	fmt.Printf("bch init:get create client instance ok\n")
	fmt.Printf("bch init end\n")
}

func GetBtcClient(coin *CoinType) *rpcclient.Client {
	var client *rpcclient.Client

	if coin.Bip44Index == BITCOIN {
                
		client = btcRpcClient
                fmt.Printf("The btc client: %v\n", client)

	} else if coin.Bip44Index == BITCOIN_CASH {
               
		client = bchRpcClient
                fmt.Printf("The bch client: %v\n", client)

	} else {

		client = nil
                fmt.Printf("Unknow coin type %d\n", coin.Bip44Index)
	}

	return client
}

func BTCGetBalance(coin *CoinType) float64 {
	client := GetBtcClient(coin)

	amount, err := client.GetBalance("")
	if err != nil {
		fmt.Printf("error get balance btc client: %v\n", err)
		return 0.0
	}

	fmt.Printf("get balance: %d\n", amount)

	return float64(amount) * 1e-8
}

func BTCGetNewAddress(coin *CoinType) string {
	client := GetBtcClient(coin)
       
        if client == nil {
             log.Printf("Get new address failed: client instance is nil")
             return ""
         }

	address, err := client.GetNewAddress("default")
	if err != nil {
		log.Printf("error get btc new address client: %v", err)
		return ""
	}

	log.Printf("get btc new address: %v", address)

	return address.EncodeAddress()
}

func BTCListTransactions(coin *CoinType, account string, count int, fromStart int) string {
	client := GetBtcClient(coin)

	trans, err := client.ListTransactionsCountFrom(account, count, fromStart)
	if err != nil {
		log.Printf("error list transactions: %v", err)
		return ""
	}

	transData, err := json.Marshal(&trans)
	if err != nil {
		log.Printf("error data of list transactions: %v", err)
		return ""
	}

	log.Printf("list transactions: %s", string(transData))

	return string(transData)
}

func BTCListReceivedByAddress(coin *CoinType) string {
	client := GetBtcClient(coin)

	payments, err := client.ListReceivedByAddress()
	if err != nil {
		log.Printf("error list transactions: %v", err)
		return ""
	}

	paymentsData, err := json.Marshal(&payments)
	if err != nil {
		log.Printf("error data of list recv by address: %v", err)
		return ""
	}

	log.Printf("list reveive by address: %s", string(paymentsData))

	return string(paymentsData)
}

func BTCGetTransaction(coin *CoinType, txId string) string {
	client := GetBtcClient(coin)

	txHash, err := chainhash.NewHashFromStr(txId)
	if err != nil {
		return ""
	}

	trans, err := client.GetTransaction(txHash)
	if err != nil {
		log.Printf("error get transactions: %v", err)
		return ""
	}

	transData, err := json.Marshal(&trans)
	if err != nil {
		log.Printf("error data of get transaction: %v", err)
		return ""
	}

	log.Printf("get transaction: %s", string(transData))

	return string(transData)
}

func BTCSendToAddress(coin *CoinType, address string, amount float64) string {
	client := GetBtcClient(coin)

	txId, err := client.SendToAddressEx(address, amount)
	if err != nil {
		log.Printf("error send to address: %v", err)
		return ""
	}

	log.Printf("Send to address ok, transaction hash:%s", txId)

	return txId
}

func BTCWalletLock(coin *CoinType) error {
	client := GetBtcClient(coin)

	err := client.WalletLock()
	if err != nil {
		log.Printf("error lock wallet: %v", err)
		return err
	}

	return nil
}

func BTCWalletPassphrase(coin *CoinType) error {
	client := GetBtcClient(coin)

	err := client.WalletPassphrase("123456", 60)
	if err != nil {
		log.Printf("error wallet passphrase: %v", err)
		return err
	}

	return nil
}
