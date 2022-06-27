#!/bin/bash
rm -rf $HOME/.elestod/

# make four elesto directories
mkdir $HOME/.elestod
mkdir $HOME/.elestod/validator1
mkdir $HOME/.elestod/validator2
mkdir $HOME/.elestod/validator3

# init all three validators
elestod init --chain-id=testing validator1 --home=$HOME/.elestod/validator1
elestod init --chain-id=testing validator2 --home=$HOME/.elestod/validator2
elestod init --chain-id=testing validator3 --home=$HOME/.elestod/validator3

# create keys for all three validators
elestod keys add validator1 --keyring-backend=test --home=$HOME/.elestod/validator1
elestod keys add validator2 --keyring-backend=test --home=$HOME/.elestod/validator2
elestod keys add validator3 --keyring-backend=test --home=$HOME/.elestod/validator3

# change staking denom to stake
cat $HOME/.elestod/validator1/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="stake"' > $HOME/.elestod/validator1/config/tmp_genesis.json && mv $HOME/.elestod/validator1/config/tmp_genesis.json $HOME/.elestod/validator1/config/genesis.json

# create validator node with tokens to transfer to the three other nodes
elestod add-genesis-account $(elestod keys show validator1 -a --keyring-backend=test --home=$HOME/.elestod/validator1) 200000000000stake --home=$HOME/.elestod/validator1
elestod gentx validator1 500000000stake --keyring-backend=test --home=$HOME/.elestod/validator1 --chain-id=testing
elestod collect-gentxs --home=$HOME/.elestod/validator1


# update staking genesis
cat $HOME/.elestod/validator1/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="240s"' > $HOME/.elestod/validator1/config/tmp_genesis.json && mv $HOME/.elestod/validator1/config/tmp_genesis.json $HOME/.elestod/validator1/config/genesis.json

# update crisis variable to stake
cat $HOME/.elestod/validator1/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="stake"' > $HOME/.elestod/validator1/config/tmp_genesis.json && mv $HOME/.elestod/validator1/config/tmp_genesis.json $HOME/.elestod/validator1/config/genesis.json

# udpate gov genesis
cat $HOME/.elestod/validator1/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="60s"' > $HOME/.elestod/validator1/config/tmp_genesis.json && mv $HOME/.elestod/validator1/config/tmp_genesis.json $HOME/.elestod/validator1/config/genesis.json
cat $HOME/.elestod/validator1/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="stake"' > $HOME/.elestod/validator1/config/tmp_genesis.json && mv $HOME/.elestod/validator1/config/tmp_genesis.json $HOME/.elestod/validator1/config/genesis.json

# port key (validator1 uses default ports)
# validator1 1317, 9090, 9091, 26658, 26657, 26656, 6060
# validator2 1316, 9088, 9089, 26655, 26654, 26653, 6061
# validator3 1315, 9086, 9087, 26652, 26651, 26650, 6062
# validator4 1314, 9084, 9085, 26649, 26648, 26647, 6063

# change app.toml values

# validator2
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1316|g' $HOME/.elestod/validator2/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9088|g' $HOME/.elestod/validator2/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9089|g' $HOME/.elestod/validator2/config/app.toml

# validator3
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:1315|g' $HOME/.elestod/validator3/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9086|g' $HOME/.elestod/validator3/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9087|g' $HOME/.elestod/validator3/config/app.toml


# change config.toml values

# validator1
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.elestod/validator1/config/config.toml
# validator2
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26655|g' $HOME/.elestod/validator2/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26654|g' $HOME/.elestod/validator2/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26653|g' $HOME/.elestod/validator2/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.elestod/validator3/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.elestod/validator2/config/config.toml
# validator3
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:26652|g' $HOME/.elestod/validator3/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:26651|g' $HOME/.elestod/validator3/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.elestod/validator3/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:26650|g' $HOME/.elestod/validator3/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.elestod/validator3/config/config.toml


# copy validator1 genesis file to validator2-3
cp $HOME/.elestod/validator1/config/genesis.json $HOME/.elestod/validator2/config/genesis.json
cp $HOME/.elestod/validator1/config/genesis.json $HOME/.elestod/validator3/config/genesis.json


# copy tendermint node id of validator1 to persistent peers of validator2-3
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(elestod tendermint show-node-id --home=$HOME/.elestod/validator1)@localhost:26656\"|g" $HOME/.elestod/validator2/config/config.toml
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(elestod tendermint show-node-id --home=$HOME/.elestod/validator1)@localhost:26656\"|g" $HOME/.elestod/validator3/config/config.toml

# start all three validators
tmux new -s validator1 -d elestod start --home=$HOME/.elestod/validator1
tmux new -s validator2 -d elestod start --home=$HOME/.elestod/validator2
tmux new -s validator3 -d elestod start --home=$HOME/.elestod/validator3


# send stake from first validator to second validator
sleep 7
elestod tx bank send validator1 $(elestod keys show validator2 -a --keyring-backend=test --home=$HOME/.elestod/validator2) 500000000stake --keyring-backend=test --home=$HOME/.elestod/validator1 --chain-id=testing --yes
sleep 7
elestod tx bank send validator1 $(elestod keys show validator3 -a --keyring-backend=test --home=$HOME/.elestod/validator3) 400000000stake --keyring-backend=test --home=$HOME/.elestod/validator1 --chain-id=testing --yes

## create second validator
sleep 7
elestod tx staking create-validator --amount=500000000stake --from=validator2 --pubkey=$(elestod tendermint show-validator --home=$HOME/.elestod/validator2) --moniker="validator2" --chain-id="testing" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="500000000" --keyring-backend=test --home=$HOME/.elestod/validator2 --yes
sleep 7
elestod tx staking create-validator --amount=400000000stake --from=validator3 --pubkey=$(elestod tendermint show-validator --home=$HOME/.elestod/validator3) --moniker="validator3" --chain-id="testing" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="400000000" --keyring-backend=test --home=$HOME/.elestod/validator3 --yes
