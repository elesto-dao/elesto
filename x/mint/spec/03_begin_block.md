<!--
order: 3
-->

# Begin-Block

Begin block operation for the `mint` module calculates `BlockInflation` to mint coins to be sent to the fee collector
and sets `LastBlockTime` value. It is worth noting that there is no inflation in genesis block as it doesn't have
`LastBlockTime`.

Begin block operation for the `mint` module executes the minting routine following the schedule defined in `abci.go`.

For each block, the `BeginBlocker` method calculates the current year based on the `blocksPerYear` estimated value, and mints the amount associated with it.

If the lookup fails, no coins are minted.