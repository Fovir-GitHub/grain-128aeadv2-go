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

function handleGenerateKey() {
  setKey("0x" + generateHex(16));
}

async function handleWrapSaveKey() {
  const data = await wrapKey({
    key: getKey(),
    password: getPassword(),
    ad: getAD(),
    isHex: isADHex(),
  });
  const blob = createBlobWithFileContent(data.key);
  downloadBlobFile(blob, "wrapped.key");
}

async function handleLoadKeyFile() {
  const b64 = await readTextFileContent(els.loadKeyFile);
  const data = await unwrapKey({
    b64: b64,
    passphrase: getPassword(),
    ad: getAD(),
    isHex: isADHex(),
  });
  setKey(data.key);
}
