/**
 * This code was GENERATED using the solita package.
 * Please DO NOT EDIT THIS FILE, instead rerun solita to update it or write a wrapper to add functionality.
 *
 * See: https://github.com/metaplex-foundation/solita
 */

import * as beet from '@metaplex-foundation/beet'
export type InitializeParams = {
  tradeFeeBasisPoints: beet.COption<number>
  pumpFee: beet.COption<beet.bignum>
}

/**
 * @category userTypes
 * @category generated
 */
export const initializeParamsBeet =
  new beet.FixableBeetArgsStruct<InitializeParams>(
    [
      ['tradeFeeBasisPoints', beet.coption(beet.u16)],
      ['pumpFee', beet.coption(beet.u64)],
    ],
    'InitializeParams'
  )
