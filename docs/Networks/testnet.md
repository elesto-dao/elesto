---
title: Joining Testnet
nodeName:  elesto-guide-1
chainId: elesto-canary-1
---

# Joining Testnet


!!! important
    Before you start, make sure you have installed the [Elesto binary](./node.md).


## Testnet chain ID

The Elesto testnet chain ID is **`{{ chainId }}`**

## Initialize Elesto node

Choose a name for your node, for this guide we'll be using `{{ nodeName }}`. The command will create the node configurations:

```shell
elestod init {{ nodeName }} --chain-id={{ chainId }}
```

Edit the configuration file `config.toml` with your editor of choice:

```shell
edit ~/.elesto/config/config.toml 
```

Under the section *P2P Configuration Options* edit the seeds and peers configuration entries as follows:

> TODO: add the nodes and peers 

```shell
# Comma separated list of seed nodes to connect to
seeds = ""

# Comma separated list of nodes to keep persistent connections to
persistent_peers = ""
```



## Set up Cosmovisor

Cosmovisor allows automatic upgrades for the node, and it is the recommended way of running an Elesto node.

Install the Cosmovisor binary:

```shell
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0
```

Create the required folder structure:

```shell
mkdir -p ~/.elesto/cosmovisor
mkdir -p ~/.elesto/cosmovisor/genesis
mkdir -p ~/.elesto/cosmovisor/genesis/bin
mkdir -p ~/.elesto/cosmovisor/upgrades
```

Set-up the environment variables:

```shell
echo "# Setup Cosmovisor" >> ~/.profile
echo "export DAEMON_NAME=elestod" >> ~/.profile
echo "export DAEMON_HOME=$HOME/.elesto" >> ~/.profile
echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=false" >> ~/.profile
echo "export DAEMON_LOG_BUFFER_SIZE=512" >> ~/.profile
echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.profile
echo "export UNSAFE_SKIP_BACKUP=true" >> ~/.profile
source ~/.profile
```

You may leave out `UNSAFE_SKIP_BACKUP=true`, however, the backup takes a decent amount of time, and public snapshots of old states are available.



Download and replace the genesis file:

```shell
cd ~/.elesto
# TODO: when the repos are public
# the correct url is https://github.com/elesto-dao/networks/raw/main/elesto-canary-1/genesis.tar.bz2
curl -L -O https://filedn.eu/lDrGVxedryyQhhwkHmVd7sJ/genesis.tar.bz2 
tar -xjf genesis.tar.bz2 && rm genesis.tar.bz2
```




Copy the current elestod binary into the cosmovisor/genesis folder:

```shell
cp ~/go/bin/elestod ~/.elesto/cosmovisor/genesis/bin
```

To check your work, ensure the version of cosmovisor and elestod are the same:

```shell
cosmovisor version
elestod version
```

These two command should both output 7.0.3

Reset private validator file to genesis state:

```
elestod unsafe-reset-all
```

### Set-up Elesto Service

Set up a service to allow cosmovisor to run in the background as well as restart automatically if it runs into any problems:

```shell
cd ~
echo "[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=elestod"
Environment="DAEMON_HOME=${HOME}/.elesto"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_LOG_BUFFER_SIZE=512"
Environment="UNSAFE_SKIP_BACKUP=true"
User=$USER
ExecStart=${HOME}/go/bin/cosmovisor start
Restart=always
RestartSec=3
LimitNOFILE=infinity
LimitNPROC=infinity
[Install]
WantedBy=multi-user.target
" >cosmovisor.service
 
```

Install the service:

```shell
sudo mv cosmovisor.service /lib/systemd/system/cosmovisor.service
sudo systemctl daemon-reload
sudo systemctl restart systemd-journald
```

### Operate the Elesto service

Start the service:

```shell
sudo systemctl start cosmovisor
```

Check the service status:

```shell
sudo systemctl status cosmovisor
```

Inspect the service logs:

```shell
journalctl -u cosmovisor -f
```


### Sync the node 

After starting the junod daemon, the chain will begin to sync to the network. The time to sync to the network will vary depending on your setup, but could take a very long time. To query the status of your node:

```shell
# Query via the RPC (default port: 26657)
curl http://localhost:26657/status | jq .result.sync_info.catching_up
```

If this command returns true then your node is still catching up. If it returns `false` then your node has caught up to the network current block and you are safe to proceed to upgrade to a validator node.

> TODO: state-sync and backups




## Upgrading to validator

??? Important "Keys and Balances" 
    To became a validator you need an account with a positive balance. Follow the [keys management](../How-To/chain_002_key_management.md) to learn how to create an account and the [faucet](../How-To/chain_001_faucet.md) how-to to learn how to get tokens for the testnet. 


To upgrade the node to a validator, you will need to submit a create-validator transaction:

```shell
elestod tx staking create-validator \
--chain-id="{{ chainId }}" \
--pubkey=$(elestod tendermint show-validator) \
--amount=[staking_amount_utsp] \
--commission-rate="[commission_rate]" \
--commission-max-rate="[maximum_commission_rate]" \
--commission-max-change-rate="[maximum_rate_of_change_of_commission]" \
--min-self-delegation="[min_self_delegation_amount]" \
--moniker="{{ nodeName }}" \
--security-contact="[security contact email/contact method]" \
--website "wesite of your validator" \
--from=[KEY_NAME]
```

??? Example "Example: Create validator for Alice"

    The following is an example to broadcast the create-validator transaction from Alice wallet 

    ```shell
    elestod tx staking create-validator \
    --chain-id="{{ chainId }}" \
    --pubkey=$(elestod tendermint show-validator) \
    --amount=9000000utsp \
    --commission-rate="0.1" \
    --commission-max-rate="0.2" \
    --commission-max-change-rate="0.1" \
    --min-self-delegation="1" \
    --moniker="{{ nodeName }}" \
    --security-contact="security@example-validator.com" \
    --website "https://example-validator.com" \
    --from=alice
    ```


For a more detailed explanation of the paramters values run the command 

```shell
elestod tx staking create-validator --help
```



### Track Validator Active Set
To see the current validator active set:

```
elestod query staking validators --limit 300 -o json | jq -r '.validators[] |
[.operator_address, .status, (.tokens|tonumber / pow(10; 6)),
.commission.update_time[0:19], .description.moniker] | @csv' | column -t -s","
```

You can search for your specific moniker by adding grep MONIKER at the end:

```
elestod query staking validators --limit 300 -o json | jq -r '.validators[] |
[.operator_address, .status, (.tokens|tonumber / pow(10; 6)),
.commission.update_time[0:19], .description.moniker] | @csv' | column -t -s"," | grep {{ nodeName }}
```

If your bond status is `BOND_STATUS_BONDED`, your validator is part of the active validator set!

