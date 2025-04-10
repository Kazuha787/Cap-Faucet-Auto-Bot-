package main

import (
	"bufio"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
)

var (
	rpcURL       = "https://carrot.megaeth.com/rpc"
	contractAddr = "0xe9b6e75c243b6100ffcb1c66e8f78f96feea727f"
	chainID      = big.NewInt(6342)
	tokenAmount  = new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18))
	mintSig      = "40c10f19"
	privateKeys  []string

	successCount = 0
	failCount    = 0
	claimCount   = 0
	currentClaim = 0
)

func main() {
	printBanner()
	loadPrivateKeys("privatkey.txt")

	if len(privateKeys) == 0 {
		color.Red("No private keys found. ❌")
		return
	}

	claimCount = promptClaimCount()

	// Loop over all wallets and execute the claims
	for i := 0; i < len(privateKeys); i++ {
		walletPrivateKey := privateKeys[i]
		color.Yellow("🌟 Using wallet %d: %s", i+1, getWalletAddress(walletPrivateKey))

		// Handle multiple transactions for each wallet
		for j := 0; j < claimCount; j++ {
			currentClaim = j + 1
			printProgress()
			animate("⏳ Claiming tokens", 3)

			success := sendClaim(walletPrivateKey)
			if success {
				color.Green("✅ Transaction %d successful for wallet %d!", currentClaim, i+1)
				successCount++
			} else {
				color.Red("❌ Transaction %d failed for wallet %d.", currentClaim, i+1)
				failCount++
			}

			if j < claimCount-1 {
				color.Cyan("⏱️ Waiting 5 seconds before next claim...\n")
				time.Sleep(5 * time.Second)
			}
		}
	}

	printSummary()
}

func sendClaim(privateKeyHex string) bool {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		color.Red("Failed to connect to RPC: %v", err)
		return false
	}
	defer client.Close()

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		color.Red("Invalid private key: %v", err)
		return false
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		color.Red("Failed to parse private key: %v", err)
		return false
	}

	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		color.Red("Failed to get nonce: %v", err)
		return false
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		color.Red("Failed to get gas price: %v", err)
		return false
	}

	methodSig := common.FromHex("0x" + mintSig)
	addressBytes := common.LeftPadBytes(fromAddr.Bytes(), 32)
	amountBytes := common.LeftPadBytes(tokenAmount.Bytes(), 32)
	data := append(methodSig, append(addressBytes, amountBytes...)...)

	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(contractAddr),
		big.NewInt(0),
		uint64(100000),
		gasPrice,
		data,
	)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		color.Red("Failed to sign transaction: %v", err)
		return false
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		color.Red("Failed to send transaction: %v", err)
		return false
	}

	color.Green("Transaction sent! Hash: %s", signedTx.Hash().Hex())
	return true
}

func loadPrivateKeys(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		color.Red("Error reading file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "0x") {
			privateKeys = append(privateKeys, line)
		}
	}
}

func promptClaimCount() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		color.Cyan("🛠️ How many transactions do you want to run per wallet? (Enter a number)")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		count, err := strconv.Atoi(input)
		if err == nil && count > 0 {
			return count
		}
		color.Red("❌ Please enter a valid number.")
	}
}

func printBanner() {
        bold := color.New(color.Bold, color.FgCyan).SprintFunc()
        fmt.Println(bold("╔════════════════════════════════════════════════════╗"))
        fmt.Println(bold("║                CAPS FAUCET MINT BOT                ║"))
        fmt.Println(bold("║         Automate your Caps Faucet (MEGA-ETH)       ║"))
        fmt.Println(bold("║    Developed by: https://t.me/Offical_Im_kazuha    ║"))
        fmt.Println(bold("║    GitHub: https://github.com/Kazuha787            ║"))
        fmt.Println(bold("╠════════════════════════════════════════════════════╣"))
        fmt.Println(bold("║  ██╗  ██╗ █████╗ ███████╗██╗   ██╗██╗  ██╗ █████╗  ║"))
        fmt.Println(bold("║  ██║ ██╔╝██╔══██╗╚══███╔╝██║   ██║██║  ██║██╔══██╗ ║"))
        fmt.Println(bold("║  █████╔╝ ███████║  ███╔╝ ██║   ██║███████║███████║ ║"))
        fmt.Println(bold("║  ██╔═██╗ ██╔══██║ ███╔╝  ██║   ██║██╔══██║██╔══██║ ║"))
        fmt.Println(bold("║  ██║  ██╗██║  ██║███████╗╚██████╔╝██║  ██║██║  ██║ ║"))
        fmt.Println(bold("║  ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝ ║"))
        fmt.Println(bold("╚════════════════════════════════════════════════════╝"))
        fmt.Println("✨ Ready to automate token claims! ✨")
        fmt.Println()
}

func printProgress() {
	color.Magenta("Progress: %d / %d | ✅ %d | ❌ %d", currentClaim, claimCount, successCount, failCount)
}

func printSummary() {
	color.Blue("\n╔═══════════════════ Summary ═══════════════════╗")
	color.Cyan("║ Total Transactions: %d", claimCount*len(privateKeys))
	color.Green("║ ✅ Successful: %d", successCount)
	color.Red("║ ❌ Failed:     %d", failCount)
	color.Blue("╚═══════════════════════════════════════════════╝")
}

func animate(msg string, dots int) {
	for i := 0; i < dots; i++ {
		fmt.Printf("\r%s%s", msg, strings.Repeat("🌑", i+1))
		time.Sleep(500 * time.Millisecond) // Adding a small delay for each animation step
	}
}

func getWalletAddress(privateKeyHex string) string {
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		color.Red("Invalid private key: %v", err)
		return "Invalid Wallet"
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		color.Red("Failed to parse private key: %v", err)
		return "Invalid Wallet"
	}

	fromAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	return fromAddr.Hex()
}
