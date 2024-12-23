/**
 * This code was GENERATED using the solita package.
 * Please DO NOT EDIT THIS FILE, instead rerun solita to update it or write a wrapper to add functionality.
 *
 * See: https://github.com/metaplex-foundation/solita
 */

import * as beet from '@metaplex-foundation/beet'
export type BuyParams = {
  amountIn: beet.bignum
  minimumAmountOut: beet.bignum
}

/**
 * @category userTypes
 * @category generated
 */
export const buyParamsBeet = new beet.BeetArgsStruct<BuyParams>(
  [
    ['amountIn', beet.u64],
    ['minimumAmountOut', beet.u64],
  ],
  'BuyParams'
)
