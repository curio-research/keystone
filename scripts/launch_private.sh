#!/bin/sh

# Launch Private network with one node
make
./build/bin/geth --networkid 13447 --datadir ./data --port 8808 --rpc.gascap 0 --rpc.txfeecap 0 --http --http.corsdomain="*" --http.api web3,eth,debug,personal,net,miner --http.port 8540  --vmdebug --unlock 0x697E9295D39C6914E4e1cf73FD8fA6490adA08b9 --password ./pass-private.txt --mine --miner.etherbase 0x697E9295D39C6914E4e1cf73FD8fA6490adA08b9 --allow-insecure-unlock --nodiscover