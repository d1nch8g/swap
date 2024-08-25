package main

import (
	"context"
	"encoding/base64"
	"log"
	"os"
	"strings"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

func main() {
	client := liteclient.NewConnectionPool()

	cfg, err := liteclient.GetConfigFromUrl(context.Background(), "https://ton.org/testnet-global.config.json")
	if err != nil {

	}

	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		log.Fatalln("connection err: ", err.Error())
		return
	}

	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(cfg)

	// bound all requests to single ton node
	ctx := client.StickyContext(context.Background())

	var seed []string
	if f, err := os.ReadFile("seed.key"); err != nil {
		log.Println("error: ", err)
		seed = wallet.NewSeed()
		err = os.WriteFile("seed.key", []byte(strings.Join(seed, " ")), os.ModePerm)
		if err != nil {
			log.Fatalln("Save seed err:", seed)
			return
		}
	} else {
		seed = strings.Split(string(f), " ")
	}

	w, err := wallet.FromSeed(api, seed, wallet.V4R1)
	if err != nil {
		log.Fatalln("FromSeed err:", err.Error())
		return
	}

	addr := w.WalletAddress()
	err = os.WriteFile("address.key", []byte(addr.String()), os.ModePerm)
	if err != nil {
		log.Fatalln("Save seed err:", seed)
		return
	}

	log.Println("wallet address:", w.WalletAddress())

	log.Println("fetching and checking proofs since config init block, it may take near a minute...")
	block, err := api.CurrentMasterchainInfo(context.Background())
	if err != nil {
		log.Fatalln("get masterchain info err: ", err.Error())
		return
	}
	log.Println("master proof checks are completed successfully, now communication is 100% safe!")

	balance, err := w.GetBalance(ctx, block)

	log.Println("Balance: ", balance, err)

	addr = address.MustParseAddr("UQAGCpx5L_noxcmqAD66VFiDhynHIBnkpE2--ZYm3RRU7qtB")

	transfer, err := w.BuildTransfer(addr, tlb.MustFromTON("0.003"), true, "Hello from tonutils-go!")
	if err != nil {
		log.Fatalln("Transfer err:", err.Error())
		return
	}

	tx, block, err := w.SendWaitTransaction(ctx, transfer)
	if err != nil {
		log.Fatalln("SendWaitTransaction err:", err.Error())
		return
	}

	balance, err = w.GetBalance(ctx, block)
	if err != nil {
		log.Fatalln("GetBalance err:", err.Error())
		return
	}

	log.Printf("transaction confirmed at block %d, hash: %s balance left: %s", block.SeqNo,
		base64.StdEncoding.EncodeToString(tx.Hash), balance.String())
}
