#!/bin/sh

# build and launch local devnet
echo 👾 Starting up ...

make
./build/bin/geth --http --http.corsdomain="*" --http.api web3,eth,debug,personal,net --vmdebug --dev --http.port 8540 --miner.gaslimit 100000000000000 --rpc.gascap 0 --dev.gaslimit 100000000000000 --miner.gasprice 0 --rpc.txfeecap 0 --dev.period 1

