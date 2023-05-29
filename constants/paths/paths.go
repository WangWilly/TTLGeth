package paths

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	walletAddr1 = os.Getenv("WALLET_HASH_1")
	walletAddr2 = os.Getenv("WALLET_HASH_2")
	privateKey1 = os.Getenv("NO0x_PRIVATE_KEY_1")
	privateKey2 = os.Getenv("NO0x_PRIVATE_KEY_2")
)

const (
	WALLET_ADDR = "walletAddr"
)

const (
	P_CHECK_ETH     = "/checkEth"
	P_CHECK_WWKF    = "/checkWwkf"
	P_TRANSFER_ETH  = "/transferEth"
	P_TRANSFER_WWKF = "/transferWwkf"
)

const (
	PARAM_PREFIX     = "/:"
	PARAM_WALLETADDR = PARAM_PREFIX + WALLET_ADDR
)

var (
	//lint:ignore U1000 Ignore unused function temporarily for debugging
	predefinedWalletAddrs = map[string]string{
		"Wallet1": walletAddr1,
		"Wallet2": walletAddr2,
	}
	predefinedPrivateKeys = map[string]string{
		walletAddr1: privateKey1,
		walletAddr2: privateKey2,
	}
)

func GetDefaultAddr(addr string) string {
	if predefineAddr, ok := predefinedWalletAddrs[addr]; ok {
		return predefineAddr
	}
	return addr
}

func GetDefaultPk(addr string) string {
	predefineAddr := GetDefaultAddr(addr)
	if predefinePk, ok := predefinedPrivateKeys[predefineAddr]; ok {
		return predefinePk
	}
	panic("Fail to find a address from a defined list")
}
