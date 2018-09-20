package main

import (
	"fmt"

	"github.com/multicoinwallet"
)

func main() {
	BCHRpc()
	//BitcoinRpc()
}

func BitcoinRpc() {
	var coinType multicoinwallet.CoinType

	coinType.Bip44Index = multicoinwallet.BITCOIN

	fmt.Printf("Test BTC start:\n")
	multicoinwallet.BTCInit()
	fmt.Printf("Test get balance:\n")
	amount := multicoinwallet.GetBalance(&coinType)
	fmt.Printf("Balance: %f\n", amount)

	/*	fmt.Printf("Test get new address:\n")
		address := multicoinwallet.GetNewAddress(&coinType, "bitcoin")
		fmt.Printf("Get new address: %s\n", address)
	*/
	fmt.Printf("Test list transactions:\n")
	transactions := multicoinwallet.ListTransactions(&coinType, "default", 10, 0)
	fmt.Printf("List transactions:\n%s\n", transactions)

	fmt.Printf("Test list receive by address:\n")
	payments := multicoinwallet.ListReceivedByAddress(&coinType)
	fmt.Printf("List receive by address:\n%s\n", payments)

	fmt.Printf("Test get transaction:\n")
	transaction := multicoinwallet.GetTransaction(&coinType, "dd2701f95a6e3a59868b10465256e839ef130d0a6728fbb92660e4061867ddb2")
	fmt.Printf("Get transaction:\n%s\n", transaction)

	fmt.Printf("Test send to address:\n")
	txId := multicoinwallet.SendToAddress(&coinType, "16kGvbDUATfGfqqNwLxWPzncVWHNJyScjc", 1.0)
	fmt.Printf("Send to address: %s\n", txId)

	fmt.Printf("Test wallet lock:\n")
	lock := multicoinwallet.WalletLock(&coinType)
	if lock != nil {
		fmt.Printf("Wallet lock fail: %v\n", lock)
	} else {
		fmt.Printf("Wallet lock:ok\n")
	}

	pass := multicoinwallet.WalletPassphrase(&coinType)
	if pass != nil {
		fmt.Printf("Wallet passphrase fail: %v\n", lock)
	} else {
		fmt.Printf("Wallet passphrase: ok\n")
	}

	fmt.Printf("Test BTC end\n")
}

func BCHRpc() {
	var coinType multicoinwallet.CoinType

	coinType.Bip44Index = multicoinwallet.BITCOIN_CASH

	fmt.Printf("Test BCH start:\n")
	multicoinwallet.BCHInit()
	fmt.Printf("Test get balance:\n")
	amount := multicoinwallet.GetBalance(&coinType)
	fmt.Printf("Balance: %f\n", amount)

	/*	fmt.Printf("Test get new address:\n")
		address := multicoinwallet.GetNewAddress(&coinType, "bitcoin")
		fmt.Printf("Get new address: %s\n", address)
	*/
	fmt.Printf("Test list transactions:\n")
	transactions := multicoinwallet.ListTransactions(&coinType, "default", 10, 0)
	fmt.Printf("List transactions:\n%s\n", transactions)

	fmt.Printf("Test list receive by address:\n")
	payments := multicoinwallet.ListReceivedByAddress(&coinType)
	fmt.Printf("List receive by address:\n%s\n", payments)

	fmt.Printf("Test get transaction:\n")
	transaction := multicoinwallet.GetTransaction(&coinType, "dd2701f95a6e3a59868b10465256e839ef130d0a6728fbb92660e4061867ddb2")
	fmt.Printf("Get transaction:\n%s\n", transaction)

	fmt.Printf("Test send to address:\n")
	txId := multicoinwallet.SendToAddress(&coinType, "16kGvbDUATfGfqqNwLxWPzncVWHNJyScjc", 1.0)
	fmt.Printf("Send to address: %s\n", txId)

	fmt.Printf("Test wallet lock:\n")
	lock := multicoinwallet.WalletLock(&coinType)
	if lock != nil {
		fmt.Printf("Wallet lock fail: %v\n", lock)
	} else {
		fmt.Printf("Wallet lock:ok\n")
	}

	pass := multicoinwallet.WalletPassphrase(&coinType)
	if pass != nil {
		fmt.Printf("Wallet passphrase fail: %v\n", lock)
	} else {
		fmt.Printf("Wallet passphrase: ok\n")
	}

	fmt.Printf("Test BCH end\n")
}
