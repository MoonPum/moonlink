// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package moonpump

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// VanityPump is the `vanityPump` instruction.
type VanityPump struct {
	Params *VanityPumpParams

	// [0] = [WRITE, SIGNER] creator
	//
	// [1] = [] config
	//
	// [2] = [WRITE] mint
	//
	// [3] = [WRITE] bondingCurve
	//
	// [4] = [WRITE] bondingCurveVault
	//
	// [5] = [WRITE] metadataAccount
	//
	// [6] = [WRITE] bondingCurveTokenAccount
	//
	// [7] = [WRITE] feeRecipient
	//
	// [8] = [] tokenProgram
	//
	// [9] = [] tokenMetadataProgram
	//
	// [10] = [] associatedTokenProgram
	//
	// [11] = [] systemProgram
	//
	// [12] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewVanityPumpInstructionBuilder creates a new `VanityPump` instruction builder.
func NewVanityPumpInstructionBuilder() *VanityPump {
	nd := &VanityPump{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *VanityPump) SetParams(params VanityPumpParams) *VanityPump {
	inst.Params = &params
	return inst
}

// SetCreatorAccount sets the "creator" account.
func (inst *VanityPump) SetCreatorAccount(creator ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(creator).WRITE().SIGNER()
	return inst
}

// GetCreatorAccount gets the "creator" account.
func (inst *VanityPump) GetCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetConfigAccount sets the "config" account.
func (inst *VanityPump) SetConfigAccount(config ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *VanityPump) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintAccount sets the "mint" account.
func (inst *VanityPump) SetMintAccount(mint ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *VanityPump) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetBondingCurveAccount sets the "bondingCurve" account.
func (inst *VanityPump) SetBondingCurveAccount(bondingCurve ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(bondingCurve).WRITE()
	return inst
}

// GetBondingCurveAccount gets the "bondingCurve" account.
func (inst *VanityPump) GetBondingCurveAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBondingCurveVaultAccount sets the "bondingCurveVault" account.
func (inst *VanityPump) SetBondingCurveVaultAccount(bondingCurveVault ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(bondingCurveVault).WRITE()
	return inst
}

// GetBondingCurveVaultAccount gets the "bondingCurveVault" account.
func (inst *VanityPump) GetBondingCurveVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccountAccount sets the "metadataAccount" account.
func (inst *VanityPump) SetMetadataAccountAccount(metadataAccount ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadataAccount).WRITE()
	return inst
}

// GetMetadataAccountAccount gets the "metadataAccount" account.
func (inst *VanityPump) GetMetadataAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetBondingCurveTokenAccountAccount sets the "bondingCurveTokenAccount" account.
func (inst *VanityPump) SetBondingCurveTokenAccountAccount(bondingCurveTokenAccount ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(bondingCurveTokenAccount).WRITE()
	return inst
}

// GetBondingCurveTokenAccountAccount gets the "bondingCurveTokenAccount" account.
func (inst *VanityPump) GetBondingCurveTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetFeeRecipientAccount sets the "feeRecipient" account.
func (inst *VanityPump) SetFeeRecipientAccount(feeRecipient ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(feeRecipient).WRITE()
	return inst
}

// GetFeeRecipientAccount gets the "feeRecipient" account.
func (inst *VanityPump) GetFeeRecipientAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *VanityPump) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *VanityPump) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *VanityPump) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *VanityPump) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *VanityPump) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *VanityPump) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *VanityPump) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *VanityPump) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetRentAccount sets the "rent" account.
func (inst *VanityPump) SetRentAccount(rent ag_solanago.PublicKey) *VanityPump {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *VanityPump) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst VanityPump) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_VanityPump,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst VanityPump) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *VanityPump) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Creator is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.BondingCurve is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.BondingCurveVault is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.MetadataAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.BondingCurveTokenAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.FeeRecipient is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *VanityPump) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("VanityPump")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               creator", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                  mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          bondingCurve", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     bondingCurveVault", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("              metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     bondingCurveToken", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("          feeRecipient", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("  tokenMetadataProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj VanityPump) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *VanityPump) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewVanityPumpInstruction declares a new VanityPump instruction with the provided parameters and accounts.
func NewVanityPumpInstruction(
	// Parameters:
	params VanityPumpParams,
	// Accounts:
	creator ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	bondingCurve ag_solanago.PublicKey,
	bondingCurveVault ag_solanago.PublicKey,
	metadataAccount ag_solanago.PublicKey,
	bondingCurveTokenAccount ag_solanago.PublicKey,
	feeRecipient ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *VanityPump {
	return NewVanityPumpInstructionBuilder().
		SetParams(params).
		SetCreatorAccount(creator).
		SetConfigAccount(config).
		SetMintAccount(mint).
		SetBondingCurveAccount(bondingCurve).
		SetBondingCurveVaultAccount(bondingCurveVault).
		SetMetadataAccountAccount(metadataAccount).
		SetBondingCurveTokenAccountAccount(bondingCurveTokenAccount).
		SetFeeRecipientAccount(feeRecipient).
		SetTokenProgramAccount(tokenProgram).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
