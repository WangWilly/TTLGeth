import Web3 from "web3";
import Contract from "web3-eth-contract";
import { AbiItem } from "web3-utils";
import C from "./constants";
import wwkfJson from "../contracts/willywangkaaFirstContract.json";

/**
 * The class defines the `getInstance` method that lets clients access
 * the unique singleton instance.
 *
 * Reference: https://refactoring.guru/design-patterns/singleton/typescript/example
 */
export class SharedWeb3 extends Web3 {
  private static instance: SharedWeb3;

  private constructor() {
    super(Web3.givenProvider);
  }

  /**
   * Keep just one instance of each subclass around.
   */
  public static getInstance(): SharedWeb3 {
    if (!SharedWeb3.instance) {
      SharedWeb3.instance = new SharedWeb3();
    }

    return SharedWeb3.instance;
  }
}

/**
 * The class defines the `getInstance` method that lets clients access
 * the unique singleton instance.
 *
 * Reference: https://refactoring.guru/design-patterns/singleton/typescript/example
 */
export class SharedWwkf {
  private static instance: Contract;

  private constructor() {
    const wwkfAbi = wwkfJson as AbiItem[];
    SharedWwkf.instance = new (SharedWeb3.getInstance().eth.Contract)(wwkfAbi, C.WWKF_CONTRACT.ADDR);
  }

  /**
   * Keep just one instance of each subclass around.
   */
  public static getInstance(): Contract {
    if (!SharedWwkf.instance) {
      new SharedWwkf();
    }

    return SharedWwkf.instance;
  }
}
