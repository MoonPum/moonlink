import BN from 'bn.js';

import {
  ASSOCIATED_TOKEN_PROGRAM_ID,
  getAssociatedTokenAddressSync,
  NATIVE_MINT,
} from '@solana/spl-token';
import {
  AccountMeta,
  Connection,
  PublicKey,
} from '@solana/web3.js';

import { PdaDeriver } from './pda-deriver';
import { PROGRAM_ID } from './pumpTypes';
import * as accounts from './pumpTypes/accounts';
import * as errors from './pumpTypes/errors';
import * as instructions from './pumpTypes/instructions';
import * as types from './pumpTypes/types';

export { accounts, errors, instructions, PdaDeriver, PROGRAM_ID, types };

const METADATA_TOKEN_PROGRAM_ID = new PublicKey(
  "metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s"
);

export class Pump {
  private pumpDeriver: PdaDeriver;
  constructor(public readonly program: PublicKey) {
    this.pumpDeriver = new PdaDeriver(program, METADATA_TOKEN_PROGRAM_ID);
  }

  mint = (creator: PublicKey, ticker: string): PublicKey => {
    return this.pumpDeriver.MintPda(creator, ticker)[0];
  };

  bondingCurveTokenAccount = (
    bondingCurve: PublicKey,
    mint: PublicKey
  ): PublicKey => {
    return getAssociatedTokenAddressSync(mint, bondingCurve, true);
  };

  userTokenAccount = (owner: PublicKey, mint: PublicKey): PublicKey => {
    return getAssociatedTokenAddressSync(mint, owner);
  };

  config = async (connection: Connection): Promise<accounts.Config> => {
    const pda = this.pumpDeriver.ConfigPda()[0];
    const info = await accounts.Config.fromAccountAddress(connection, pda);
    if (!info) {
      throw new Error("Config not found");
    }
    return info;
  };

  bondingCurve = async (
    connection: Connection,
    mint: PublicKey
  ): Promise<accounts.BondingCurve> => {
    const pda = this.pumpDeriver.BondingCurvePda(mint)[0];
    const info = await accounts.BondingCurve.fromAccountAddress(
      connection,
      pda
    );
    if (!info) {
      throw new Error("BondingCurve not found");
    }
    return info;
  };

  initialize = async (
    admin: PublicKey,
    feeRecipient: PublicKey,
    operator: PublicKey,
    tradeFeeBasisPoints: number | null,
    pumpFee: number | null
  ) => {
    const config = this.pumpDeriver.ConfigPda()[0];
    return instructions.createInitializeInstruction(
      {
        admin,
        config,
        feeRecipient,
        operator,
      } satisfies instructions.InitializeInstructionAccounts,
      {
        params: {
          tradeFeeBasisPoints,
          pumpFee,
        } satisfies types.InitializeParams,
      } satisfies instructions.InitializeInstructionArgs,
      this.program
    );
  };

  updateConfig = async (
    admin: PublicKey,
    feeRecipient: PublicKey | null,
    operator: PublicKey | null,
    tradeFeeBasisPoints: number | null,
    pumpFee: number | null
  ) => {
    const config = this.pumpDeriver.ConfigPda()[0];
    return instructions.createUpdateConfigInstruction(
      {
        admin,
        config,
      } satisfies instructions.UpdateConfigInstructionAccounts,
      {
        params: {
          feeRecipient,
          operator,
          tradeFeeBasisPoints,
          pumpFee,
        } satisfies types.UpdateConfigParams,
      } satisfies instructions.UpdateConfigInstructionArgs,
      this.program
    );
  };

  transferOwnership = async (admin: PublicKey, newAdmin: PublicKey) => {
    const config = this.pumpDeriver.ConfigPda()[0];
    return instructions.createTransferOwnershipInstruction(
      {
        admin,
        config,
      } satisfies instructions.TransferOwnershipInstructionAccounts,
      {
        params: {
          newAdmin,
        } satisfies types.TransferOwnershipParams,
      } satisfies instructions.TransferOwnershipInstructionArgs,
      this.program
    );
  };

  pump = async (
    connection: Connection,
    creator: PublicKey,
    name: string,
    ticker: string,
    uri: string
  ) => {
    const config = this.pumpDeriver.ConfigPda()[0];
    const mint = this.pumpDeriver.MintPda(creator, ticker)[0];
    const bondingCurve = this.pumpDeriver.BondingCurvePda(mint)[0];
    const bondingCurveVault =
      this.pumpDeriver.bondingCurveVaultPda(bondingCurve)[0];
    const metadataAccount = this.pumpDeriver.MetadataPda(mint)[0];
    const bondingCurveTokenAccount = this.bondingCurveTokenAccount(
      bondingCurve,
      mint
    );

    const configInfo = await this.config(connection);

    return instructions.createPumpInstruction(
      {
        creator,
        config,
        feeRecipient: configInfo.feeRecipient,
        mint,
        bondingCurve,
        bondingCurveVault,
        metadataAccount,
        bondingCurveTokenAccount,
        tokenMetadataProgram: METADATA_TOKEN_PROGRAM_ID,
        associatedTokenProgram: ASSOCIATED_TOKEN_PROGRAM_ID,
      } satisfies instructions.PumpInstructionAccounts,
      {
        params: {
          name,
          ticker,
          uri,
        } satisfies types.PumpParams,
      } satisfies instructions.PumpInstructionArgs,
      this.program
    );
  };

  vanityPump = async (
    connection: Connection,
    creator: PublicKey,
    seeds: Uint8Array,
    name: string,
    ticker: string,
    uri: string
  ) => {
    const config = this.pumpDeriver.ConfigPda()[0];
    const mint = this.pumpDeriver.VanityMintPda(seeds)[0];
    const bondingCurve = this.pumpDeriver.BondingCurvePda(mint)[0];
    const bondingCurveVault =
      this.pumpDeriver.bondingCurveVaultPda(bondingCurve)[0];
    const metadataAccount = this.pumpDeriver.MetadataPda(mint)[0];
    const bondingCurveTokenAccount = this.bondingCurveTokenAccount(
      bondingCurve,
      mint
    );

    const configInfo = await this.config(connection);

    return instructions.createVanityPumpInstruction(
      {
        creator,
        config,
        feeRecipient: configInfo.feeRecipient,
        mint,
        bondingCurve,
        bondingCurveVault,
        metadataAccount,
        bondingCurveTokenAccount,
        tokenMetadataProgram: METADATA_TOKEN_PROGRAM_ID,
        associatedTokenProgram: ASSOCIATED_TOKEN_PROGRAM_ID,
      } satisfies instructions.VanityPumpInstructionAccounts,
      {
        params: {
          seeds: Array.from(seeds),
          name,
          ticker,
          uri,
        } satisfies types.VanityPumpParams,
      } satisfies instructions.VanityPumpInstructionArgs,
      this.program
    );
  };

  buy = async (
    connection: Connection,
    buyer: PublicKey,
    mint: PublicKey,
    amountIn: BN,
    minimumAmountOut: BN
  ) => {
    const config = this.pumpDeriver.ConfigPda()[0];
    const bondingCurve = this.pumpDeriver.BondingCurvePda(mint)[0];
    const bondingCurveVault =
      this.pumpDeriver.bondingCurveVaultPda(bondingCurve)[0];
    const bondingCurveTokenAccount = this.bondingCurveTokenAccount(
      bondingCurve,
      mint
    );
    const buyerTokenAccount = this.userTokenAccount(buyer, mint);

    const configInfo = await this.config(connection);

    return instructions.createBuyInstruction(
      {
        buyer,
        config,
        bondingCurve,
        mint,
        bondingCurveTokenAccount,
        bondingCurveVault,
        buyerTokenAccount,
        feeRecipient: configInfo.feeRecipient,
        associatedTokenProgram: ASSOCIATED_TOKEN_PROGRAM_ID,
      } satisfies instructions.BuyInstructionAccounts,
      {
        params: {
          amountIn,
          minimumAmountOut,
        } satisfies types.BuyParams,
      } satisfies instructions.BuyInstructionArgs,
      this.program
    );
  };
  sell = async (
    connection: Connection,
    seller: PublicKey,
    mint: PublicKey,
    amountIn: BN,
    minimumAmountOut: BN
  ) => {
    const config = this.pumpDeriver.ConfigPda()[0];
    const bondingCurve = this.pumpDeriver.BondingCurvePda(mint)[0];
    const bondingCurveVault =
      this.pumpDeriver.bondingCurveVaultPda(bondingCurve)[0];
    const bondingCurveTokenAccount = this.bondingCurveTokenAccount(
      bondingCurve,
      mint
    );
    const sellerTokenAccount = this.userTokenAccount(seller, mint);

    const configInfo = await this.config(connection);

    return instructions.createSellInstruction(
      {
        seller,
        config,
        bondingCurve,
        mint,
        bondingCurveTokenAccount,
        bondingCurveVault,
        sellerTokenAccount,
        feeRecipient: configInfo.feeRecipient,
      } satisfies instructions.SellInstructionAccounts,
      {
        params: {
          amountIn,
          minimumAmountOut,
        } satisfies types.SellParams,
      } satisfies instructions.SellInstructionArgs,
      this.program
    );
  };

  graduate = async (
    connection: Connection,
    mint: PublicKey,
    creator: PublicKey,
    nonce: number,
    openTime: number,
    remainingAccounts: AccountMeta[]
  ) => {
    const config = this.pumpDeriver.ConfigPda()[0];
    const bondingCurve = this.pumpDeriver.BondingCurvePda(mint)[0];
    const bondingCurveVault =
      this.pumpDeriver.bondingCurveVaultPda(bondingCurve)[0];
    const bondingCurveTokenAccount = this.bondingCurveTokenAccount(
      bondingCurve,
      mint
    );

    const configInfo = await this.config(connection);

    const operatorTokenAccount = this.userTokenAccount(
      configInfo.operator,
      mint
    );
    const operatorNativeMintTokenAccount = this.userTokenAccount(
      configInfo.operator,
      NATIVE_MINT
    );

    return instructions.createGraduateInstruction(
      {
        operator: configInfo.operator,
        config,
        bondingCurve,
        mint,
        bondingCurveTokenAccount,
        bondingCurveVault,
        nativeMint: NATIVE_MINT,
        admin: configInfo.admin,
        creator,
        operatorTokenAccount,
        operatorNativeMintTokenAccount,
        anchorRemainingAccounts: remainingAccounts,
      } satisfies instructions.GraduateInstructionAccounts,
      {
        params: {
          nonce,
          openTime,
        } satisfies types.GrudateParams,
      } satisfies instructions.GraduateInstructionArgs,
      this.program
    );
  };
}
