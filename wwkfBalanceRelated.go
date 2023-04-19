package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
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

func accessTokenBalance(contractAddr string, userAddr string) *big.Int {
	methodId := hexutil.Encode(contractUtilObtainMethodId("balanceOf(address)"))
	// %064s: padded with 0 to 64 bytes
	data := methodId + fmt.Sprintf("%064s", userAddr[2:])
	client, err := rpc.DialHTTP(W3NET_URL)
	if err != nil {
		log.Fatal("Fail to dial to W3NET_URL using rpc protocal: ", err)
	}
	defer client.Close()

	req := request{contractAddr, data}
	var resp string
	if err := client.Call(&resp, "eth_call", req, "latest"); err != nil {
		log.Fatal("Fail to call web3 API: ", err)
	}

	// remove leading zero of resp
	// %064s: the string is padded with 0 to 64 bytes
	formattedResp := strings.TrimLeft(resp[2:], "0")
	if len(formattedResp) == 0 {
		formattedResp = "0"
	}
	balance, err := hexutil.DecodeBig("0x" + formattedResp)
	if err != nil {
		log.Fatal("Fail to decode 'formattedResp' into 'big.Int': ", err)
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
