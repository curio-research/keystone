#!/bin/sh

# Remove all block data locally and re-generate genesis block again
make
rm -rf ./data/geth
./build/bin/geth --datadir ./data init './genesis_private.json'