# Cap Faucet Auto Bot
![IMG_20250410_094909_184](https://github.com/user-attachments/assets/1f62dc38-606c-47c0-9553-6632811530d3)


Welcome to the **Cap Faucet Auto Bot** repository! 🚀 This bot is designed to automate the process of claiming tokens from the **Cap Faucet** using Ethereum-based wallets. It provides an easy-to-use terminal interface powered by **Bubble Tea**, and it supports automated claims using multiple wallet private keys. 

## Features

- **Ethereum Token Claiming**: Automates the token claiming process directly from the Cap Faucet.
- **Multi-wallet Support**: Easily claim tokens using multiple Ethereum wallet private keys.
- **Dynamic Terminal UI**: Beautiful and interactive terminal UI with progress tracking, built using the **Bubble Tea** library. 🌈
- **Real-time Progress Updates**: See your claim progress in real-time with clear success and failure counts.
- **Simple Setup**: Quick to set up and run, just provide your private keys and the bot does the rest. ⚡️

## Requirements

Before running the bot, ensure you have the following installed:

- **Go 1.18+**: Make sure Go is installed and configured properly.
- **Ethereum Client (Infura, Alchemy, etc.)**: You need an Ethereum RPC URL to interact with the Ethereum network.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Kazuha787/Cap-Faucet-Auto-Bot-.git
   ```
2. Navigate to the project directory:
```sh
cd Cap-Faucet-Auto-Bot-
```
3. Install required dependencies:
```
go mod tidy
```

4. Add your Ethereum private keys to the privatekey.txt file. Each private key should be on a separate line and start with 0x.
```
nano privatkey.txt
```

6. Run the bot:
```
go run main.go
```
The bot will start, and you’ll be prompted to enter the number of claims you want to make.

Configuration

Ethereum RPC URL

The bot connects to the Ethereum network via an RPC URL. You can use a public Ethereum RPC provider like Infura or Alchemy.

Replace the rpcURL variable in main.go with your own Ethereum RPC URL:

var rpcURL = `YOUR_ETHEREUM_RPC_URL`

Private Keys

Store your Ethereum private keys in the privatekey.txt file. Each private key should be placed on a new line:
```
0xYOUR_PRIVATE_KEY_1
0xYOUR_PRIVATE_KEY_2
```

Make sure to never share your private keys with anyone, as they can be used to access your wallets.

## Usage

When you run the bot, you will be prompted to enter the number of claims you want to make. The bot will then:

Select a wallet key from your list.

Claim tokens from the Cap Faucet.

Show real-time progress and status of each claim.

Output the transaction hash for successful claims.

Example Output
```
╔════════════════════════════════════════════════════╗
║              Cap Faucet Auto Bot                   ║
║        Automate your Cap Faucet Token Claims!      ║
║ Developed by: https://github.com/Kazuha787         ║
╠════════════════════════════════════════════════════╣
Progress: 1 / 5 | ✅ 1 | ❌ 0
Using wallet 1: 0xPRIVATE_KEY_1
Claiming tokens...
✅ Claim successful!
Waiting 5 seconds before next claim...

Progress: 2 / 5 | ✅ 2 | ❌ 0
Using wallet 2: 0xPRIVATE_KEY_2
Claiming tokens...
✅ Claim successful!
Waiting 5 seconds before next claim...
```
## Contributing

We welcome contributions! Feel free to open issues, fork the repo, and submit pull requests. If you have any suggestions or improvements, don't hesitate to share them.

Steps to Contribute

1. Fork the repository.

2. Create a new branch: git checkout -b feature-name.

3. Make your changes.

4. Commit your changes: git commit -am 'Add new feature'.

5. Push to your branch: git push origin feature-name

6. Create a new pull request.

# License
This project is licensed under the MIT License - see the LICENSE file for details.
Made with ❤️ by Kazuha787.
```
This README provides a complete guide for your repository, including installation, usage, and contributing guidelines. Feel free to modify or enhance it based on your specific needs!
```
