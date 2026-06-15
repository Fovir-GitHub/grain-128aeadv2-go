import { onClick } from "../lib/dom.js";
import { els, setNonce } from "../lib/elements.js";
import { generateHex } from "../lib/utils.js";

export function RegisterNonceManagementEvents() {
  onClick(els.generateIVButton, handleGenerateIV);
}

// Generate a random IV in the format of `0x` + 12-byte (96-bit) hex.
export function handleGenerateIV() {
  setNonce("0x" + generateHex(12));
}
