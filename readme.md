# Time to Learn Geth

Base on the template `.env.template`, fill up all variable and create a file `.env` for all crediential information. Have fun and join the virtual world. 🥷

📖 Instruction:
- ✅: Completed goal
- 👉: Short-term goal
- 📝: Long-term goal
- 📖: Some relative note
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
- 👉 https://www.honeybadger.io/blog/golang-go-package-management/
- 📝 https://github.com/liyue201/erc20-go/blob/main/erc20/token.go


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


## 🌶️ GO cheat sheet

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

### Why Ethereum mainet would not be confused with testnet?
Ethereum mainnet and testnet have different **network IDs** and **chain IDs** that distinguish them from each other. 📌: "For example, the Ethereum mainnet network ID is 1, while the other most commonly used testnets have network IDs of 3, 4, and 42 for Ropsten, Rinkeby, and Kovan, respectively." Network ID and chain ID are parameters that are used by nodes to identify which network they belong to and which transactions they should validate. Therefore, Ethereum mainnet would not be confused with testnet by nodes or clients that use these parameters correctly.

Source: Conversation with Bing, 4/12/2023(1) Networks | ethereum.org. https://ethereum.org/en/developers/docs/networks/ Accessed 4/12/2023.
(2) Comparison of the different TestNets - Ethereum Stack Exchange. https://ethereum.stackexchange.com/questions/27048/comparison-of-the-different-testnets Accessed 4/12/2023.
(3) Testnet Vs Mainnet in Blockchain - A Complete Guide | Shardeum. https://shardeum.org/blog/testnet-vs-mainnet/ Accessed 4/12/2023.

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

