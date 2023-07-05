package testutil

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/scinterface/lagrange"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	optintoSig = "0xf73b7519"
)

// RegisterOperator registers an operator to the lagrange service.
func RegisterOperator(client *ethclient.Client, auth *bind.TransactOpts, stakeAddr, slasherAddr common.Address) error {
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return err
	}
	gasLimit := uint64(3000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	addrHex := stakeAddr.Hash()
	callData := append(common.FromHex(optintoSig), addrHex[:]...)

	tx := types.NewTransaction(
		nonce,
		slasherAddr,
		big.NewInt(0),
		gasLimit,
		gasPrice,
		callData,
	)
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return err
	}
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		return err
	}

	time.Sleep(3 * time.Second)
	service, err := lagrange.NewLagrange(stakeAddr, client)
	if err != nil {
		return err
	}
	_, err = service.Register(auth, big.NewInt(0), []byte{}, math.MaxUint32)
	return err
}
