package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

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
	BUILT_WWKF_PATH    = "./contracts/willywangkaaFirstContract.json"
	W3NET_URL          = "https://" + NET_BRANCH_NAME + ".infura.io/v3/" + INFURA_PROJECT_ID
)

// Ethereum network: mainnet, sepolia
// Available testnets: Sepolia, Goerli, Ropsten, Rinkeby and Kovan
var (
	ctx         = context.Background()
	client, err = ethclient.DialContext(ctx, W3NET_URL)
)

/*
🙋‍♂️: make sure the connection to the service is safe and sound.

Access Ethereum network (with Infura help, users don't need to
maintain a Ethereum node locally) and checkout the latest
block number.
*/
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

func transferEtherWithAmount(pkFrom string, addrTo string, amountInWei int64) (bool, string) {
	// Transform hexdecimal string into byte object (*ecdsa.PrivateKey)
	keyPair, err := crypto.HexToECDSA(pkFrom)
	if err != nil {
		log.Fatal("Fail to recover '*ecdsa.PrivateKey' from the given private key: ", err)
	}
	byteAddrFrom := crypto.PubkeyToAddress(keyPair.PublicKey)
	byteAddrTo := common.HexToAddress(addrTo)

	// nonce: the number of transactions from the address
	nonce, err := client.PendingNonceAt(ctx, byteAddrFrom)
	if err != nil {
		log.Fatal("Fail to pend the 'nonce' at 'addrFrom': ", err)
	}

	amount := big.NewInt(amountInWei)
	gasLimit := uint64(21000) // 🤔: intrinsic gas too low
	gas, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal("Fail to obtain optimal gas price: ", err)
	}

	// chainId: https://openethereum.github.io/Chain-specification
	chainId, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal("Fail to obtain current network (", NET_BRANCH_NAME, ") id: ", err)
	}

	transaction := types.NewTransaction(nonce, byteAddrTo, amount, gasLimit, gas, nil)
	signedTx, err := types.SignTx(transaction, types.NewEIP155Signer(chainId), keyPair)
	if err != nil {
		log.Fatal("Fail to sign transection:", err)
	}
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatal("Fail to send transection: ", err)
	}

	return true, signedTx.Hash().Hex()
} // 📌

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
	if err != nil { // check `ethclient.DialContext`
		log.Println(err)
	}

	// 📌✅ [Start]
	// currentBlock()

	// 📌✅ [Create wallet]
	// addr, pubKey, privKey := createWallet()
	// fmt.Println("addr:", addr, "\npubKey:", pubKey, "\nprivKey:", privKey)

	// 📌✅ [Ethereum transection]
	// fmt.Println("Before transection:")
	// fmt.Println("Ether balance of", WALLET_ADDR_1, ":", checkWalletBalance(WALLET_ADDR_1))
	// fmt.Println("Ether balance of", WALLET_ADDR_2, ":", checkWalletBalance(WALLET_ADDR_2))
	// status, signedTxStr := transferEtherWithAmount(
	// 	PRIVATE_KEY_1,
	// 	WALLET_ADDR_2,
	// 	100000000000000000,
	// )
	// if status {
	// 	fmt.Println("transaction sent:", signedTxStr)
	// }
	// fmt.Println("After transection:")
	fmt.Println("Ether balance of", WALLET_ADDR_1, ":", checkWalletBalance(WALLET_ADDR_1))
	fmt.Println("Ether balance of", WALLET_ADDR_2, ":", checkWalletBalance(WALLET_ADDR_2))

	// 📌✅ [Transfer ERC20 token (WWKF)]
	// check balance first
	// fmt.Println("Before WWKF transection:")
	// fmt.Println(
	// 	"WWKF balance of",
	// 	WALLET_ADDR_1,
	// 	":",
	// 	accessTokenBalance(WWKF_CONTRACT_ADDR, WALLET_ADDR_1),
	// )
	// fmt.Println(
	// 	"WWKF balance of",
	// 	WALLET_ADDR_2,
	// 	":",
	// 	accessTokenBalance(WWKF_CONTRACT_ADDR, WALLET_ADDR_2),
	// )
	// status, signedTxStr := transferTokenWithAmount(
	// 	WWKF_CONTRACT_ADDR,
	// 	PRIVATE_KEY_2,
	// 	WALLET_ADDR_1,
	// 	500000000,
	// )
	// if status {
	// 	fmt.Println("complete transfering WWKF:", signedTxStr)
	// }
	// fmt.Println("After WWKF transection:")
	// fmt.Println("Ether balance of", WALLET_ADDR_1, ":", checkWalletBalance(WALLET_ADDR_1))
	// fmt.Println("Ether balance of", WALLET_ADDR_2, ":", checkWalletBalance(WALLET_ADDR_2))
	fmt.Println(
		"WWKF balance of",
		WALLET_ADDR_1,
		":",
		accessTokenBalance(WWKF_CONTRACT_ADDR, WALLET_ADDR_1),
	)
	fmt.Println(
		"WWKF balance of",
		WALLET_ADDR_2,
		":",
		accessTokenBalance(WWKF_CONTRACT_ADDR, WALLET_ADDR_2),
	)

}
