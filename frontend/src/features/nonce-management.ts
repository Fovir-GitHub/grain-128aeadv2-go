import { onClick } from "../lib/dom.js";
import { els, setNonce } from "../lib/elements.js";
import { generateHex } from "../lib/utils.js";

// TODO: Handle automatically generate unique IV.

export function RegisterNonceManagementEvents() {
  onClick(els.generateIVButton, handleGenerateIV);
}

function handleGenerateIV() {
  setNonce("0x" + generateHex(12));
}
