package moonpumpbot

import (
	log "github.com/corgi-kx/logcustom"
	"github.com/gagliardetto/solana-go"
	"test/moonpump"
)

var (
	token      = solana.MustPublicKeyFromBase58("Ef68Rtp4v1WmferTSTRQCyzj1icZiL936RQSdmoYNzDP")
	feeAccount = solana.MustPublicKeyFromBase58("4ezJDzYBSstCXQ2bTaZA4m3gLi6JnZ59D44SKQhqJQqs")
)
var configPda, bondingCurvePda, bondingCurveTokenPda, signerTokenPda, bondingVaultCurvePda solana.PublicKey

func init() {
	var err error
	configPda, _, err = solana.FindProgramAddress([][]byte{
		[]byte("config"),
	}, moonpump.ProgramID)
	if err != nil {
		log.Fatal("Get Moonpump configPda err: ", err)
	}
	bondingCurvePda, _, err = solana.FindProgramAddress([][]byte{
		[]byte("bonding-curve"),
		token.Bytes(),
	}, moonpump.ProgramID)
	if err != nil {
		log.Fatal("Get Moonpump bondingCurve err: ", err)
	}
	bondingCurveTokenPda, _, err = solana.FindAssociatedTokenAddress(bondingCurvePda, token)
	if err != nil {
		log.Fatal("Get Moonpump bondingCurveTokenAccount err: ", err)
	}

	signerTokenPda, _, err = solana.FindAssociatedTokenAddress(signer.PublicKey(), token)
	if err != nil {
		log.Fatal("Get Moonpump buyerTokenAccount err: ", err)
	}

	bondingVaultCurvePda, _, err = solana.FindProgramAddress([][]byte{
		[]byte("bonding-curve"),
		bondingCurvePda.Bytes(),
	}, moonpump.ProgramID)
	if err != nil {
		log.Fatal("Get Moonpump bondingVaultCurve err: ", err)
	}

}

func buy(solAmount uint64) {
	builder := moonpump.NewBuyInstructionBuilder()
	param := moonpump.BuyParams{
		AmountIn:         solAmount,
		MinimumAmountOut: 0,
	}
	builder.SetParams(param)
	builder.SetBuyerAccount(signer.PublicKey())
	builder.SetConfigAccount(configPda)
	builder.SetBondingCurveAccount(bondingCurvePda)
	builder.SetMintAccount(token)
	builder.SetBondingCurveTokenAccountAccount(bondingCurveTokenPda)
	builder.SetBuyerTokenAccountAccount(signerTokenPda)
	builder.SetBondingCurveVaultAccount(bondingVaultCurvePda)
	builder.SetFeeRecipientAccount(feeAccount)
	builder.SetTokenProgramAccount(solana.TokenProgramID)
	builder.SetAssociatedTokenProgramAccount(solana.SPLAssociatedTokenAccountProgramID)
	builder.SetSystemProgramAccount(solana.SystemProgramID)
	instruction := builder.Build()

	signature := sendTx(instruction)
	log.Info("Moonpump buy success signature: ", signature)
}

func sell(tokenAmount uint64) {
	builder := moonpump.NewSellInstructionBuilder()
	param := moonpump.SellParams{
		AmountIn:         tokenAmount,
		MinimumAmountOut: 1000,
	}
	builder.SetParams(param)
	builder.SetSellerAccount(signer.PublicKey())
	builder.SetConfigAccount(configPda)
	builder.SetBondingCurveAccount(bondingCurvePda)
	builder.SetBondingCurveVaultAccount(bondingVaultCurvePda)
	builder.SetMintAccount(token)
	builder.SetBondingCurveTokenAccountAccount(bondingCurveTokenPda)
	builder.SetSellerTokenAccountAccount(signerTokenPda)
	builder.SetFeeRecipientAccount(feeAccount)
	builder.SetTokenProgramAccount(solana.TokenProgramID)
	builder.SetSystemProgramAccount(solana.SystemProgramID)
	instruction := builder.Build()

	signature := sendTx(instruction)
	log.Info("Moonpump sell success signature: ", signature)
}

func launch(param moonpump.PumpParams) {
	builder := moonpump.NewPumpInstructionBuilder()
	builder.SetParams(param)
	builder.SetCreatorAccount(signer.PublicKey())
	builder.SetConfigAccount(configPda)

	mintPda, _, err := solana.FindProgramAddress([][]byte{
		signer.PublicKey().Bytes(),
		[]byte(param.Ticker),
	}, moonpump.ProgramID)
	if err != nil {
		log.Fatal("Get Moonpump bondingCurve err: ", err)
	}
	log.Info("token mint address: ", mintPda)
	builder.SetMintAccount(mintPda)

	newBondingCurvePda, _, err := solana.FindProgramAddress([][]byte{
		[]byte("bonding-curve"),
		mintPda.Bytes(),
	}, moonpump.ProgramID)
	if err != nil {
		log.Fatal("Get Moonpump bondingCurve err: ", err)
	}
	builder.SetBondingCurveAccount(newBondingCurvePda)

	newBondingVaultCurvePda, _, err := solana.FindProgramAddress([][]byte{
		[]byte("bonding-curve"),
		newBondingCurvePda.Bytes(),
	}, moonpump.ProgramID)
	if err != nil {
		log.Fatal("Get Moonpump bondingVaultCurve err: ", err)
	}
	builder.SetBondingCurveVaultAccount(newBondingVaultCurvePda)

	metadataPda, _, err := solana.FindProgramAddress([][]byte{
		[]byte("metadata"),
		solana.TokenMetadataProgramID.Bytes(),
		mintPda.Bytes(),
	}, solana.TokenMetadataProgramID)
	if err != nil {
		log.Fatal("Get Moonpump metadataPda err: ", err)
	}
	builder.SetMetadataAccountAccount(metadataPda)

	newBondingCurveTokenPda, _, err := solana.FindAssociatedTokenAddress(newBondingCurvePda, mintPda)
	if err != nil {
		log.Fatal("Get Moonpump bondingCurveTokenAccount err: ", err)
	}
	builder.SetBondingCurveTokenAccountAccount(newBondingCurveTokenPda)
	builder.SetFeeRecipientAccount(feeAccount)
	builder.SetTokenProgramAccount(solana.TokenProgramID)
	builder.SetTokenMetadataProgramAccount(solana.TokenMetadataProgramID)
	builder.SetAssociatedTokenProgramAccount(solana.SPLAssociatedTokenAccountProgramID)
	builder.SetSystemProgramAccount(solana.SystemProgramID)
	builder.SetRentAccount(solana.SysVarRentPubkey)
	instruction := builder.Build()

	signature := sendTx(instruction)
	log.Info("Moonpump launch success signature: ", signature)
}