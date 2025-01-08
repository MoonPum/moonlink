import { PublicKey } from '@solana/web3.js';

const CONFIG_SEEDS = Buffer.from("config", "utf-8");
const BONDING_CURVE_SEEDS = Buffer.from("bonding-curve", "utf-8");
const METADATA_SEEDS = Buffer.from("metadata", "utf-8");

export class PdaDeriver {
  constructor(
    public readonly program: PublicKey,
    public readonly metadataTokenProgram: PublicKey
  ) {}

  ConfigPda = (): [PublicKey, number] => {
    return PublicKey.findProgramAddressSync([CONFIG_SEEDS], this.program);
  };

  BondingCurvePda = (mint: PublicKey): [PublicKey, number] => {
    return PublicKey.findProgramAddressSync(
      [BONDING_CURVE_SEEDS, mint.toBuffer()],
      this.program
    );
  };

  bondingCurveVaultPda = (bondingCurve: PublicKey): [PublicKey, number] => {
    return PublicKey.findProgramAddressSync(
      [BONDING_CURVE_SEEDS, bondingCurve.toBuffer()],
      this.program
    );
  };

  MintPda = (creator: PublicKey, ticker: string): [PublicKey, number] => {
    return PublicKey.findProgramAddressSync(
      [creator.toBuffer(), Buffer.from(ticker, "utf-8")],
      this.program
    );
  };

  VanityMintPda = (seeds: Uint8Array): [PublicKey, number] => {
    return PublicKey.findProgramAddressSync([seeds], this.program);
  };

  MetadataPda = (mint: PublicKey): [PublicKey, number] => {
    return PublicKey.findProgramAddressSync(
      [METADATA_SEEDS, this.metadataTokenProgram.toBuffer(), mint.toBuffer()],
      this.metadataTokenProgram
    );
  };
}
