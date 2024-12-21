package moonpumpbot

import (
	"context"
	"github.com/gagliardetto/solana-go"
	computebudget "github.com/gagliardetto/solana-go/programs/compute-budget"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"log"
	"math/big"
	"strconv"
	"test/moonpump"
)

var wsClient *ws.Client
var rpcClient *rpc.Client
var signer *solana.Wallet

func init() {
	var err error
	signer, err = solana.WalletFromPrivateKeyBase58("pri")
	if err != nil {
		panic(err)
	}
	wsClient, err = ws.Connect(context.Background(), "rpc")
	if err != nil {
		panic(err)
	}

	rpcClient = rpc.New("https://rpc")

}

func sendTx(instruction *moonpump.Instruction) solana.Signature {
	computeLimitInstruction, err := computebudget.NewSetComputeUnitLimitInstructionBuilder().SetUnits(300000).ValidateAndBuild()
	if err != nil {
		panic(err)
	}
	computeUnitPriceInstruction, err := computebudget.NewSetComputeUnitPriceInstructionBuilder().SetMicroLamports(300000).ValidateAndBuild()
	if err != nil {
		panic(err)
	}
	tx, err := solana.NewTransaction([]solana.Instruction{
		computeLimitInstruction,
		computeUnitPriceInstruction,
		instruction,
	}, recentHash(), solana.TransactionPayer(signer.PublicKey()))
	if err != nil {
		panic(err)
	}

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if signer.PublicKey().Equals(key) {
			return &signer.PrivateKey
		}
		return nil
	})
	if err != nil {
		log.Fatal("failed to sign transaction: ", err)
	}

	//resp, err := rpcClient.SimulateTransaction(context.Background(), tx)
	//if err != nil {
	//	log.Fatal("failed to Simulate transaction: ", err)
	//}
	//
	//log.Println(resp.Value.Logs)

	sig, err := confirm.SendAndConfirmTransaction(context.Background(), rpcClient, wsClient, tx)
	if err != nil {
		log.Fatal("failed to send transaction: ", err)
	}
	return sig
}

func recentHash() solana.Hash {
	resp, err := rpcClient.GetLatestBlockhash(context.Background(), rpc.CommitmentRecent)
	if err != nil {
		panic(err)
	}
	return resp.Value.Blockhash
}

func tokenBalance(account solana.PublicKey) uint64 {
	tokenAccountPda, _, err := solana.FindAssociatedTokenAddress(account, token)
	if err != nil {
		log.Fatal("Get Moonpump tokenAccountPda err: ", err)
	}
	resp, err := rpcClient.GetTokenAccountBalance(context.Background(), tokenAccountPda, rpc.CommitmentConfirmed)
	if err != nil {
		log.Fatal("failed to GetTokenAccountBalance: ", err)
	}
	balance, err := strconv.ParseUint(resp.Value.Amount, 10, 64)
	if err != nil {
		log.Fatal("failed to parse amount: ", err)
	}
	return balance
}

func LampToSol(lamports uint64) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetUint64(lamports), new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))
}

func SolToLamp(sol float64) uint64 {
	lamports, _ := new(big.Float).Mul(new(big.Float).SetFloat64(sol), new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL)).Uint64()
	return lamports
}