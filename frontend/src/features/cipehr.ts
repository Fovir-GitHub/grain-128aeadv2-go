import { decrypt, encrypt } from "../lib/api.js";
import { onChange, onClick } from "../lib/dom.js";
import {
  els,
  enabledAutoGenerateIV,
  getCipherInput,
  getCipherOutput,
  getKey,
  getNonce,
  getPlaintextOutputText,
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
import { handleGenerateIV } from "./nonce-management.js";

export function RegisterCipherEvents() {
  onClick(els.encryptButton, handleEncrypt);
  onClick(els.decryptButton, handleDecrypt);
  onChange(els.loadPlaintextFile, handleLoadPlaintextFile);
  onClick(els.saveEncFileButton, handleSaveEncFile);
  onClick(els.saveDecFileButton, handleSaveDecFile);
  onChange(els.loadEncFile, handleLoadEncFile);
}

// Handle encryption request.
async function handleEncrypt() {
  // Generate IV if automatically generate IV is enabled.
  if (enabledAutoGenerateIV()) {
    handleGenerateIV();
  }

  // Call API.
  const data = await encrypt({
    key: getKey(),
    nonce: getNonce(),
    plaintext: getCipherInput(),
    isInputHex: isCipehrInputHex(),
  });

  // Fill in updated data and clear the plaintext of text format output.
  setPlaintextOutputText("");
  setCipherOutput(data.output);
  setLoadedLFSR(data.loadedLFSR);
  setLoadedNFSR(data.loadedNFSR);
  setInitLFSR(data.initLFSR);
  setInitNFSR(data.initNFSR);
}

// Handle loading plaintext file.
async function handleLoadPlaintextFile() {
  const plaintext = await readTextFileContent(els.loadPlaintextFile);
  setCipherInput(plaintext);
}

// Save ciphertext into `ciphertext.enc` file.
function handleSaveEncFile() {
  const ciphertext = getCipherOutput();
  if (!ciphertext) {
    alert("Empty output");
    return;
  }
  const blob = createBlobWithFileContent(ciphertext);
  downloadBlobFile(blob, "ciphertext.enc");
}

// Handle decryption request.
async function handleDecrypt() {
  // Call API.
  const data = await decrypt({
    key: getKey(),
    ciphertext: getCipherInput(),
  });

  // Fill in values.
  setCipherOutput(data.output);
  setLoadedLFSR(data.loadedLFSR);
  setLoadedNFSR(data.loadedNFSR);
  setInitLFSR(data.initLFSR);
  setInitNFSR(data.initNFSR);
  setPlaintextOutputText(hex2string(data.output));
}

// Loading `.enc` files.
async function handleLoadEncFile() {
  const ciphertext = await readTextFileContent(els.loadEncFile);
  setCipherInput(ciphertext);
}

// Save plaintext into a `plaintext.dec` file.
// The file is in JSON format, which contains both hex and text format of plaintext.
function handleSaveDecFile() {
  // If the text format of plaintext is empty, it is not a decryption.
  if (getPlaintextOutputText() == "") {
    alert("no decryption");
    return;
  }

  // Create JSON string.
  const content = JSON.stringify({
    hex: getCipherOutput(),
    text: getPlaintextOutputText(),
  });

  const blob = createBlobWithFileContent(content);
  downloadBlobFile(blob, "plaintext.dec");
}
