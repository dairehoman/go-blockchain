package main

import (
	"flag"
	"log"
	"time"

	"../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func addBlock(client proto.BlockchainClient) {

	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatal("unable to add block : ", block.Hash)
	}
}

func getBlockchain(client proto.BlockchainClient) {
	bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatal("unable to get blockchain data: ", err)
	}

	log.Println("blocks:")
	for _, block := range bc.Blocks {
		log.Printf("hash: %s, prev block hash: %s, data: %s", block.Hash, block.PrevBlockHash, block.Data)
	}

}

func main() {
	add := flag.Bool("add", false, "add new block")
	get := flag.Bool("get", false, "get blockchain data")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial into server: ", err)
	}

	client := proto.NewBlockchainClient(conn)

	if *add {
		addBlock(client)
	}

	if *get {
		getBlockchain(client)
	}
}
