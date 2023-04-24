package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	wwkf "ttlGeth/bindings/wwkf"

	_ "github.com/joho/godotenv/autoload"
)

// Cardinal setting
var (
	NET_BRANCH_NAME    = "sepolia"
	INFURA_PROJECT_ID  = os.Getenv("INFURA_PROJECT_ID")
	PRIVATE_KEY_1      = os.Getenv("NO0x_PRIVATE_KEY_1")
	PUBLIC_KEY_1       = os.Getenv("PUBLIC_KEY_1")
	WALLET_ADDR_1      = os.Getenv("WALLET_HASH_1")
	PRIVATE_KEY_2      = os.Getenv("NO0x_PRIVATE_KEY_2")
	PUBLIC_KEY_2       = os.Getenv("PUBLIC_KEY_2")
	WALLET_ADDR_2      = os.Getenv("WALLET_HASH_2")
	WWKF_CONTRACT_ADDR = "0x03378DAa43739f2361FE67175aD6bF2666309748"
	W3NET_URL          = "https://" + NET_BRANCH_NAME + ".infura.io/v3/" + INFURA_PROJECT_ID
	W3WSS_URL          = "wss://" + NET_BRANCH_NAME + ".infura.io/ws/v3/" + INFURA_PROJECT_ID
)

// Ethereum network: mainnet, sepolia
// Available testnets: Sepolia, Goerli, Ropsten, Rinkeby and Kovan
var (
	ctx                   = context.Background()
	client, dialErr       = ethclient.DialContext(ctx, W3NET_URL)
	instance, contractErr = wwkf.NewWwkf(common.HexToAddress(WWKF_CONTRACT_ADDR), client)
	wssClient, wssDialErr = ethclient.DialContext(ctx, W3WSS_URL)
)

/*
🙋‍♂️: make sure the connection to the service is safe and sound.

Access Ethereum network (with Infura help, users don't need to
maintain a Ethereum node locally) and checkout the latest
block number.
*/
//lint:ignore U1000 Ignore unused function temporarily for debugging
func currentBlock() *big.Int {
	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("Fail to obtain current block number from ", NET_BRANCH_NAME, ": ", err)
	}
	return block.Number()
}

/*
Wallets are composed of three main components; the public key,
the private key,and the public address.
*/
//lint:ignore U1000 Ignore unused function temporarily for debugging
func createWallet() (string, string, string) {
	keyPair, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("Fail to generate a private key for a new wallet: ", err)
	}

	// Transform the byte information into hexdecimal string (for readibility)
	privateKey := crypto.FromECDSA(keyPair)
	hexStrPrivateKey := hexutil.Encode(privateKey)
	publicKey := crypto.FromECDSAPub(&keyPair.PublicKey)
	hexStrPublicKey := hexutil.Encode(publicKey)
	thePublicAddress := crypto.PubkeyToAddress(keyPair.PublicKey).Hex()
	return thePublicAddress, hexStrPublicKey, hexStrPrivateKey
} // 📌

func checkWalletBalance(addr string) *big.Int {
	byteAddr := common.HexToAddress(addr)
	// Get the balance of the given 'addr' at the current block
	balance, err := client.BalanceAt(ctx, byteAddr, nil)
	if err != nil {
		log.Fatal("Fail to check eth balance of ", addr, ": ", err)
	}
	return balance
}

func transferEtherWithAmount(pkFrom string, addrTo string, amountInWei int64) string {
	// Transform hexdecimal string into byte object (*ecdsa.PrivateKey)
	keyPair, err := crypto.HexToECDSA(pkFrom)
	if err != nil {
		log.Panic("Fail to recover '*ecdsa.PrivateKey' from the given private key: ", err)
	}
	byteAddrFrom := crypto.PubkeyToAddress(keyPair.PublicKey)
	byteAddrTo := common.HexToAddress(addrTo)

	// nonce: the number of transactions from the address
	nonce, err := client.PendingNonceAt(ctx, byteAddrFrom)
	if err != nil {
		log.Panic("Fail to pend the 'nonce' at 'addrFrom': ", err)
	}

	amount := big.NewInt(amountInWei)
	gasLimit := uint64(21000) // 🤔: intrinsic gas too low
	gas, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Panic("Fail to obtain optimal gas price: ", err)
	}

	// chainId: https://openethereum.github.io/Chain-specification
	chainId, err := client.NetworkID(ctx)
	if err != nil {
		log.Panic("Fail to obtain current network (", NET_BRANCH_NAME, ") id: ", err)
	}

	transaction := types.NewTransaction(nonce, byteAddrTo, amount, gasLimit, gas, nil)
	signedTx, err := types.SignTx(transaction, types.NewEIP155Signer(chainId), keyPair)
	if err != nil {
		log.Panic("Fail to sign transection:", err)
	}
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Panic("Fail to send transection: ", err)
	}

	return signedTx.Hash().Hex()
} // 📌

/*
- ref1: https://golangbot.com/channels/

- ref2: https://blog.wu-boy.com/2020/05/understant-golang-context-in-10-minutes/
*/
func waitForTxCompletion(txStr string, retStatus chan<- bool, maxWaitInSec int) {
	isExpired := make(chan bool)

	// create a new channel that will receive the latest block headers
	headers := make(chan *types.Header)

	// call the SubscribeNewHead method with the headers channel
	sub, err := wssClient.SubscribeNewHead(ctx, headers)
	if err != nil {
		log.Panic("Fail to subscribe new head: ", err)
	}

	// use a select statement to listen for new messages or errors
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Panic("Error occurs in the subscribtion: ", err)
			case header := <-headers:
				// get the full block by passing the header hash to BlockByHash function
				block, err := wssClient.BlockByHash(ctx, header.Hash())
				if err != nil {
					log.Fatal(err)
				}
				// fmt.Println(block.Hash().Hex())
				// fmt.Println(block.Number().Uint64())
				// fmt.Println(block.Nonce())

				// loop through the transactions in the block and check if any of them matches your transaction hash
				for _, tx := range block.Transactions() {
					if tx.Hash().Hex() == txStr {
						// do something with your transaction
						retStatus <- true
					}
				}
			case <-isExpired:
				log.Println("WaitForTxCompletion is expired in", maxWaitInSec, "secs.")
				retStatus <- false
			}
		}
	}()

	time.Sleep(time.Duration(maxWaitInSec) * time.Second)
	isExpired <- true
}

/*
sepolia faucet: https://www.infura.io/faucet
- common in Go
  - `log.Println`

- common in address process
  - `common.HexToAddress`

- common in Geth
  - `client.BalanceAt`
*/
func main() {
	// [Initialization]
	if dialErr != nil {
		log.Panic(dialErr)
	}
	if contractErr != nil {
		log.Panic(contractErr)
	}
	if wssDialErr != nil {
		log.Panic(wssDialErr)
	}

	// 📌✅ [Start]
	// currentBlock()

	// 📌✅ [Create wallet]
	// addr, pubKey, privKey := createWallet()
	// fmt.Println("addr:", addr, "\npubKey:", pubKey, "\nprivKey:", privKey)

	// 📌✅ [Ethereum transection]
	fmt.Println("- Before transection:")
	fmt.Println("Ether balance of", WALLET_ADDR_1, ":", checkWalletBalance(WALLET_ADDR_1))
	fmt.Println("Ether balance of", WALLET_ADDR_2, ":", checkWalletBalance(WALLET_ADDR_2))
	signedTxStr := transferEtherWithAmount(
		PRIVATE_KEY_1,
		WALLET_ADDR_2,
		1000000000,
	)
	fmt.Println("complete transfering WWKF:", signedTxStr)
	retStatus := make(chan bool)
	fmt.Println("wait for the TX getting approved:", signedTxStr)
	go waitForTxCompletion(signedTxStr, retStatus, 60)
	txStatus := <-retStatus
	if !txStatus {
		fmt.Println("fail to wait for the TX getting approved")
	}
	fmt.Println("- After transection:")
	fmt.Println("Ether balance of", WALLET_ADDR_1, ":", checkWalletBalance(WALLET_ADDR_1))
	fmt.Println("Ether balance of", WALLET_ADDR_2, ":", checkWalletBalance(WALLET_ADDR_2))

	// 📌✅ [Transfer ERC20 token (WWKF)]
	fmt.Println("- Before WWKF transection:")
	fmt.Println(
		"WWKF balance of",
		WALLET_ADDR_1,
		":",
		getWwkfBalance(WALLET_ADDR_1),
	)
	fmt.Println(
		"WWKF balance of",
		WALLET_ADDR_2,
		":",
		getWwkfBalance(WALLET_ADDR_2),
	)
	signedTxStr = transferWwkfWithAmount(
		PRIVATE_KEY_1,
		WALLET_ADDR_2,
		int64(5000),
	)
	fmt.Println("complete transfering WWKF:", signedTxStr)
	retStatus = make(chan bool)
	fmt.Println("wait for the TX getting approved:", signedTxStr)
	go waitForTxCompletion(signedTxStr, retStatus, 60)
	txStatus = <-retStatus
	if !txStatus {
		fmt.Println("fail to wait for the TX getting approved")
	}
	fmt.Println("- After WWKF transection:")
	fmt.Println("Ether balance of", WALLET_ADDR_1, ":", checkWalletBalance(WALLET_ADDR_1))
	fmt.Println("Ether balance of", WALLET_ADDR_2, ":", checkWalletBalance(WALLET_ADDR_2))
	fmt.Println(
		"WWKF balance of",
		WALLET_ADDR_1,
		":",
		getWwkfBalance(WALLET_ADDR_1),
	)
	fmt.Println(
		"WWKF balance of",
		WALLET_ADDR_2,
		":",
		getWwkfBalance(WALLET_ADDR_2),
	)

}
