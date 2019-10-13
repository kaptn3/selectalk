docker run --rm \
-v ~/go/src/github.com/rodsher/selectel/deployments/ethereum/genesis.json:/root/genesis.json \
-v ~/.ethereum:/root/.ethereum \
ethereum/client-go --datadir /root/.ethereum init /root/genesis.json