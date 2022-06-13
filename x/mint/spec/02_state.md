<!--
order: 2
-->

# State

## Params

Minting params are held in the global params store.

- Params: `mint/params -> ProtocolBuffer(params)`

### Params values

| Name              | Type     | Description                                     |
| ----------------- | -------- | ----------------------------------------------- |
| `mint_denom`      | `string` | denomination of coin to be minted               |
| `blocks_per_year` | `int64`  | estimate number of blocks in a year             |
| `max_supply`      | `int64`  | total max supply                                |
| `team_address`    | `string` | address to which development reward is assigned |
| `team_reward`     | `string` | amount of development reward                    |
