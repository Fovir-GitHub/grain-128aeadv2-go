import { GenerateHex } from "./utils.js";

export function RegisterNonceManagementEvents() {
  registerGenerateKeyClick();
}

// handleGenerateKeyClick fills the nonce input box
// with a random 96-bit (12 bytes) hex string
// with a prefix `0x`.
function handleGenerateKeyClick() {
  const input = document.getElementById(
    "nonce-management-input",
  ) as HTMLInputElement;
  const key = "0x" + GenerateHex(12);
  input.value = key;
}

function registerGenerateKeyClick() {
  const btn = document.getElementById(
    "nonce-management-generate-key",
  ) as HTMLButtonElement;
  btn.addEventListener("click", handleGenerateKeyClick);
}
