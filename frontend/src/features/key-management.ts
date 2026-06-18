import { unwrapKey, wrapKey } from "../lib/api.js";
import { onChange, onClick } from "../lib/dom.js";
import {
  els,
  getAD,
  getKey,
  getPassword,
  isADHex,
  setKey,
} from "../lib/elements.js";
import {
  createBlobWithFileContent,
  downloadBlobFile,
  generateHex,
  readTextFileContent,
} from "../lib/utils.js";

export function RegisterKeyManagementEvents() {
  onClick(els.generateKeyButton, handleGenerateKey);
  onClick(els.wrapSaveKeyButton, handleWrapSaveKey);
  onChange(els.loadKeyFile, handleLoadKeyFile);
}

// Fill key field with `0x` and a random 16-byte (128-bit) hex string.
function handleGenerateKey() {
  setKey("0x" + generateHex(16));
}

// Wrap the key with given password and associated data and store in `wrapped.key`.
async function handleWrapSaveKey() {
  // Call API.
  const data = await wrapKey({
    key: getKey(),
    password: getPassword(),
    ad: getAD(),
    isHex: isADHex(),
  });
  const blob = createBlobWithFileContent(data.key);
  downloadBlobFile(blob, "wrapped.key");
}

// Load a `.key` file and unwrap it.
async function handleLoadKeyFile() {
  const content = await readTextFileContent(els.loadKeyFile);
  const data = await unwrapKey({
    content: content,
    passphrase: getPassword(),
    ad: getAD(),
    isHex: isADHex(),
  });
  setKey(data.key);
}
