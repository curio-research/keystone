// ------------------------------------------------
// under dev mode we prefund accounts
// ------------------------------------------------

import { providers, utils } from "ethers";
import fs from "fs";

const fund = async () => {
  try {
    console.log("💶 Prefunding accounts");

    var prefundArr: Record<string, Balance> = JSON.parse(fs.readFileSync("./scripts/prefundDevAccounts.json", "utf-8"));

    const provider = new providers.JsonRpcProvider("http://127.0.0.1:8540");
    const listAccounts = await provider.listAccounts();
    const originAcc = listAccounts[0];

    const originSigner = provider.getSigner(originAcc);

    // fund accounts
    for (const [address, value] of Object.entries(prefundArr)) {
      await originSigner.sendTransaction({
        from: originSigner._address,
        to: address,
        value: utils.parseEther(value.balance),
      });
    }
  } catch (err) {
    console.log(err);
  }
};

fund();

interface Balance {
  balance: string;
}
