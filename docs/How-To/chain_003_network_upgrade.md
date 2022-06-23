---
title: Network upgrade
---

This section covers how to execute a software upgrade for the Elesto network.



## Submit the Upgrade Request

The first step is to open the proposal for the network upgrade

```sh
elestod tx gov submit-proposal software-upgrade \
 testnet-upgrade-2022-06-21 \
 --upgrade-height 1151280 \
 --deposit 200000000utsp \
 --title "testnet-upgrade-2022-06-21 upgrade" \
 --description "testnet upgrade introducing mint and credentials module" \
 --from elesto1ms2wrq8k04cug7ea6ekf60nfke6a8vu8pwm684 \
 --chain-id elesto-canary-1 \
 -b block
```

List proposals 

```
elestod query gov proposals -o json | jq
```

Vote proposal

```sh
elestod tx gov vote 1 yes -b block -y --chain-id elesto-canary-1 --from elesto1ms2wrq8k04cug7ea6ekf60nfke6a8vu8pwm684 --deposit 
```

To check the voting status:

```
elestod query gov proposal 1 --output json  | jq
```

??? Example "Example: query a proposal"

    ```shell
    ➜ elestod query gov proposal 1 --output json  | jq
    ```

    ```json
    {
        "proposal_id": "1",
        "content": {
            "@type": "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
            "title": "testnet-upgrade-2022-06-21 upgrade",
            "description": "testnet upgrade introducing mint and credentials module",
            "plan": {
            "name": "testnet-upgrade-2022-06-21",
            "time": "0001-01-01T00:00:00Z",
            "height": "1151280",
            "info": "",
            "upgraded_client_state": null
            }
        },
        "status": "PROPOSAL_STATUS_PASSED",
        "final_tally_result": {
            "yes": "267616000000",
            "abstain": "0",
            "no": "0",
            "no_with_veto": "0"
        },
        "submit_time": "2022-06-20T15:07:20.362575929Z",
        "deposit_end_time": "2022-06-22T15:07:20.362575929Z",
        "total_deposit": [
            {
            "denom": "utsp",
            "amount": "220000000"
            }
        ],
        "voting_start_time": "2022-06-20T15:07:20.362575929Z",
        "voting_end_time": "2022-06-22T15:07:20.362575929Z"
    }

    ```

When the proposal has been successfully voted, check the upgrade plan with the command:

```
elestod query upgrade plan -o json  | jq
```

??? Example "Example: check upgrade plan"

    ```shell
    ➜ elestod query upgrade plan -o json  | jq
    ```

    ```json
    {
        "name": "testnet-upgrade-2022-06-21",
        "time": "0001-01-01T00:00:00Z",
        "height": "1151280",
        "info": "",
        "upgraded_client_state": null
    }


    ```


## Prepare the Binaries for Cosmosvisor

Once the government proposal for the upgrade has passed, it is time to install the binaries so Cosmosvisor can perform the upgrade.

!!! Warning
    Cosmosvisor can automatically fetch binaries from the internet, but it is recommended to install the binaries manually to make sure the binaries are correct and they are working on your infrastructure.


The first step is to identify the upgrade name:

```
elestod query upgrade plan -o json  | jq .name
```

??? Example "Example: get the upgrade name"

    ```shell
    ➜ elestod query upgrade plan -o json  | jq .name
    ```

    ```json
    "testnet-upgrade-2022-06-21"
    ```


Then we can create the folder for the upgrade in the Cosmosvisor folder structure:

```
mkdir -p .elesto/cosmovisor/upgrades/$UPGRADE_NAME/bin
```

where 

- `$UPGRADE_NAME` is the name of the upgrade obtained in the previous command

??? Example "Example: create the upgrade folder"

    ```shell
    ➜ mkdir -p .elesto/cosmovisor/upgrades/testnet-upgrade-2022-06-21/bin
    ```

Now we can download the new node binary, make it executable, check if it's working, and move it to the upgrade folder


```
curl -LO $BINARY_URL
chmod +x elestod
./elestod version
$NEW_NODE_VERSION
mv elestod .elesto/cosmovisor/upgrades/$UPGRADE_NAME/bin
```

where 

- `$BINARY_URL` is the URL pointing to the binary compatible with your architecture and operative system.
- `$NEW_NODE_VERSION` is the node version that will run after the upgrade.
  
??? Example "Example: download the binary "

    ```shell
    ➜ curl -LO https://github.com/elesto-dao/elesto/releases/download/v2.0.0-rc1/elestod
    ➜ chmod +x elestod
    ➜ ./elestod version #verify that the binary is working properly
    2.0.0-rc1
    ➜ mv elestod .elesto/cosmovisor/upgrades/testnet-upgrade-2022-06-21/bin
    ```

That's it! the Cosmosvisor software should take care of automatically upgrading the node.

For more information about how Cosmosvisor will apply the upgrade, check the [dedicated reference documentation](https://docs.cosmos.network/master/run-node/cosmovisor.html#detecting-upgrades).