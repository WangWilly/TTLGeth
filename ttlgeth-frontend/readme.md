# TTLGETH frontend

## Reviews
```bash
npm create vite@latest ttlgeth-frontend -- --template react-ts
```

- https://vitejs.dev/guide/static-deploy.html
- https://tailwindcss.com/docs/guides/vite
  - https://flaviocopes.com/fix-unknown-at-rule-tailwind/
  - https://stackoverflow.com/questions/58139175/running-actions-in-another-directory

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
(window as any).ethereum
```

- Use a //@ts-ignore comment to suppress the error on the line where you access the ethereum property. For example:

```ts
// @ts-ignore
window.ethereum
```

For more details, you can refer to these StackOverflow answers¹²³.

Source: Conversation with Bing, 4/27/2023(1) Property 'ethereum' does not exist on type 'Window & typeof globalThis' in Next Typescript project. https://ethereum.stackexchange.com/questions/135989/property-ethereum-does-not-exist-on-type-window-typeof-globalthis-in-next Accessed 4/27/2023.
(2) TypeScript error: Property 'X' does not exist on type 'Window'. https://stackoverflow.com/questions/56457935/typescript-error-property-x-does-not-exist-on-type-window Accessed 4/27/2023.
(3) javascript - typescript ignore window unknown property - Stack Overflow. https://stackoverflow.com/questions/71042635/typescript-ignore-window-unknown-property Accessed 4/27/2023.
