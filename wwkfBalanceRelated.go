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

type request struct {
	To   string `json:"to"`
	Data string `json:"data"`
}

/*
- `byteArr := []byte(stringObj)`
- `sha3.NewLegacyKeccak256`
*/
func contractUtilObtainMethodId(strFnSignature string) []byte {
	byteFnSignature := []byte(strFnSignature)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(byteFnSignature)
	methodId := hash.Sum(nil)[:4]
	return methodId
}

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

func transferTokenWithAmount(contractAddr string, pkFrom string, addrTo string, tokenAmount int64) (bool, string) {
	keyPair, err := crypto.HexToECDSA(pkFrom)
	if err != nil {
		log.Fatal("Fail to convert pkFrom into ECDSA: ", err)
		return false, ""
	}
	byteAddrFrom := crypto.PubkeyToAddress(keyPair.PublicKey)

	byteTokenAddr := common.HexToAddress(contractAddr)
	byteAddrTo := common.HexToAddress(addrTo)
	methodId := contractUtilObtainMethodId("transfer(address,uint256)")
	paddedByteAddrTo := common.LeftPadBytes(byteAddrTo.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedByteAddrTo))
	bigIntTokenAmount := big.NewInt(tokenAmount)
	paddedByteTokenAmount := common.LeftPadBytes(bigIntTokenAmount.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedByteTokenAmount))

	nonce, err := client.PendingNonceAt(ctx, byteAddrFrom)
	if err != nil {
		log.Fatal("Fail to pend nonce at pkFrom: ", err)
		return false, ""
	}

	ethAmount := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal("Fail to have suggested gas price: ", err)
	}

	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedByteAddrTo...)
	data = append(data, paddedByteTokenAmount...)
	gasLimit := contractUtilSuggestGas(byteAddrFrom, byteTokenAddr, data)
	// fmt.Println(gasLimit)

	// 👉 to be refactored (`transferEtherWithAmount`)
	tx := types.NewTransaction(nonce, byteTokenAddr, ethAmount, gasLimit, gasPrice, data)
	chainId, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal("Fail to obtain NetworkID: ", err)
		return false, ""
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), keyPair)
	if err != nil {
		log.Fatal("Fail to sign transection: ", err)
		return false, ""
	}
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal("Fail to send transection: ", err)
		return false, ""
	}

	return true, signedTx.Hash().Hex()
}
