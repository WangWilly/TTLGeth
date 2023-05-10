import { Component } from "react";
import Web3 from "web3";
import { Contract } from "web3-eth-contract";
import { Log } from "web3-core";
import C from "./utils/constants";
import { SharedWeb3, SharedWwkf } from "./utils/sharedWeb3";

type Props = {};
type State = {
  contract: null | Contract;
  logs: Log[];
};

const web3 = SharedWeb3.getInstance();
const wwkfInstance = SharedWwkf.getInstance();

class EventLogs extends Component<Props, State> {
  state: State = {
    contract: null,
    logs: [],
  };

  constructor(props: Props) {
    super(props);
  }

  componentDidMount() {
    if (web3 == null) {
      throw new Error("Metamask provider is not available.");
    }

    console.log(`contract.options.address: ${wwkfInstance.options.address}`);
    console.log(`C.WWKF_CONTRACT.TOPICS.TRANSFER: ${C.WWKF_CONTRACT.TOPICS.TRANSFER}`);
    // Subscribe to logs events with optional parameters; Promised results
    const sub = web3.eth.subscribe(
      "logs",
      {
        fromBlock: C.WWKF_CONTRACT.BEGIN_BLOCK_ID,
        address: wwkfInstance.options.address,
        topics: [C.WWKF_CONTRACT.TOPICS.TRANSFER],
      },
      (error, result) => {
        if (error) {
          throw new Error(`subscription error in EventLogs: ${error}`);
        }
        console.log("Obtain new event");
        // update the state with the new log object
        this.setState((prevState) => ({
          logs: [...prevState.logs, result],
        }));
      }
    );
    // debugger;
    console.log(`sub: ${sub}`);
  }

  render() {
    return (
      <div>
        <p className="text-2xl">WWKF event logs of 'Transfer':</p>
        <ul className="text-left">
          {this.state.logs.reverse().map((log) => {
            const txLogIdx = `${log.transactionHash.substring(log.transactionHash.length - 8)}-${log.logIndex}`;
            const transferedAmount = Web3.utils.toNumber(log.data);
            return (
              <li className="rounded-md overflow-auto bg-zinc-400 text-white m-2 p-2" key={txLogIdx}>
                {/* display some properties of the log object */}
                <p>Tx-logIdx: {txLogIdx}</p>
                <p>Block number: {log.blockNumber}</p>
                <p>Transfered amount: {transferedAmount}</p>
                <p>From: 0x{log.topics[1].substring(26)}</p>
                <p>To: 0x{log.topics[2].substring(26)}</p>
              </li>
            );
          })}
        </ul>
      </div>
    );
  }
}

export default EventLogs;
