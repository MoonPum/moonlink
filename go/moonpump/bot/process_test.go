package moonpumpbot

import (
	"fmt"
	"github.com/gagliardetto/solana-go"
	"math/rand"
	"test/moonpump"
	"testing"
	"time"
)

func TestLaunch(t *testing.T) {
	param := moonpump.PumpParams{
		Name:   "tradingview",
		Ticker: "TV",
		Uri:    "https://gateway.pinata.cloud/ipfs/QmTzeofN9zr6BNd4iAFgEtovh2N1zps7VMghrk8pyzNWQz",
	}
	launch(param)
}

func TestSell(t *testing.T) {
	tokenAmount := tokenBalance(signer.PublicKey())
	fmt.Println(tokenAmount)
	sell(tokenAmount)
}

func TestBuy(t *testing.T) {
	buy(5 * solana.LAMPORTS_PER_SOL)
}
