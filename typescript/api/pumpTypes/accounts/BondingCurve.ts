/**
 * This code was GENERATED using the solita package.
 * Please DO NOT EDIT THIS FILE, instead rerun solita to update it or write a wrapper to add functionality.
 *
 * See: https://github.com/metaplex-foundation/solita
 */

import * as web3 from '@solana/web3.js'
import * as beet from '@metaplex-foundation/beet'
import * as beetSolana from '@metaplex-foundation/beet-solana'
import { Status, statusBeet } from '../types/Status'

/**
 * Arguments used to create {@link BondingCurve}
 * @category Accounts
 * @category generated
 */
export type BondingCurveArgs = {
  creator: web3.PublicKey
  virtualSolReserves: beet.bignum
  virtualTokenReserves: beet.bignum
  realSolReserves: beet.bignum
  realTokenReserves: beet.bignum
  k: beet.bignum
  status: Status
  bump: number
  mint: web3.PublicKey
  padding: Uint8Array
}

export const bondingCurveDiscriminator = [23, 183, 248, 55, 96, 216, 172, 96]
/**
 * Holds the data for the {@link BondingCurve} Account and provides de/serialization
 * functionality for that data
 *
 * @category Accounts
 * @category generated
 */
export class BondingCurve implements BondingCurveArgs {
  private constructor(
    readonly creator: web3.PublicKey,
    readonly virtualSolReserves: beet.bignum,
    readonly virtualTokenReserves: beet.bignum,
    readonly realSolReserves: beet.bignum,
    readonly realTokenReserves: beet.bignum,
    readonly k: beet.bignum,
    readonly status: Status,
    readonly bump: number,
    readonly mint: web3.PublicKey,
    readonly padding: Uint8Array
  ) {}

  /**
   * Creates a {@link BondingCurve} instance from the provided args.
   */
  static fromArgs(args: BondingCurveArgs) {
    return new BondingCurve(
      args.creator,
      args.virtualSolReserves,
      args.virtualTokenReserves,
      args.realSolReserves,
      args.realTokenReserves,
      args.k,
      args.status,
      args.bump,
      args.mint,
      args.padding
    )
  }

  /**
   * Deserializes the {@link BondingCurve} from the data of the provided {@link web3.AccountInfo}.
   * @returns a tuple of the account data and the offset up to which the buffer was read to obtain it.
   */
  static fromAccountInfo(
    accountInfo: web3.AccountInfo<Buffer>,
    offset = 0
  ): [BondingCurve, number] {
    return BondingCurve.deserialize(accountInfo.data, offset)
  }

  /**
   * Retrieves the account info from the provided address and deserializes
   * the {@link BondingCurve} from its data.
   *
   * @throws Error if no account info is found at the address or if deserialization fails
   */
  static async fromAccountAddress(
    connection: web3.Connection,
    address: web3.PublicKey,
    commitmentOrConfig?: web3.Commitment | web3.GetAccountInfoConfig
  ): Promise<BondingCurve> {
    const accountInfo = await connection.getAccountInfo(
      address,
      commitmentOrConfig
    )
    if (accountInfo == null) {
      throw new Error(`Unable to find BondingCurve account at ${address}`)
    }
    return BondingCurve.fromAccountInfo(accountInfo, 0)[0]
  }

  /**
   * Provides a {@link web3.Connection.getProgramAccounts} config builder,
   * to fetch accounts matching filters that can be specified via that builder.
   *
   * @param programId - the program that owns the accounts we are filtering
   */
  static gpaBuilder(
    programId: web3.PublicKey = new web3.PublicKey(
      'DrcCNopZABtZev6WkiBDsrwPfbJAcexx5jnVDZ7egLmS'
    )
  ) {
    return beetSolana.GpaBuilder.fromStruct(programId, bondingCurveBeet)
  }

  /**
   * Deserializes the {@link BondingCurve} from the provided data Buffer.
   * @returns a tuple of the account data and the offset up to which the buffer was read to obtain it.
   */
  static deserialize(buf: Buffer, offset = 0): [BondingCurve, number] {
    return bondingCurveBeet.deserialize(buf, offset)
  }

  /**
   * Serializes the {@link BondingCurve} into a Buffer.
   * @returns a tuple of the created Buffer and the offset up to which the buffer was written to store it.
   */
  serialize(): [Buffer, number] {
    return bondingCurveBeet.serialize({
      accountDiscriminator: bondingCurveDiscriminator,
      ...this,
    })
  }

  /**
   * Returns the byteSize of a {@link Buffer} holding the serialized data of
   * {@link BondingCurve} for the provided args.
   *
   * @param args need to be provided since the byte size for this account
   * depends on them
   */
  static byteSize(args: BondingCurveArgs) {
    const instance = BondingCurve.fromArgs(args)
    return bondingCurveBeet.toFixedFromValue({
      accountDiscriminator: bondingCurveDiscriminator,
      ...instance,
    }).byteSize
  }

  /**
   * Fetches the minimum balance needed to exempt an account holding
   * {@link BondingCurve} data from rent
   *
   * @param args need to be provided since the byte size for this account
   * depends on them
   * @param connection used to retrieve the rent exemption information
   */
  static async getMinimumBalanceForRentExemption(
    args: BondingCurveArgs,
    connection: web3.Connection,
    commitment?: web3.Commitment
  ): Promise<number> {
    return connection.getMinimumBalanceForRentExemption(
      BondingCurve.byteSize(args),
      commitment
    )
  }

  /**
   * Returns a readable version of {@link BondingCurve} properties
   * and can be used to convert to JSON and/or logging
   */
  pretty() {
    return {
      creator: this.creator.toBase58(),
      virtualSolReserves: (() => {
        const x = <{ toNumber: () => number }>this.virtualSolReserves
        if (typeof x.toNumber === 'function') {
          try {
            return x.toNumber()
          } catch (_) {
            return x
          }
        }
        return x
      })(),
      virtualTokenReserves: (() => {
        const x = <{ toNumber: () => number }>this.virtualTokenReserves
        if (typeof x.toNumber === 'function') {
          try {
            return x.toNumber()
          } catch (_) {
            return x
          }
        }
        return x
      })(),
      realSolReserves: (() => {
        const x = <{ toNumber: () => number }>this.realSolReserves
        if (typeof x.toNumber === 'function') {
          try {
            return x.toNumber()
          } catch (_) {
            return x
          }
        }
        return x
      })(),
      realTokenReserves: (() => {
        const x = <{ toNumber: () => number }>this.realTokenReserves
        if (typeof x.toNumber === 'function') {
          try {
            return x.toNumber()
          } catch (_) {
            return x
          }
        }
        return x
      })(),
      k: (() => {
        const x = <{ toNumber: () => number }>this.k
        if (typeof x.toNumber === 'function') {
          try {
            return x.toNumber()
          } catch (_) {
            return x
          }
        }
        return x
      })(),
      status: 'Status.' + Status[this.status],
      bump: this.bump,
      mint: this.mint.toBase58(),
      padding: this.padding,
    }
  }
}

/**
 * @category Accounts
 * @category generated
 */
export const bondingCurveBeet = new beet.FixableBeetStruct<
  BondingCurve,
  BondingCurveArgs & {
    accountDiscriminator: number[] /* size: 8 */
  }
>(
  [
    ['accountDiscriminator', beet.uniformFixedSizeArray(beet.u8, 8)],
    ['creator', beetSolana.publicKey],
    ['virtualSolReserves', beet.u64],
    ['virtualTokenReserves', beet.u64],
    ['realSolReserves', beet.u64],
    ['realTokenReserves', beet.u64],
    ['k', beet.u128],
    ['status', statusBeet],
    ['bump', beet.u8],
    ['mint', beetSolana.publicKey],
    ['padding', beet.bytes],
  ],
  BondingCurve.fromArgs,
  'BondingCurve'
)
