package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/dollarkillerx/urllib"
)

func TestAlchemy(t *testing.T) {
	//address := "0xd1f84ec6652b315e77ebffa649f6742754574796"
	address2 := "0xa10A5e7E2b2CC6fcaF4e74B00a163B10ee060eBf"

	//kpl(address)
	kpl(address2)
}

func kpl(address string) {
	//input := map[string]interface{}{
	//	"jsonrpc": "2.0",
	//	"method":  "alchemy_getTokenBalances",
	//	"headers": map[string]interface{}{
	//		"Content-Type": "application/json",
	//	},
	//	"params": []string{
	//		address,
	//		"erc721",
	//		"erc1155",
	//	},
	//	"id": 42,
	//}

	input := map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "alchemy_getAssetTransfers",
		"params": []map[string]interface{}{
			{
				"fromBlock": "0x0",
				"toBlock":   "latest",
				"category": []string{
					"erc721",
				},
				"contractAddresses": []string{
					"0xa10A5e7E2b2CC6fcaF4e74B00a163B10ee060eBf",
				},
				"withMetadata":     false,
				"excludeZeroValue": true,
				"maxCount":         "0x3e8",
				//"fromAddress":      address,
			},
		},
	}

	i, bytes, err := urllib.
		Post("https://polygon-mainnet.g.alchemy.com/v2/g9_dtmuJeERcETxHWLr2fBqgdz9qcie4").
		SetJsonObject(input).Byte()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(i)
	fmt.Println(string(bytes))

	os.WriteFile("alchemy.json", bytes, 00666)
}

func TestP2(t *testing.T) {
	r := "0x0000000000000000000000000000000000000000000000000000000000000001"
	val := r[2:]
	i, err := strconv.ParseInt(val, 16, 64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(i)
}

func TestSig(t *testing.T) {
	sig := "0xd8f28606f13a6ce568ce6b6222251ee2bffd02b41fbe039ae8129ed1d02e84313554a5b57ac5efe89f4c45a9ea78192a4436fbf783fbcef14840939b24dd0dd41b"
	data := "Login to the website"
	address := "0xe916857755caee785ec694f096583f1be994892d"

	fmt.Println(verifySig(
		address,
		sig,
		[]byte(data),
	))

	fmt.Println(verifySig2(
		address,
		sig,
		[]byte(data),
	))

	fmt.Println(verifySig3(
		address,
		sig,
		[]byte(data),
	))
}

func verifySig(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	return from == recoveredAddr.Hex()
}

func verifySig2(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	return from == recoveredAddr.Hex()
}

func verifySig3(from, sigHex string, msg []byte) bool {
	fromAddr := common.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)
	if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27

	pubKey, err := crypto.SigToPub(signHash(msg), sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Println("addr: ", recoveredAddr)
	return fromAddr == recoveredAddr
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
