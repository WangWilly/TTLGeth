package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

/*
Deprecated.
*/
//lint:ignore U1000 Ignore unused function temporarily for debugging
type request struct {
	To   string `json:"to"`
	Data string `json:"data"`
}

/*
Deprecated.
- `byteArr := []byte(stringObj)`
- `sha3.NewLegacyKeccak256`
*/
//lint:ignore U1000 Ignore unused function temporarily for debugging
func contractUtilObtainMethodId(strFnSignature string) []byte {
	byteFnSignature := []byte(strFnSignature)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(byteFnSignature)
	methodId := hash.Sum(nil)[:4]
	return methodId
}

/*
Deprecated.
*/
//lint:ignore U1000 Ignore unused function temporarily for debugging
func contractUtilSuggestGas(byteAddrFrom common.Address, byteTokenAddr common.Address, data []byte) uint64 {
	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:  byteAddrFrom,
		To:    &byteTokenAddr,
		Value: big.NewInt(0),
		Data:  data,
	})
	if err != nil {
		log.Println("Fail to estimate gas price: ", err)
		return uint64(1000000)
	}
	return gasLimit
}

func getWwkfBalance(userAddr string) *big.Int {
	balance, err := instance.BalanceOf(
		&bind.CallOpts{},
		common.HexToAddress(userAddr),
	)
	if err != nil {
		log.Panic("Fail to access the balance of the WWKF: ", err)
	}
	return balance
}

func transferWwkfWithAmount(pkFrom string, addrTo string, tokenAmount int64) string {
	authedTransector := getAuthedTransector(pkFrom)
	reply, err := instance.Transfer(
		authedTransector,
		common.HexToAddress(addrTo),
		big.NewInt(tokenAmount),
	)
	if err != nil {
		log.Panic("Fail to transfer WWKF token: ", err)
	}
	return reply.Hash().Hex()
}

func getAuthedTransector(pk string) *bind.TransactOpts {
	keyPair, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Panic("Fail to transform pk string into ECDSA: ", err)
	}
	byteAddr := crypto.PubkeyToAddress(keyPair.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, byteAddr)
	if err != nil {
		log.Panic("Fail to pend the nonce: ", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Panic("Fail to obain the current chainID: ", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(keyPair, chainID)
	if err != nil {
		log.Panic("Fail to create a new keyed transection: ", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)          // in wei
	auth.GasLimit = uint64(3000000)     // in units. TODO: require improvment
	auth.GasPrice = big.NewInt(1000000) // in wei. TODO: require improvment

	return auth
}
