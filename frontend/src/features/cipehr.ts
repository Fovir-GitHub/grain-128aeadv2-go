import { encrypt } from "../lib/api.js";
import { onChange, onClick } from "../lib/dom.js";
import {
  els,
  getCipherInput,
  getCipherOutput,
  getKey,
  getNonce,
  isCipehrInputHex,
  setCipherInput,
  setCipherOutput,
  setInitLFSR,
  setInitNFSR,
  setLoadedLFSR,
  setLoadedNFSR,
} from "../lib/elements.js";
import {
  createBlobWithFileContent,
  downloadBlobFile,
  readTextFileContent,
} from "../lib/utils.js";

export function RegisterCipherEvents() {
  onClick(els.encryptButton, handleEncrypt);
  onChange(els.loadEncFile, handleLoadEncFile);
  onClick(els.saveEncFileButton, handleSaveEncFile);
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

async function handleLoadEncFile() {
  const plaintext = await readTextFileContent(els.loadEncFile);
  setCipherInput(plaintext);
}

function handleSaveEncFile() {
  const ciphertext = getCipherOutput();
  if (!ciphertext) {
    alert("Empty output");
    return;
  }
  const blob = createBlobWithFileContent(ciphertext);
  downloadBlobFile(blob, "ciphertext.enc");
}
