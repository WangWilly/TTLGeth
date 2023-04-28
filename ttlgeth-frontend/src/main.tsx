import React from "react";
import ReactDOM from "react-dom/client";
import WalletView from "./WalletView.tsx";
import EventLogs from "./EventLogs.tsx";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <div className="container mx-9 px-4">
      <p className="text-base">
        Install <a href="https://metamask.io/">Metamask</a> first then click "connect wallet" to check your{" "}
        <a href="https://sepolia.etherscan.io/token/0x03378daa43739f2361fe67175ad6bf2666309748">WWKF</a> token.
      </p>
      <div className="columns-2">
        <div className="break-after-column">
          <WalletView />
        </div>
        <div>
          <EventLogs />
        </div>
      </div>
    </div>
  </React.StrictMode>
);
