import { Component } from "react";
import Web3 from "web3";
import { Contract } from "web3-eth-contract";
import { Log } from "web3-core";
import { AbiItem } from "web3-utils";
import C from "./utils/constants";
import wwkfJson from "./contracts/willywangkaaFirstContract.json";

type Props = {};
type State = {
  contract: null | Contract;
  logs: Log[];
};

// 🤔: siglton method
const wwkfAbi = wwkfJson as AbiItem[];
const web3 = new Web3(Web3.givenProvider); // use Metamask provider
const wwkfInstance = new web3.eth.Contract(wwkfAbi, C.WWKF_CONTRACT.ADDR);

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
        // topics: [C.WWKF_CONTRACT.TOPICS.TRANSFER],
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
        <p className="text-2xl">WWKF event logs:</p>
        <ul className="text-left">
          {this.state.logs.map((log) => (
            <li className="rounded-md overflow-auto bg-zinc-400 text-white m-2 p-2" key={log.logIndex}>
              {/* display some properties of the log object */}
              {/* <p>Address: {log.address}</p> */}
              <p>LogIndex: {log.logIndex}</p>
              <p>Tx: {log.transactionHash}</p>
              <p>Data: {log.data}</p>
              <p>Block Number: {log.blockNumber}</p>
              {/* <p>Topics: {log.topics.join(",")}</p> */}
            </li>
          ))}
        </ul>
      </div>
    );
  }
}

export default EventLogs;
