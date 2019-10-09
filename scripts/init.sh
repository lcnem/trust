#!/bin/sh

echo "Enter key-name"
read key_name
trustcli keys add $key_name

echo "Enter moniker"
read moniker

echo "Enter chain-id (If left empty, randomly created)"
read chain_id
trustd init $moniker --chain-id $chain_id

trustd add-genesis-account $key_name 1000000000stake
valconspub=$(trustd tendermint show-validator)
trustd gentx --amount 1000000stake --name $key_name --pubkey $valconspub
trustd collect-gentxs
