package multicoinwallet


const (
	BITCOIN = 0
	BITCOIN_CASH = 1
)

type CoinType struct {
	Name       string
	Symbol     string
	Bip44Index uint32
}

func AddCoin(coin *CoinType) {
	if coin.Bip44Index == BITCOIN {
		BTCInit()
	}else if coin.Bip44Index == BITCOIN_CASH {
		BCHInit()
	}
}


func GetBalance(coin *CoinType) float64 {
	if (coin.Bip44Index == BITCOIN) || (coin.Bip44Index == BITCOIN_CASH) {
		return BTCGetBalance(coin)
	}
    return 0.0
}

func GetNewAddress(coin *CoinType, label string) string {
	if  (coin.Bip44Index == BITCOIN) || (coin.Bip44Index == BITCOIN_CASH) {
	    return BTCGetNewAddress(coin)
	}

	return ""
}

func ListTransactions(coin *CoinType, account string, count int, fromStart int) string {
	if  (coin.Bip44Index == BITCOIN) || (coin.Bip44Index == BITCOIN_CASH) {
		return BTCListTransactions(coin, account, count, fromStart)
	}

	return ""
}

func ListReceivedByAddress(coin *CoinType) string {
	if  (coin.Bip44Index == BITCOIN) || (coin.Bip44Index == BITCOIN_CASH) {
		return BTCListReceivedByAddress(coin)
	}

	return ""
}

func GetTransaction(coin *CoinType, txHashId string) string {
	if  (coin.Bip44Index == BITCOIN) || (coin.Bip44Index == BITCOIN_CASH) {
		return BTCGetTransaction(coin, txHashId)
	}

	return ""
}

func SendToAddress(coin *CoinType, address string, amount float64) string {
    WalletPassphrase(coin)
	
	if  (coin.Bip44Index == BITCOIN) || (coin.Bip44Index == BITCOIN_CASH) {
		return BTCSendToAddress(coin, address, amount)
	}
	
	WalletLock(coin)

	return ""
}

func WalletLock(coin *CoinType) error {
	if  (coin.Bip44Index == BITCOIN) || (coin.Bip44Index == BITCOIN_CASH) {
		return BTCWalletLock(coin)
	}

	return nil	
}


func WalletPassphrase(coin *CoinType) error {
	if  (coin.Bip44Index == BITCOIN) || (coin.Bip44Index == BITCOIN_CASH) {
		return BTCWalletPassphrase(coin)
	}

	return nil	
}


