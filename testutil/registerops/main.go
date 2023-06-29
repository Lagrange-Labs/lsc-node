package main

import (
	"context"
	"fmt"

	"github.com/Lagrange-Labs/lagrange-node/testutil"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	StakeAddress   = common.HexToAddress("0xf32358f5C8FFfCF1a7bDb58b270a082abb7Ba1A6")
	SlasherAddress = common.HexToAddress("0x6Bf0fF4eBa00E3668c0241bb1C622CDBFE55bbE0")
	PrivateKeys    = []string{
		"0x232d99bc62cf95c358fb496e9f820ec299f43417397cea32f9f365daf4748429",
		"0x25f536330df3a72fa381bfb5ea5552b2731523f08580e7a0e2e69618a9643faa",
		"0xc262364335471942e02e79d760d1f5c5ad7a34463303851cacdd15d72e68b228",
	}
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(fmt.Errorf("failed to connect to ethereum node: %v", err))
	}

	for _, privateKey := range PrivateKeys {
		auth, err := utils.GetSigner(context.Background(), client, privateKey)
		if err != nil {
			panic(fmt.Errorf("failed to get signer: %v", err))
		}
		err = testutil.RegisterOperator(client, auth, StakeAddress, SlasherAddress)
		if err != nil {
			panic(fmt.Errorf("failed to register operator: %v", err))
		}
	}
}
