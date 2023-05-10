import Web3 from "web3";

export default {
  WWKF_CONTRACT: {
    ADDR: "0x03378DAa43739f2361FE67175aD6bF2666309748",
    BEGIN_BLOCK_ID: 3296031,
    TOPICS: {
      TRANSFER: Web3.utils.sha3("Transfer(address,address,uint256)") ?? "",
    },
  },
};
