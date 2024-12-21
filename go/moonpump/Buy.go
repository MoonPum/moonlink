// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package moonpump

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Buy is the `buy` instruction.
type Buy struct {
	Params *BuyParams

	// [0] = [WRITE, SIGNER] buyer
	//
	// [1] = [] config
	//
	// [2] = [WRITE] bondingCurve
	//
	// [3] = [] mint
	//
	// [4] = [WRITE] bondingCurveTokenAccount
	//
	// [5] = [WRITE] buyerTokenAccount
	//
	// [6] = [WRITE] bondingCurveVault
	//
	// [7] = [WRITE] feeRecipient
	//
	// [8] = [] tokenProgram
	//
	// [9] = [] associatedTokenProgram
	//
	// [10] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBuyInstructionBuilder creates a new `Buy` instruction builder.
func NewBuyInstructionBuilder() *Buy {
	nd := &Buy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *Buy) SetParams(params BuyParams) *Buy {
	inst.Params = &params
	return inst
}

// SetBuyerAccount sets the "buyer" account.
func (inst *Buy) SetBuyerAccount(buyer ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(buyer).WRITE().SIGNER()
	return inst
}

// GetBuyerAccount gets the "buyer" account.
func (inst *Buy) GetBuyerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetConfigAccount sets the "config" account.
func (inst *Buy) SetConfigAccount(config ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *Buy) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetBondingCurveAccount sets the "bondingCurve" account.
func (inst *Buy) SetBondingCurveAccount(bondingCurve ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(bondingCurve).WRITE()
	return inst
}

// GetBondingCurveAccount gets the "bondingCurve" account.
func (inst *Buy) GetBondingCurveAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAccount sets the "mint" account.
func (inst *Buy) SetMintAccount(mint ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *Buy) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBondingCurveTokenAccountAccount sets the "bondingCurveTokenAccount" account.
func (inst *Buy) SetBondingCurveTokenAccountAccount(bondingCurveTokenAccount ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(bondingCurveTokenAccount).WRITE()
	return inst
}

// GetBondingCurveTokenAccountAccount gets the "bondingCurveTokenAccount" account.
func (inst *Buy) GetBondingCurveTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetBuyerTokenAccountAccount sets the "buyerTokenAccount" account.
func (inst *Buy) SetBuyerTokenAccountAccount(buyerTokenAccount ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(buyerTokenAccount).WRITE()
	return inst
}

// GetBuyerTokenAccountAccount gets the "buyerTokenAccount" account.
func (inst *Buy) GetBuyerTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetBondingCurveVaultAccount sets the "bondingCurveVault" account.
func (inst *Buy) SetBondingCurveVaultAccount(bondingCurveVault ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(bondingCurveVault).WRITE()
	return inst
}

// GetBondingCurveVaultAccount gets the "bondingCurveVault" account.
func (inst *Buy) GetBondingCurveVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetFeeRecipientAccount sets the "feeRecipient" account.
func (inst *Buy) SetFeeRecipientAccount(feeRecipient ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(feeRecipient).WRITE()
	return inst
}

// GetFeeRecipientAccount gets the "feeRecipient" account.
func (inst *Buy) GetFeeRecipientAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Buy) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Buy) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *Buy) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *Buy) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Buy) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Buy) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst Buy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Buy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Buy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Buy) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Buyer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.BondingCurve is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.BondingCurveTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.BuyerTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.BondingCurveVault is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.FeeRecipient is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *Buy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Buy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 buyer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          bondingCurve", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                  mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     bondingCurveToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("            buyerToken", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     bondingCurveVault", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("          feeRecipient", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj Buy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Buy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewBuyInstruction declares a new Buy instruction with the provided parameters and accounts.
func NewBuyInstruction(
	// Parameters:
	params BuyParams,
	// Accounts:
	buyer ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	bondingCurve ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	bondingCurveTokenAccount ag_solanago.PublicKey,
	buyerTokenAccount ag_solanago.PublicKey,
	bondingCurveVault ag_solanago.PublicKey,
	feeRecipient ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *Buy {
	return NewBuyInstructionBuilder().
		SetParams(params).
		SetBuyerAccount(buyer).
		SetConfigAccount(config).
		SetBondingCurveAccount(bondingCurve).
		SetMintAccount(mint).
		SetBondingCurveTokenAccountAccount(bondingCurveTokenAccount).
		SetBuyerTokenAccountAccount(buyerTokenAccount).
		SetBondingCurveVaultAccount(bondingCurveVault).
		SetFeeRecipientAccount(feeRecipient).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetSystemProgramAccount(systemProgram)
}