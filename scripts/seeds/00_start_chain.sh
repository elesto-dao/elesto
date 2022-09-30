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
    echo "y" | elestod keys add validator
    echo "y" | elestod keys add regulator
    echo "y" | elestod keys add emti # e-money token issuer
    echo "y" | elestod keys add arti # asset-referenced token issuer
    echo "video adult rule exhaust tube crater lunch route clap pudding poet pencil razor pluck veteran hill stock thunder sense riot fox oppose glare bar" | elestod keys add bob --recover --keyring-backend test
    echo "y" | elestod keys add alice

    echo "Adding genesis account"
    elestod add-genesis-account $(elestod keys show validator -a) 40000000000000stake
    # this is to have the accounts on chain
    elestod add-genesis-account $(elestod keys show emti -a) 40000000000000stake
    elestod add-genesis-account $(elestod keys show arti -a) 40000000000000stake
    elestod add-genesis-account $(elestod keys show bob -a) 20000000000000stake
    elestod add-genesis-account $(elestod keys show alice -a) 20000000000000stake
    ## add the regulator
    elestod add-genesis-account $(elestod keys show regulator -a) 40000000000000stake
    elestod gentx validator 30000000000000stake --chain-id elesto
    elestod collect-gentxs

    # the community tax in the distribution must be disabled, since the community tax is
    # already distributed by the mint module
    echo "$( jq '.app_state.distribution.params.community_tax = "0.1"' ~/.elesto/config/genesis.json )" > ~/.elesto/config/genesis.json

    echo
fi


echo "Starting Elesto chain"
elestod start
