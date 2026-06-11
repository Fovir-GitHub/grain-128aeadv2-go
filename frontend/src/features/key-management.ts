import { onClick } from "../lib/dom.js";
import { els, setKey } from "../lib/elements.js";
import { generateHex } from "../lib/utils.js";

export function RegisterKeyManagementEvents() {
  onClick(els.generateKeyButton, handleGenerateKey);
}

function handleGenerateKey() {
  setKey("0x" + generateHex(16));
}
