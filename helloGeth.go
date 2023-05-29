package main

import (
	"context"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	wwkf "ttlGeth/bindings/wwkf"
	C "ttlGeth/constants/paths"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-gonic/gin"
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
						return
					}
				}
			case <-isExpired:
				log.Println("WaitForTxCompletion is expired in", maxWaitInSec, "secs.")
				retStatus <- false
				return
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
	// initialization
	if dialErr != nil {
		log.Panic(dialErr)
	}
	if contractErr != nil {
		log.Panic(contractErr)
	}
	if wssDialErr != nil {
		log.Panic(wssDialErr)
	}

	api := gin.Default()
	api.SetTrustedProxies(nil)
	api.StaticFile("/favicon.ico", "./ttlgeth-frontend/public/wwkfIcon.svg")

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Welcome": "hello geth"})
	})

	api.GET(C.P_CHECK_ETH+C.PARAM_WALLETADDR, func(c *gin.Context) {
		walletAddr := C.GetDefaultAddr(c.Param(C.WALLET_ADDR))
		c.JSON(http.StatusOK, gin.H{"ethBalance": checkWalletBalance(walletAddr)})
	})

	api.GET(C.P_CHECK_WWKF+C.PARAM_WALLETADDR, func(c *gin.Context) {
		walletAddr := C.GetDefaultAddr(c.Param(C.WALLET_ADDR))
		c.JSON(http.StatusOK, gin.H{"wwkfBalance": getWwkfBalance(walletAddr)})
	})

	api.GET(C.P_TRANSFER_ETH, func(c *gin.Context) {
		addrFrom := C.GetDefaultAddr(c.Query("addrFrom"))
		pkFrom := C.GetDefaultPk(addrFrom)
		addrTo := C.GetDefaultAddr(c.Query("addrTo"))
		amount, err := strconv.ParseInt(c.DefaultQuery("amount", "1000000000"), 10, 64)
		if err != nil {
			errMsg := "Fail to parse amount into a integer: " + err.Error()
			log.Panic(errMsg)
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": errMsg})
		}
		prevAddrfromBalance, prevAddrtoBalance := checkWalletBalance(addrFrom), checkWalletBalance(addrTo)
		signedTxStr := transferEtherWithAmount(pkFrom, addrTo, amount)
		retStatus := make(chan bool)
		go waitForTxCompletion(signedTxStr, retStatus, 60)
		txStatus := <-retStatus
		if !txStatus {
			errMsg := "Fail to wait for the TX getting approved"
			log.Panic(errMsg)
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		}

		currAddrfromBalance, currAddrtoBalance := checkWalletBalance(addrFrom), checkWalletBalance(addrTo)
		c.JSON(
			http.StatusOK,
			gin.H{
				"txAddr":                     signedTxStr,
				"addrfrom":                   addrFrom,
				"addrto":                     addrTo,
				"amount":                     amount,
				"previousAddrfromEthBalance": prevAddrfromBalance,
				"previousAddrtoEthBalance":   prevAddrtoBalance,
				"currentAddrfromEthBalance":  currAddrfromBalance,
				"currentAddrtoEthBalance":    currAddrtoBalance,
			},
		)
	})

	api.GET(C.P_TRANSFER_WWKF, func(c *gin.Context) {
		addrFrom := C.GetDefaultAddr(c.Query("addrFrom"))
		pkFrom := C.GetDefaultPk(addrFrom)
		addrTo := C.GetDefaultAddr(c.Query("addrTo"))
		amount, err := strconv.ParseInt(c.DefaultQuery("amount", "5000"), 10, 64)
		if err != nil {
			errMsg := "Fail to parse amount into a integer: " + err.Error()
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": errMsg})
			log.Panic(errMsg)
		}
		prevAddrfromEthBalance, prevAddrtoEthBalance := checkWalletBalance(addrFrom), checkWalletBalance(addrTo)
		prevAddrfromWwkfBalance, prevAddrtoWwkfBalance := getWwkfBalance(addrFrom), getWwkfBalance(addrTo)
		signedTxStr := transferWwkfWithAmount(pkFrom, addrTo, amount)
		retStatus := make(chan bool)
		go waitForTxCompletion(signedTxStr, retStatus, 60)
		txStatus := <-retStatus
		if !txStatus {
			errMsg := "Fail to wait for the TX getting approved"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			log.Panic(errMsg)
		}

		currAddrfromEthBalance, currAddrtoEthBalance := checkWalletBalance(addrFrom), checkWalletBalance(addrTo)
		currAddrfromWwkfBalance, currAddrtoWwkfBalance := getWwkfBalance(addrFrom), getWwkfBalance(addrTo)
		c.JSON(
			http.StatusOK,
			gin.H{
				"txAddr":                      signedTxStr,
				"addrfrom":                    addrFrom,
				"addrto":                      addrTo,
				"wwkfAmount":                  amount,
				"previousAddrfromWwkfBalance": prevAddrfromWwkfBalance,
				"previousAddrtoWwkfBalance":   prevAddrtoWwkfBalance,
				"currentAddrfromWwkfBalance":  currAddrfromWwkfBalance,
				"currentAddrtoWwkfBalance":    currAddrtoWwkfBalance,
				"previousAddrfromEthBalance":  prevAddrfromEthBalance,
				"previousAddrtoEthBalance":    prevAddrtoEthBalance,
				"currentAddrfromEthBalance":   currAddrfromEthBalance,
				"currentAddrtoEthBalance":     currAddrtoEthBalance,
			},
		)
	})

	api.Run()
}
