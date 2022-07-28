# Mint module

The mint module takes on the task of handling the creation and destruction of tokens based on the current status of the circulating supply.
The minting mechanism allows for a flexible inflation rate determined by market demand targeting a particular bonded-stake ratio and effects a balance between market liquidity and staked supply.

By default, the Cosmos SDK mint module allows for an infinite supply. To meet a different set of requirements, this custom mint module specifies a fixed total token supply and a custom minting schedule.

- For each block, the chain looks up the predefined token amount to be minted and mints it.
- After the network goes past 10 years of activity, the minting schedule halts, and no more tokens are added to the total supply.

From a security perspective, reasonable care has been taken to preserve the current block year, making the mint module capable of working even in case of an entire chain reset.
