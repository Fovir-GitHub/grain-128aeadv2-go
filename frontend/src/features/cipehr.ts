import { decrypt, encrypt } from "../lib/api.js";
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
  setPlaintextOutputText,
} from "../lib/elements.js";
import {
  createBlobWithFileContent,
  downloadBlobFile,
  hex2string,
  readTextFileContent,
} from "../lib/utils.js";

export function RegisterCipherEvents() {
  onClick(els.encryptButton, handleEncrypt);
  onClick(els.decryptButton, handleDecrypt);
  onChange(els.loadPlaintextFile, handleLoadPlaintextFile);
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

async function handleLoadPlaintextFile() {
  const plaintext = await readTextFileContent(els.loadPlaintextFile);
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

async function handleDecrypt() {
  const data = await decrypt({
    key: getKey(),
    ciphertext: getCipherInput(),
  });

  setCipherOutput(data.output);
  setLoadedLFSR(data.loadedLFSR);
  setLoadedNFSR(data.loadedNFSR);
  setInitLFSR(data.initLFSR);
  setInitNFSR(data.initNFSR);
  setPlaintextOutputText(hex2string(data.output));
}
