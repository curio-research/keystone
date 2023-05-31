#!/bin/bash

if [ -n "$1" ]
then
  case "$1" in
    --prod)
      # Prod Env
      echo "Running in Prod Mode"
      geth --datadir ./data init './genesis_private.json'
      geth --networkid 13447 --datadir ./data --port 8808 --rpc.gascap 0 --rpc.txfeecap 0 --http --http.corsdomain="*" --http.api web3,eth,debug,personal,net,miner --http.addr 0.0.0.0 --vmdebug --unlock 0x697E9295D39C6914E4e1cf73FD8fA6490adA08b9 --password ./pass-private.txt --mine --miner.etherbase 0x697E9295D39C6914E4e1cf73FD8fA6490adA08b9 --allow-insecure-unlock --nodiscover
      ;;
    *) echo "$1 is not an option" ;;
  esac
else
  # Default: Dev Env
  echo "Running in Dev Mode"
  geth --vmdebug --dev --dev.gaslimit 100000000000000 --dev.period 1 --http --http.corsdomain="*" --http.api web3,eth,debug,personal,net --http.addr 0.0.0.0 --rpc.gascap 0 --rpc.txfeecap 0 --miner.gaslimit 100000000000000 --miner.gasprice 0
fi
