import { useState } from "react";
import Web3 from "web3";
import wwkfJson from "./contracts/willywangkaaFirstContract.json";
import { AbiItem } from "web3-utils";

const wwkfAbi = wwkfJson as AbiItem[];
const web3 = new Web3(Web3.givenProvider); // use Metamask provider
const WWKF_CONTRACT_ADDR = "0x03378DAa43739f2361FE67175aD6bF2666309748"; // TODO: move to constant
const wwkfInstance = new web3.eth.Contract(wwkfAbi, WWKF_CONTRACT_ADDR);

declare global {
  interface Window {
    ethereum: any;
  }
}

function App() {
  const [accountAddr, setAccountAddr] = useState("");
  const [wwkfBalance, setWwkfBalance] = useState(0);

  // Click handler of button for connectting wallet
  const connectWallet = async () => {
    // Request access to accounts
    const accounts = await window.ethereum.request({
      method: "eth_requestAccounts",
    });
    // Set the first accounts as user accountAddr
    setAccountAddr(accounts[0]);
  };

  // Click handler of button for getting wwkf balance
  const getWwkfBalance = async () => {
    // Call the balanceOf function of wwkfInstance
    const result = await wwkfInstance.methods.balanceOf(accountAddr).call();
    // set the result as wwkfBalance state
    setWwkfBalance(result);
  };

  return (
    <>
      <div>
        <h1>WWKF naive demo</h1>
        <p>Edit <code>src/App.tsx</code> and save to test HMR</p>
        <p>Account addr: {accountAddr}</p>
        <p>Wwkf balance: {wwkfBalance}</p>
        <div className="card">
          <button onClick={connectWallet}>Connect Wallet</button>
          <button onClick={getWwkfBalance}>Get Wwkf Balance</button>
        </div>
      </div>
    </>
  );
}

export default App;