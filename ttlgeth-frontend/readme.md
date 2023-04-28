# TTLGETH frontend

## Reviews

```bash
npm create vite@latest ttlgeth-frontend -- --template react-ts
```

- https://vitejs.dev/guide/static-deploy.html
  - 📌 https://ithelp.ithome.com.tw/articles/10200767
  - https://dev.to/colocodes/how-to-debug-a-react-app-51l4
- https://tailwindcss.com/docs/guides/vite
  - https://flaviocopes.com/fix-unknown-at-rule-tailwind/
  - https://www.makeuseof.com/tailwind-css-in-react/
- lintting and formatting
  - https://www.digitalocean.com/community/tutorials/linting-and-formatting-with-eslint-in-vs-code
  - https://www.alphr.com/use-prettier-vs-code/
  - https://www.robinwieruch.de/how-to-use-prettier-vscode/
- Web3JS
  - https://ethereum.stackexchange.com/questions/2024/how-to-access-the-event-log-by-knowing-the-contract-address-web3
- nit.
  - https://stackoverflow.com/questions/65504958/web3-js-extending-the-window-interface-type-definitions

```ts
// ttlgeth-frontend/vite.config.ts
export default defineConfig({
  plugins: [react()],
  base: "./", // 📌
});
```

## 🤖 Robot summary

Quick accession of network result.

### How to use vite to create static HTML of typescript react?

Based on web search results, you can use Vite to create static HTML of TypeScript React by following these steps:

- Install Vite as a dev dependency in your project and set up the npm scripts for building and previewing your site. You can use the `vite build` command to generate a `dist` folder with the static HTML and assets. You can also use the `vite preview` command to test the build locally.
- Set the correct base in `vite.config.js`. If you are deploying to a subdirectory, you need to specify the base path as `/<SUBDIRECTORY>/`.
- Scaffold a Vite + React + TypeScript project using the `react-ts` template. You can use the `npm create vite@latest my-react-ts-app -- --template react-ts` command to do this. Alternatively, you can use other community templates that support TypeScript and React.
- Deploy the `dist` folder to any platform that supports static hosting, such as Vercel, GitHub Pages, Render, Netlify, etc. You may need to configure some additional settings depending on the platform.

For more details, you can refer to the official Vite documentation¹, this tutorial², and this StackOverflow answer³.

Source: Conversation with Bing, 4/27/2023(1) Deploying a Static Site | Vite. https://vitejs.dev/guide/static-deploy Accessed 4/27/2023.
(2) Getting Started | Vite. https://vitejs.dev/guide/ Accessed 4/27/2023.
(3) How to deploy React TypeScript app to Vercel with Vite. https://stackoverflow.com/questions/72001670/how-to-deploy-react-typescript-app-to-vercel-with-vite Accessed 4/27/2023.

### How to apply tailwind css in a vite project?

Based on web search results, you can load an ABI file in TypeScript by following these steps:

- Install a library that can read and parse ABI files, such as web3.js or ethers.js. You can use npm or yarn to install them as dependencies in your project.
- Import the ABI file as a JSON module in your TypeScript file. You may need to enable the `resolveJsonModule` and `esModuleInterop` options in your `tsconfig.json` file for this to work. For example, `import MyContractABI from './MyContract.json'`.
- Import the type definitions for the ABI from the library you are using. For example, `import { AbiItem } from 'web3-utils'` or `import { ContractInterface } from 'ethers'`.
- Cast the imported ABI file to the appropriate type. For example, `const abi = MyContractABI as AbiItem[]` or `const abi = MyContractABI as ContractInterface`.
- Use the ABI file to interact with your contract using the library methods. For example, `const contract = new web3.eth.Contract(abi, address)` or `const contract = new ethers.Contract(address, abi, provider)`.

For more details, you can refer to these StackOverflow answers¹²³.

Source: Conversation with Bing, 4/27/2023(1) web3js - web3.js import contract abi from .json file - Ethereum Stack Exchange. https://ethereum.stackexchange.com/questions/21665/web3-js-import-contract-abi-from-json-file Accessed 4/27/2023.
(2) Trouble with web3.eth.Contract, ABI usage with Typescript. https://ethereum.stackexchange.com/questions/94601/trouble-with-web3-eth-contract-abi-usage-with-typescript Accessed 4/27/2023.
(3) How to import ABI types from ethers.js typescript. https://stackoverflow.com/questions/73109666/how-to-import-abi-types-from-ethers-js-typescript Accessed 4/27/2023.

### How to solve "Property 'ethereum' does not exist on type 'Window & typeof globalThis'."?

Based on web search results, you can solve the error "Property 'ethereum' does not exist on type 'Window & typeof globalThis'" by following these methods:

- Declare a global interface for Window that includes the ethereum property. You can do this in a separate declaration file or at the top of your TypeScript file. For example:

```ts
declare global {
  interface Window {
    ethereum: any;
  }
}
```

- Use a type assertion to cast window to any before accessing the ethereum property. For example:

```ts
(window as any).ethereum;
```

- Use a //@ts-ignore comment to suppress the error on the line where you access the ethereum property. For example:

```ts
// @ts-ignore
window.ethereum;
```

For more details, you can refer to these StackOverflow answers¹²³.

Source: Conversation with Bing, 4/27/2023(1) Property 'ethereum' does not exist on type 'Window & typeof globalThis' in Next Typescript project. https://ethereum.stackexchange.com/questions/135989/property-ethereum-does-not-exist-on-type-window-typeof-globalthis-in-next Accessed 4/27/2023.
(2) TypeScript error: Property 'X' does not exist on type 'Window'. https://stackoverflow.com/questions/56457935/typescript-error-property-x-does-not-exist-on-type-window Accessed 4/27/2023.
(3) javascript - typescript ignore window unknown property - Stack Overflow. https://stackoverflow.com/questions/71042635/typescript-ignore-window-unknown-property Accessed 4/27/2023.

### Show me an example of subscribing to logs events

Here is an example of subscribing to logs events using web3js, based on web search results¹:

```javascript
// create a web3 instance with a provider
const Web3 = require("web3");
const web3 = new Web3("https://mainnet.infura.io/v3/endpoint");

// create a contract object with the ABI and address
const abi = [ ... ]; // the ABI of the contract
const address = "0x123456..."; // the address of the contract
const contract = new web3.eth.Contract(abi, address);

// subscribe to logs events with optional parameters
var subscription = web3.eth.subscribe('logs', {
  address: contract.options.address,
  topics: [ ... ] // the topics to filter by
}, function(error, result) {
  if (!error) {
    // do something with the log object
    console.log(result);
  }
});

// unsubscribe from the subscription
subscription.unsubscribe(function(error, success) {
  if (success) {
    console.log("Successfully unsubscribed!");
  }
});
```

Source: Conversation with Bing, 4/28/2023(1) web3.eth.subscribe — web3.js 1.0.0 documentation. https://web3js.readthedocs.io/en/v1.2.11/web3-eth-subscribe.html Accessed 4/28/2023.
(2) myContract.events.MyEvent() vs web3.eth.subscribe('logs', ...);. https://ethereum.stackexchange.com/questions/91646/mycontract-events-myevent-vs-web3-eth-subscribelogs Accessed 4/28/2023.
(3) web3js subscribe logs to fast for Javascript to handle. https://stackoverflow.com/questions/71890934/web3js-subscribe-logs-to-fast-for-javascript-to-handle Accessed 4/28/2023.
(4) error in code example of subscribing logs #1762 - Github. https://github.com/web3/web3.js/issues/1762 Accessed 4/28/2023.

### 📖 What is the usage of `componentDidMount`?

The `componentDidMount` method is a lifecycle method of a class component in React. **It is invoked after the component is mounted on the DOM.** It is a good place to 📌 **perform any side effects, such as fetching data, setting up subscriptions, or updating the state**. You can also access the props and state of the component in this method. Here is an example of using `componentDidMount`, based on web search results²:

```javascript
import React, { Component } from "react";

class App extends Component {
  constructor(props) {
    super(props);
    // 📌 state
    this.state = {
      users: [],
    };
  }

  // fetch some users data from an API
  componentDidMount() {
    // Promised receive
    fetch("https://jsonplaceholder.typicode.com/users")
      .then((response) => response.json())
      .then((data) => {
        // update the state with the fetched data
        this.setState({
          users: data,
        });
      });
  }

  render() {
    return (
      <div>
        <h1>Users</h1>
        <ul>
          {this.state.users.map((user) => (
            <li key={user.id}>
              <p>Name: {user.name}</p>
              <p>Email: {user.email}</p>
            </li>
          ))}
        </ul>
      </div>
    );
  }
}

export default App;
```

Source: Conversation with Bing, 4/28/2023(1) ReactJS componentDidMount() Method - GeeksforGeeks. https://www.geeksforgeeks.org/reactjs-componentdidmount-method/ Accessed 4/28/2023.
(2) Component – React. https://react.dev/reference/react/Component Accessed 4/28/2023.
(3) React Functional Component: how to use componentDidMount(). https://stackoverflow.com/questions/71434321/react-functional-component-how-to-use-componentdidmount Accessed 4/28/2023.
(4) React JS componentDidMount() Methods and Functional Components - KnowledgeHut. https://www.knowledgehut.com/blog/web-development/react-js-componentdidmount Accessed 4/28/2023.

### 📖 Can I show the retrieved event logs using react?

One possible way to show the retrieved event logs using react and web3js is to use the **componentDidMount method of your react component to subscribe to the logs events** and **then use the setState method to update the component state with the event data**. You can then render the event data in your component using JSX. Here is a simplified example, based on web search results²:

```javascript
import React, { Component } from 'react';
import Web3 from 'web3';

class EventLogs extends Component {
  constructor(props) {
    super(props);
    this.state = {
      web3: null,
      contract: null,
      logs: []
    };
  }

  componentDidMount() {
    // create a web3 instance with a provider
    const web3 = new Web3("https://mainnet.infura.io/v3/endpoint");
    this.setState({ web3 });

    // create a contract object with the ABI and address
    const abi = [ ... ]; // the ABI of the contract
    const address = "0x123456..."; // the address of the contract
    const contract = new web3.eth.Contract(abi, address);
    this.setState({ contract });

    // Subscribe to logs events with optional parameters; Promised results
    var subscription = web3.eth.subscribe('logs', {
      address: contract.options.address,
      topics: [ ... ] // the topics to filter by
    }, (error, result) => {
      if (!error) {
        // update the state with the new log object
        this.setState(prevState => ({
          logs: [...prevState.logs, result]
        }));
      }
    });
  }

  render() {
    return (
      <div>
        <h1>Event Logs</h1>
        <ul>
          {this.state.logs.map(log => (
            <li key={log.id}>
              {/* display some properties of the log object */}
              <p>Address: {log.address}</p>
              <p>Data: {log.data}</p>
              <p>Block Number: {log.blockNumber}</p>
            </li>
          ))}
        </ul>
      </div>
    );
  }
}

export default EventLogs;
```

Source: Conversation with Bing, 4/28/2023(1) Listening for events in a React DApp - Ethereum Stack Exchange. https://ethereum.stackexchange.com/questions/43260/listening-for-events-in-a-react-dapp Accessed 4/28/2023.
(2) React: Incorporating event listening into web3js promise. https://ethereum.stackexchange.com/questions/42800/react-incorporating-event-listening-into-web3js-promise Accessed 4/28/2023.
(3) web3js - Ethereum Event log - Ethereum Stack Exchange. https://ethereum.stackexchange.com/questions/46544/ethereum-event-log Accessed 4/28/2023.
(4) web3js - How to get only latest log from my event? - Ethereum Stack Exchange. https://ethereum.stackexchange.com/questions/41976/how-to-get-only-latest-log-from-my-event Accessed 4/28/2023.

Sure, here is an example of a typescript component with react, based on web search results¹:

```typescript
// import React and the types for props and state
import React, { Component } from "react";
import { Props, State } from "./types";

// define a class component that extends React.Component
class Counter extends Component<Props, State> {
  // define the initial state using the State type
  state: State = {
    count: 0,
  };

  // define a method to increment the count
  increment = () => {
    this.setState((prevState) => ({
      count: prevState.count + 1,
    }));
  };

  // define a method to decrement the count
  decrement = () => {
    this.setState((prevState) => ({
      count: prevState.count - 1,
    }));
  };

  // define a method to reset the count
  reset = () => {
    this.setState({
      count: 0,
    });
  };

  // render the component using JSX
  render() {
    return (
      <div>
        <h1>Counter</h1>
        <p>The current count is {this.state.count}</p>
        <button onClick={this.increment}>+</button>
        <button onClick={this.decrement}>-</button>
        <button onClick={this.reset}>Reset</button>
      </div>
    );
  }
}

export default Counter;
```

Source: Conversation with Bing, 4/28/2023(1) How to use TypeScript with React: A tutorial with examples. https://blog.logrocket.com/how-use-typescript-react-tutorial-examples/ Accessed 4/28/2023.
(2) Using TypeScript with React | DigitalOcean. https://www.digitalocean.com/community/tutorials/react-typescript-with-react Accessed 4/28/2023.
(3) Build a component library with React and TypeScript. https://blog.logrocket.com/build-component-library-react-typescript/ Accessed 4/28/2023.
