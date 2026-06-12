import { onClick } from "../lib/dom.js";
import { els, setNonce } from "../lib/elements.js";
import { generateHex } from "../lib/utils.js";

export function RegisterNonceManagementEvents() {
  onClick(els.generateIVButton, handleGenerateIV);
}

export function handleGenerateIV() {
  setNonce("0x" + generateHex(12));
}
