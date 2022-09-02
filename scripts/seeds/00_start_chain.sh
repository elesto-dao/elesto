#!/bin/bash

GENESIS_FILE=~/.elesto/config/genesis.json
if [ -f $GENESIS_FILE ]
then
    echo "Genesis file exist, would you like to delete it? (y/n)"
    read delete_config
fi

if [[
	$delete_config == "Y" ||
	$delete_config == "y" ||
	! -f $GENESIS_FILE
   ]];
then
    rm -r ~/.elesto

    echo "Initialising chain"
    elestod init --chain-id=elesto elesto

    echo "Preparing genesis"
    elestod prepare-genesis testnet elesto

    echo "Adding local keys"
    echo "y" | elestod keys add validator
    echo "y" | elestod keys add regulator
    echo "y" | elestod keys add emti 
    echo "y" | elestod keys add arti 
    echo "video adult rule exhaust tube crater lunch route clap pudding poet pencil razor pluck veteran hill stock thunder sense riot fox oppose glare bar" | elestod keys add bob --recover --keyring-backend test
    echo "y" | elestod keys add alice

    echo "Adding genesis accounts"
    elestod add-genesis-account $(elestod keys show validator -a) 6000000000000uelesto
    elestod add-genesis-account $(elestod keys show emti -a) 5000000000000uelesto
    elestod add-genesis-account $(elestod keys show arti -a) 5000000000000uelesto
    elestod add-genesis-account $(elestod keys show bob -a --keyring-backend test) 1000000000000uelesto
    elestod add-genesis-account $(elestod keys show alice -a) 1000000000000uelesto
    elestod add-genesis-account $(elestod keys show regulator -a) 2000000000000uelesto

    echo "Creating validator genesis transaction"
    elestod gentx validator 5000000000000uelesto --chain-id elesto

    echo "Adding validator genesis transaction to genesis file"
    elestod collect-gentxs

    # the community tax in the distribution must be disabled, since the community tax is
    # already distributed by the mint module
    echo "$( jq '.app_state.distribution.params.community_tax = "0.1"' ~/.elesto/config/genesis.json )" > ~/.elesto/config/genesis.json

    echo
fi


echo "Starting Elesto chain"
elestod start
