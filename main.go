package main

import (
	"context"

	"github.com/cedricfung/gao/lottery"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	GethRPC = "http://104.197.245.214:8545"
)

func main() {
	conn, err := ethclient.Dial(GethRPC)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	worker := lottery.NewWorker(ctx, conn)
	worker.Loop(ctx)
}
