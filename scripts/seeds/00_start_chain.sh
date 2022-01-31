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
    echo "y" | elestod keys add bob
    echo "y" | elestod keys add alice
 
    echo "Adding genesis account"
    elestod add-genesis-account $(elestod keys show validator -a) 1000000000stake --
    # this is to have the accounts on chain 
    elestod add-genesis-account $(elestod keys show emti -a) 1000stake
    elestod add-genesis-account $(elestod keys show arti -a) 1000stake
    elestod add-genesis-account $(elestod keys show bob -a) 1000stake
    elestod add-genesis-account $(elestod keys show alice -a) 1000stake
    ## add the regulator
    elestod add-genesis-account $(elestod keys show regulator -a) 1000stake --regulator $(elestod keys show regulator -a) --
    elestod gentx validator 700000000stake --chain-id elesto
    elestod collect-gentxs
fi


echo "Starting Elesto chain"
elestod start
