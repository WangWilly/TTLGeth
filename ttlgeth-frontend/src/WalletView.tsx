import { useState } from "react";
import Web3 from "web3";
import wwkfJson from "./contracts/willywangkaaFirstContract.json";
import { AbiItem } from "web3-utils";
import C from "./utils/constants";

// 🤔: siglton method
const wwkfAbi = wwkfJson as AbiItem[];
const web3 = new Web3(Web3.givenProvider); // use Metamask provider
const wwkfInstance = new web3.eth.Contract(wwkfAbi, C.WWKF_CONTRACT.ADDR);

function WalletView() {
  const [accountAddr, setAccountAddr] = useState("");
  const [wwkfBalance, setWwkfBalance] = useState(0);

  // Click handler of button for connectting wallet
  const connectWallet = async () => {
    // Request access to accounts
    const accounts = await window.ethereum.request<string[]>({
      method: "eth_requestAccounts",
    });
    if (accounts == null || accounts == undefined || accounts.length == 0) {
      throw new Error("No account is obtained.");
    }
    // Set the first accounts as user accountAddr
    setAccountAddr((accounts as string[])[0]);
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
        <p className="text-2xl">Check your WWKF status:</p>
        <div className="rounded-md overflow-auto bg-zinc-400 text-white p-2 m-2">
          <p>Account addr: {accountAddr}</p>
          <p>Wwkf balance: {wwkfBalance}</p>
        </div>
        <div className="shadow-sm bg-slate-500 p-8 rounded-md flex flex-row m-2">
          <button className="basis-1/2 m-2" onClick={connectWallet}>
            Connect Wallet
          </button>
          <button className="basis-1/2 m-2" onClick={getWwkfBalance}>
            Get Wwkf Balance
          </button>
        </div>
      </div>
    </>
  );
}

export default WalletView;
