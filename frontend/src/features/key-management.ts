import { wrapKey } from "../lib/api.js";
import { onClick } from "../lib/dom.js";
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
} from "../lib/utils.js";

export function RegisterKeyManagementEvents() {
  onClick(els.generateKeyButton, handleGenerateKey);
  onClick(els.wrapSaveKeyButton, handleWrapSaveKey);
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
