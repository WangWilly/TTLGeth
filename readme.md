# Time to Learn Geth

Base on the template `.env.template`, fill up all variable and create a file `.env` for all crediential information. Have fun and join the virtual world. 🥷

📖 Instruction:

- ✅: Completed goal
- 👉: Short-term goal
- 📝: Long-term goal
- 📖: Some relative note (📌: highlight)
- 🤔: Not sure the credit
- 🔜: Haven't completed but browsed
- ❌: Abandoned or failed

📖 Easy run:

```
go run *.go
```

📖 Reviews:

- ✅ [Connect ethereum node using Infura](https://blog.logrocket.com/ethereum-development-using-go-ethereum/#connecting-ethereum-node-infura-go)
  - 📖 [Contract migration1](https://trufflesuite.com/docs/truffle/how-to/contracts/run-migrations/)
  - 📖 [Contract migration2](https://betterprogramming.pub/how-to-write-complex-truffle-migrations-86d4b85d7783)
- ✅ [Get ERC-20 balance](https://levelup.gitconnected.com/how-to-get-balance-of-an-ethereum-smart-contract-91ce4e7b4c4e)
- ✅ [Transfer ERC-20 token](https://goethereumbook.org/en/transfer-tokens/)
- ✅ [Develop Go with binding of contract](https://goethereumbook.org/en/smart-contract-load/)
  - https://medium.com/nerd-for-tech/smart-contract-with-golang-d208c92848a9
  - https://github.com/02amanag/SmartContractWithGolang/blob/main/cmd/main.go
- ✅ [Building a REST API with Golang using Gin and Gorm](https://blog.logrocket.com/rest-api-golang-gin-gorm/)
  - https://davidchan0519.github.io/2019/05/06/gorm-automigrate/
  - https://github.com/gin-gonic/examples/blob/master/favicon/main.go
  - https://gin-gonic.com/docs/examples/querystring-param/
  - https://go.dev/src/math/const.go
  - https://stackoverflow.com/questions/21532113/golang-converting-string-to-int64
- 👉 [Manage private key](https://goethereumbook.org/en/keystore/)
- 🔜 https://www.honeybadger.io/blog/golang-go-package-management/
- 📝 **https://goethereumbook.org/event-subscribe/**
  - **https://blog.logrocket.com/rest-api-golang-gin-gorm/**
  - https://gin-gonic.com/
  - **https://github.com/labstack/echo**
  - https://willh.gitbook.io/build-web-application-with-golang-zhtw/03.0/03.2
- GitHub Pages
  - https://github.com/ngseke/ngseke.me/blob/master/.github/workflows/cd.yml
  - https://vitejs.dev/guide/static-deploy.html
  - https://stackoverflow.com/questions/58139175/running-actions-in-another-directory
  - https://stackoverflow.com/questions/58139175/running-actions-in-another-directory
  - https://nkaushik.com/nodejs/dependencies-lock-file-not-found/#:~:text=Fix%201%3A%20Remove%20cache%3A
  - [🎥 WWKF VIEWER DEMO](https://youtu.be/AAHxSEEqN7A)
- [🌶️ Dive Into Go](/docs/diveIntoGo.md)

## Achievement

- 🙋‍♂️ Protagonists
  - Wallet1: `0x29F4E2b19BEBc88F6BD3cf794D1F93af0D76B4E4`
  - Wallet2: `0xaAc4Dd69227C8e6412CB0B0479eF57f6245e7e75`
- 📖 History
  - [✅](https://sepolia.etherscan.io/tx/0xfa3b69c6477411911528322d321d607a649e99ebafa75a5a01a196e505bc7aa0) Share the acquired eth conins to another generated wallet.
  - [✅](https://sepolia.etherscan.io/tx/0xc8ba741daff2597b3b805dd093273d908f69ad601a639fa64a954ce3b00b176e) Create WWKF token. ([📖 token history](https://sepolia.etherscan.io/address/0x03378daa43739f2361fe67175ad6bf2666309748), [📖 token information](https://sepolia.etherscan.io/token/0x03378daa43739f2361fe67175ad6bf2666309748))
  - [Link](./docs/wwkfTransectionDetail.md) to the raw outputs of the following transection.
  - [❌](https://sepolia.etherscan.io/tx/0xb56d48fdab44c133959c5ed89ac059b2ef9fc8e75a3fda1145b33b34c33d1896) Fail to transfer WWKF token. (error: transfer excessive amounts of WWKF)
  - [❌](https://sepolia.etherscan.io/tx/0x64e1fa9d239b03f0d274e5fb2ac3f89379a43e3918d62975b3b63cc7caf1b119) Fail to transfer WWKF token. (error: transfer excessive amounts of WWKF)
  - [❌](https://sepolia.etherscan.io/tx/0x2105c498bf22a718ec2f0e62205e49e367e2fdad51141ceb66cbdd80ad894970) Fail to transfer WWKF token. (error: transfer excessive amounts of WWKF)
  - [✅](https://sepolia.etherscan.io/tx/0xf26f5deefe13c60eb419e97485ba2d783190bae804d04bc6ff952f30f2d5849f) Complete transfering WWKF token.
  - [✅](https://sepolia.etherscan.io/tx/0x0f313aa1a67826a6d8519043d12e83a60861bbffbee264a4ea93e063b1bff8aa) Complete transfering WWKF token.
  - [✅](https://sepolia.etherscan.io/tx/0x5952cd7f5a5af01882214b170f7289e272fa8da19c75e6bc1ff1eb978066eabe) Complete transfering WWKF token.
  - [✅](https://sepolia.etherscan.io/tx/0x74da57c57959d4337d10ffdd034d88376a9efba3146437037bb8a911c9d5404a) Complete transfering WWKF token.
  - [✅](https://sepolia.etherscan.io/tx/0x399f6db082e6951c4b6ae8d4f4522ae0680477f098dd056a1a45ab4eed27a9af) Complete transfering WWKF token.

## 🌶️ GO cheat sheet

📌 **Goroutine**

- https://golangbot.com/channels/
- https://blog.wu-boy.com/2020/05/understant-golang-context-in-10-minutes/

📌 General usage

- `len()`: `len(v Type) int`

```Go
// [string]
s := "Hello"
fmt.Println(len(s)) // 👉 5
// [array]
a := [3]int{1, 2, 3}
fmt.Println(len(a)) // 👉 3
// [map]
m := map[string]int{"one": 1, "two": 2}
fmt.Println(len(m)) // 👉 2
```

📌 Interaction

- `fmt.Print`
- `fmt.Println`
- `log.Fatal`

📌 Handle String

- `strings.TrimLeft`: remove left string.
  - `strings.TrimLeft(resp[2:], "0")`

📌 Handle Web3 Address

- `hexutil.Encode`: encode byte into a string with leading '0x'
- `hexutil.DecodeBig`: decode string with leading '0x' into a `big.Int`
- `common.HexToAddress`: transfrom hex string into byte address

📖 Reference

- https://blog.wu-boy.com/2019/04/how-to-load-env-file-in-go/

## 🤖 Robot summary

Quick accession of network result.

### 🤔 Is there a web3 package in Go?

I found some web search results that might be helpful.

According to the first result¹, **there is a web3 package in Go that is hosted on GitHub and provides various functions and types for interacting with web3 blockchains**. You can use web3 to create and send transactions, query the network and contract state, generate code from ABI files, and more. You can import web3 in your Go code and use its methods.

According to the second result², there is another web3 package in Go that is also hosted on GitHub and provides Ethereum Golang API, JSON-RPC client, and smart contract transaction API. You can use web3 to interact with an Ethereum node using JSON RPC, deploy and call smart contracts, and more. You can install web3 using go get or download the binary from GitHub.

According to the third result³, there is yet another web3 package in Go that is hosted on GitHub and provides a wrapper for Ethereum's web3.js for use with gopherjs. You can use web3 to access the same functionality as web3.js in your Go code, such as sending transactions, querying the state of the network, and accessing smart contracts. You can install web3 using go get or download the source code from GitHub.

I hope this helps you with your query.😊

Source: Conversation with Bing, 4/20/2023(1) web3 package - github.com/gochain/web3 - Go Packages. https://pkg.go.dev/github.com/gochain/web3 Accessed 4/20/2023.
(2) chenzhijie/go-web3: Ethereum Golang API, JSON-RPC client, Smart contract transaction API. - Github. https://github.com/chenzhijie/go-web3 Accessed 4/20/2023.
(3) web3 package - github.com/jaynagpaul/go-web3 - Go Packages. https://pkg.go.dev/github.com/jaynagpaul/go-web3 Accessed 4/20/2023.

### 🤔 How to obatin Smart contract events?

I found some web search results that might be helpful.

Smart contract events are signals that smart contracts emit when something relevant happens within them based on their code. **For example, a smart contract may emit an event when a token is transferred, a state variable is changed, or an error occurs.** Smart contract events can be used to notify developers and software systems that something meaningful has happened, and **to enable communication between smart contracts and dapps¹²**.

**To obtain smart contract events, you need to listen to them using a web3 library that can interact with the blockchain.** **One such library is ethers.js, which provides methods for connecting to a provider, getting a contract instance, and subscribing to contract events.** You can also filter events by topics or arguments, and access the event data and logs³.

I hope this helps you. If you need more details, you can check out the web search results below.

¹: Logging data from smart contracts with events | ethereum.org
²: Blockchain Explained: What are smart contract events and logs?
³: How to Listen to Smart Contract Events Using Ethers.js

Source: Conversation with Bing, 4/20/2023(1) How to Listen to Smart Contract Events Using Ethers.js. https://bing.com/search?q=smart+contract+events Accessed 4/20/2023.
(2) What are Polling Cycles in Shardeum? : r/shardeum - Reddit. https://www.reddit.com/r/shardeum/comments/12rfapk/what_are_polling_cycles_in_shardeum/ Accessed 4/20/2023.
(3) Logging data from smart contracts with events | ethereum.org. https://ethereum.org/en/developers/tutorials/logging-events-smart-contracts/ Accessed 4/20/2023.
(4) How to Listen to Smart Contract Events Using Ethers.js. https://moralis.io/how-to-listen-to-smart-contract-events-using-ethers-js/ Accessed 4/20/2023.
(5) Web3 Syncing – How to Sync Smart Contract Web3 Events. https://moralis.io/web3-syncing-how-to-sync-smart-contract-web3-events/ Accessed 4/20/2023.
(6) Blockchain Explained: What are smart contract events and logs?. https://medium.com/web3labs/blockchain-explained-what-are-smart-contract-events-and-logs-bd44f1e99b61 Accessed 4/20/2023.

### 📖 How to use JSON RPC in Ethereum?

I found some web search results that might be helpful.

According to the first result¹, **JSON RPC is a protocol that allows you to interact with an Ethereum node using HTTP, WebSocket, or IPC.** You can use JSON RPC to send requests and receive responses from the node, such as getting the balance of an account, sending a transaction, or subscribing to events. **You can use various libraries or tools to make JSON RPC calls, such as web3.js, curl, or Postman.**

According to the second result², **JSON RPC is a way to communicate with an Ethereum node and perform actions on the network, such as staking on the Goerli testnet.** You can use JSON RPC to create and send transactions, query the state of the network, and access smart contracts. You can use abigen to generate a Go file that binds your smart contract and use Geth to connect with the node.

According to the third result³, 📌 **JSON RPC is a method to access data from an Ethereum node and decode it into a human-readable format. You can use JSON RPC to get information about blocks, transactions, receipts, logs, and events.** You can also use JSON RPC to execute smart contract functions and monitor the network status.

I hope this helps you with your query.😊

Source: Conversation with Bing, 4/20/2023(1) How to Stake on the Goerli ETH Testnet (Ethereum). https://docs.blockdaemon.com/docs/how-to-stake-on-the-goerli-eth-testnet-ethereum Accessed 4/20/2023.
(2) OKTC docs | User guide | Developer docs | OKX. https://www.okx.com/en-in/oktc/docs/dev/quick-start/join-oktc-testnet Accessed 4/20/2023.
(3) Decoding Decentralization. An Ethereum Blockchain Data Primer… | by Alexander Sack | Coinmonks | Apr, 2023 - Medium. https://medium.com/coinmonks/decoding-decentralization-f3af30068e75 Accessed 4/20/2023.

### 🤔 How etherscan obtain the history of a contract?

I found some web search results that might be helpful.

According to the first result¹, etherscan does not store all the events and transaction history, but rather indexes them using a custom backend that connects to an Ethereum node and fetches transactions using **JSON RPC**. **Etherscan also uses a separate service to track internal transactions that are not part of the blockchain.**

According to the second result², you can use go-ethereum to explore all transactions for a given account by using the eth_getBlockByNumber JSON RPC method and filtering by the account address. You can also use the **web3.js library to make it easier to interact with the Ethereum node**.

According to the third result³, you can find a question and answer site for Ethereum enthusiasts where you can ask or answer questions related to Ethereum and smart contracts.

I hope this helps you with your query.😊

Source: Conversation with Bing, 4/20/2023(1) Does Etherscan store all the events and transaction history?. https://ethereum.stackexchange.com/questions/52104/does-etherscan-store-all-the-events-and-transaction-history Accessed 4/20/2023.
(2) How To Tell If An NFT Is Real: A Quick Guide For Savvy Collectors. https://www.forbes.com/sites/williamanderson/2023/04/13/how-to-tell-if-an-nft-is-real-a-concise-guide-for-savvy-collectors/ Accessed 4/20/2023.
(3) How to explore all transactions for a given account?. https://ethereum.stackexchange.com/questions/8547/how-to-explore-all-transactions-for-a-given-account Accessed 4/20/2023.

### 📖 Why Ethereum mainet would not be confused with testnet?

Ethereum mainnet and testnet have different **network IDs** and **chain IDs** that distinguish them from each other. 📌 **For example, the Ethereum mainnet network ID is 1, while the other most commonly used testnets have network IDs of 3, 4, and 42 for Ropsten, Rinkeby, and Kovan, respectively.** Network ID and chain ID are parameters that are used by nodes to identify which network they belong to and which transactions they should validate. Therefore, Ethereum mainnet would not be confused with testnet by nodes or clients that use these parameters correctly.

Source: Conversation with Bing, 4/12/2023(1) Networks | ethereum.org. https://ethereum.org/en/developers/docs/networks/ Accessed 4/12/2023.
(2) Comparison of the different TestNets - Ethereum Stack Exchange. https://ethereum.stackexchange.com/questions/27048/comparison-of-the-different-testnets Accessed 4/12/2023.
(3) Testnet Vs Mainnet in Blockchain - A Complete Guide | Shardeum. https://shardeum.org/blog/testnet-vs-mainnet/ Accessed 4/12/2023.

### 📖 What's the meaning of a starsign in Golang?

**A starsign (or asterisk) is a symbol used in Go to indicate pointers or dereferencing.** A **pointer** is a variable that stores the memory address of another variable. **Dereferencing** is the process of accessing the value stored at that memory address². 📌 **For example, if x is a pointer to an integer variable y, then \*x is the value of y.**

Source: Conversation with Bing, 4/20/2023(1) What does the asterisk do in "Go"? - Stack Overflow. https://stackoverflow.com/questions/3552626/what-does-the-asterisk-do-in-go Accessed 4/20/2023.
(2) strings package - strings - Go Packages. https://pkg.go.dev/strings Accessed 4/20/2023.
(3) Go StartsWith(str string) - Stack Overflow. https://stackoverflow.com/questions/12667327/go-startswithstr-string Accessed 4/20/2023.

### 📖 What's the meaning of Panic in Go?

According to web search results, **panic in Go is a way of handling unexpected or exceptional situations that cause the program to crash**. It is similar to an exception in other languages, but it also **terminates the normal execution flow and prints the error message and stack trace**. You can use the panic() function to deliberately trigger a panic, or it can happen automatically when some specific situation arises, such as an out-of-bounds array access. 📌 **You can also use defer and recover to handle panics gracefully and prevent the program from crashing.** ¹²³⁴

Source: Conversation with Bing, 4/20/2023(1) Panic in Golang - GeeksforGeeks. https://www.geeksforgeeks.org/panic-in-golang/ Accessed 4/20/2023.
(2) Handling Panics in Go | DigitalOcean. https://www.digitalocean.com/community/tutorials/handling-panics-in-go Accessed 4/20/2023.
(3) Golang Panic | How does Panic Work in Go Language with Examples - EduCBA. https://www.educba.com/golang-panic/ Accessed 4/20/2023.
(4) Golang panic handing [capture, defer, recover, log]. https://www.golinuxcloud.com/golang-panic/ Accessed 4/20/2023.

### ❌ How to save byte object in Go?

One way to save a byte object in Go is to use the ioutil.WriteFile function from the standard library¹. This function takes a file name, a byte slice, and a permission mode as arguments and writes the byte slice to the file. You can find an example of how to use this function in this article¹.

Another way to save a byte object in Go is to use the os.WriteFile function from the os package². This function is similar to ioutil.WriteFile but it is not deprecated as of Go 1.16². You can find an example of how to use this function in this answer².

Source:
[👉] How to covert a []byte object to an image and store it as a jpeg image on disk - Stack Overflow. https://stackoverflow.com/questions/52946929/how-to-covert-a-byte-object-to-an-image-and-store-it-as-a-jpeg-image-on-disk
[👉] Convert back byte array into file using golang - Stack Overflow. https://stackoverflow.com/questions/32687985/convert-back-byte-array-into-file-using-golang

Source: Conversation with Bing, 4/11/2023(1) go - How to covert a []byte object to an image and store it as a jpeg image on disk - Stack Overflow. https://stackoverflow.com/questions/52946929/how-to-covert-a-byte-object-to-an-image-and-store-it-as-a-jpeg-image-on-disk Accessed 4/11/2023.
(2) go - Convert back byte array into file using golang - Stack Overflow. https://stackoverflow.com/questions/32687985/convert-back-byte-array-into-file-using-golang Accessed 4/11/2023.
(3) go - Save []byte with AQL? - Stack Overflow. https://stackoverflow.com/questions/66769768/save-byte-with-aql Accessed 4/11/2023.
(4) Save byte[,,] in ScriptableObject - Unity Answers. https://answers.unity.com/questions/817868/save-byte-in-scriptableobject.html Accessed 4/11/2023.

### ❌ How to track a smart contract using Golang?

However, I found some web search results that might be helpful.

One possible way to track smart contracts using Golang is to create Go bindings for the specific type of contract. Go bindings are generated from the Application Binary Interface (ABI) of the smart contract, which describes how to call its functions and what kind of data you get back. **You can use the abigen tool, which is part of go-ethereum, to generate Go bindings from an ABI JSON file or a Solidity source file².**

Another possible way to track smart contracts using Golang is to use a framework like Truffle, which can compile and deploy smart contracts and generate ABI files for them. You can then use abigen to create Go bindings from the ABI files and interact with the smart contracts through Golang API¹.

A third possible way to track smart contracts using Golang is to use online tools like remix.ethereum.org or Etherscan ABI API to compile and get the ABI of a smart contract. You can then use abigen to create Go bindings from the ABI and interact with the smart contract through Golang API²³.

I hope this helps you. If you need more details, you can check out the web search results below.

¹: Smart Contract with Golang. Interacting with smart contracts… | by Aman Agarwal | Nerd For Tech | Medium
²: Creating Go Bindings for Ethereum Smart Contracts · Chris Hager
³: GitHub - zupzup/smart-contracts-with-go: A simple example of how to deploy and interact with ETH smart contracts using Go

Source: Conversation with Bing, 4/20/2023(1) Creating Go Bindings for Ethereum Smart Contracts. https://www.metachris.com/2021/05/creating-go-bindings-for-ethereum-smart-contracts/ Accessed 4/20/2023.
(2) Smart Contract with Golang. Interacting with smart contracts… | by Aman Agarwal | Nerd For Tech | Medium. https://medium.com/nerd-for-tech/smart-contract-with-golang-d208c92848a9 Accessed 4/20/2023.
(3) GitHub - zupzup/smart-contracts-with-go: A simple example of how to deploy and interact with ETH smart contracts using Go. https://github.com/zupzup/smart-contracts-with-go Accessed 4/20/2023.

### Can i use web3js in frontend, and if so can you give me an example?

Yes, you can use web3js in frontend to connect your web application to Ethereum blockchain and smart contracts. You need to install web3js as a dependency in your project and create an instance of web3 with a provider that connects to an Ethereum node. You can use different providers such as Infura, Alchemy, or your own local node. You also need the contract ABI and address to interact with your smart contract functions¹. Here is an example of a React component that allows users to connect their Metamask wallet and call a contract function:

```javascript
import React, { useState } from "react";
import Web3 from "web3";
import contractABI from "./contractABI.json";

const web3 = new Web3(Web3.givenProvider); // use Metamask provider
const contractAddress = "0x123456789"; // your contract address
const contract = new web3.eth.Contract(contractABI, contractAddress); // create contract instance

function App() {
  const [account, setAccount] = useState(""); // user account address
  const [value, setValue] = useState(0); // value from contract

  // connect wallet button click handler
  const connectWallet = async () => {
    // request access to accounts
    const accounts = await window.ethereum.request({
      method: "eth_requestAccounts",
    });
    // set the first account as user account
    setAccount(accounts[0]);
  };

  // get value button click handler
  const getValue = async () => {
    // call the getValue function from contract
    const result = await contract.methods.getValue().call();
    // set the result as value state
    setValue(result);
  };

  return (
    <div className="App">
      <h1>Web3 Example</h1>
      <p>Account: {account}</p>
      <p>Value: {value}</p>
      <button onClick={connectWallet}>Connect Wallet</button>
      <button onClick={getValue}>Get Value</button>
    </div>
  );
}

export default App;
```

Source: Conversation with Bing, 4/25/2023(1) Web3.js — Javascript Ethereum API. https://web3js.org/ Accessed 4/25/2023.
(2) Building a web3 frontend with React - DEV Community. https://dev.to/rounakbanik/building-a-web3-frontend-with-react-340c Accessed 4/25/2023.
(3) Web3 Development Explained for Frontend Developers | Kong Inc.. https://konghq.com/blog/engineering/web3-basics-for-frontend-developers Accessed 4/25/2023.
(4) Tutorial: Building a web3 frontend with React - Medium. https://medium.com/scrappy-squirrels/tutorial-building-a-web3-frontend-with-react-e0a87ea3bad Accessed 4/25/2023.

### How the current environment inject to the 'Web3.givenProvider'?

The current environment injects to the Web3.givenProvider property by exposing a global API object that is ERC1193 compliant, such as **`window.ethereum` for Metamask or Mist¹²**. Web3.js detects this object and assigns it to Web3.givenProvider automatically. If the environment does not provide such an object, Web3.givenProvider will be null². The environment can inject the provider by using browser extensions, web3 browsers, or other methods³⁴.

Source: Conversation with Bing, 4/25/2023(1) ethereum - What is the difference between currentprovider and givenprovider in web3.js - Stack Overflow. https://stackoverflow.com/questions/55822581/what-is-the-difference-between-currentprovider-and-givenprovider-in-web3-js Accessed 4/25/2023.
(2) web3 | API | web3.js. https://docs.web3js.org/api/web3/ Accessed 4/25/2023.
(3) Difference between Javascript VM vs injected w3 vs web3 provider. https://ethereum.stackexchange.com/questions/40227/difference-between-javascript-vm-vs-injected-w3-vs-web3-provider Accessed 4/25/2023.
(4) Web3 Providers | DApp World. https://dapp-world.com/smartbook/web3-providers--CWMi Accessed 4/25/2023.

### Can i use web3 with Metamask as provider to subscribe event from a blockchain branch 'sepolia'?

Yes, you can use web3js with Metamask as provider to subscribe to events from a blockchain branch 'sepolia'. You need to install web3js as a dependency in your project and create an instance of web3 with Web3.givenProvider, which will use the provider injected by Metamask¹². You also need to add the sepolia network to your Metamask by using a custom RPC URL that connects to a sepolia node. You can use different services such as Alchemy, Infura, or Chainlist to get a sepolia RPC URL³ . You also need the contract ABI and address to interact with your smart contract events. You can use the web3.eth.Contract.events or web3.eth.subscribe methods to subscribe to specific events or logs from your contract . For example:

```javascript
import Web3 from "web3";
import contractABI from "./contractABI.json";

const web3 = new Web3(Web3.givenProvider); // use Metamask provider
const contractAddress = "0x123456789"; // your contract address
const contract = new web3.eth.Contract(contractABI, contractAddress); // create contract instance

// subscribe to an event called MyEvent
contract.events
  .MyEvent()
  .on("data", function (event) {
    console.log(event); // log the event object
  })
  .on("error", function (error) {
    console.log(error); // log the error
  });
```

Source: Conversation with Bing, 4/25/2023(1) How to Add Sepolia to Metamask. https://docs.alchemy.com/docs/how-to-add-sepolia-to-metamask Accessed 4/25/2023.
(2) How to Add the Sepolia Test Network to Your MetaMask Wallet. https://www.alchemy.com/overviews/how-to-add-sepolia-to-metamask Accessed 4/25/2023.
(3) How to Add Sepolia Testnet to MetaMask | Simple Guide (2023). https://www.datawallet.com/crypto/add-sepolia-to-metamask Accessed 4/25/2023.

## 🤖 Further robot summary

### How do I use web3.js?

I found some web search results that might be helpful.

According to the first result¹, web3.js is a JavaScript library that allows you to interact with an Ethereum node using JSON RPC. You can use web3.js to send transactions, query the state of the network, and access smart contracts. You can install web3.js using npm or yarn and use it in your JavaScript code.

According to the second result², web3 is a command line tool and a Golang package that allows you to interact with web3 enabled blockchains, such as GoChain and Ethereum. You can use web3 to get block and transaction details, transfer tokens, deploy and call smart contracts, and more. You can install web3 using go get or download the binary from GitHub.

According to the third result³, web3 is a Golang package that provides various functions and types for interacting with web3 blockchains. You can use web3 to create and send transactions, query the network and contract state, generate code from ABI files, and more. You can import web3 in your Golang code and use its methods.

I hope this helps you with your query.😊

Source: Conversation with Bing, 4/20/2023(1) Go Tutorial - W3School. https://www.w3schools.com/go/index.php Accessed 4/20/2023.
(2) GitHub - gochain/web3: A CLI tool to interact with web3 blockchains - Ethereum, GoChain, etc. https://github.com/gochain/web3 Accessed 4/20/2023.
(3) web3 package - github.com/gochain/web3 - Go Packages. https://pkg.go.dev/github.com/gochain/web3 Accessed 4/20/2023.
