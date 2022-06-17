
# Get Testnet Tokens (Faucet)

!!! Hint 
    To learn how to manage keys, see the [Manage Keys](./chain_002_key_management.md) how-to.

The Elesto native token denom is `tsp`. To obtain `tsp` tokens from the testnet faucet, use the following command:

```sh 
curl -X POST -d $KEY_ADDRESS $FAUCET_URL 
```

where:

- `$KEY_ADDRESS` is the blockchain address you want to top-up
- `$FAUCET_URL` is the URL of the faucet service

??? Example "Example: get tokens for Alice's account"
    
    ```shell
    curl -X POST \
    -d "{\"address\": \"$(cosmos-cashd keys show alice -a)\"}" \
    https://faucet.cosmos-cash.app.beta.starport.cloud

    ```
