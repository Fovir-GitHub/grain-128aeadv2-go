import { encrypt } from "../lib/api.js";
import { onClick } from "../lib/dom.js";
import {
  els,
  getCipherInput,
  getKey,
  getNonce,
  isCipehrInputHex,
  setCipherOutput,
  setInitLFSR,
  setInitNFSR,
  setLoadedLFSR,
  setLoadedNFSR,
} from "../lib/elements.js";

export function RegisterCipherEvents() {
  onClick(els.encryptButton, handleEncrypt);
}

async function handleEncrypt() {
  const data = await encrypt({
    key: getKey(),
    nonce: getNonce(),
    plaintext: getCipherInput(),
    isInputHex: isCipehrInputHex(),
  });

  setCipherOutput(data.output);
  setLoadedLFSR(data.loadedLFSR);
  setLoadedNFSR(data.loadedNFSR);
  setInitLFSR(data.initLFSR);
  setInitNFSR(data.initNFSR);
}
