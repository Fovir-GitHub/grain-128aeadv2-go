export const els = {
  key: getInputElementById("key-management-key-input"),
  password: getInputElementById("key-management-password"),
  ad: getInputElementById("key-management-ad"),
  adHex: getInputElementById("key-management-input-mode-hex"),
  generateKeyButton: getButtonElementById(
    "key-management-generate-key",
  ),
  wrapSaveKeyButton: getButtonElementById(
    "key-management-wrap-save-key",
  ),
  loadKeyFile: getInputElementById("key-management-load-key-file"),
  nonce: getInputElementById("nonce-management-input"),
  generateIVButton: getButtonElementById(
    "nonce-management-generate-key",
  ),
  cipherInput: getInputElementById("cipher-input"),
  cipherInputHex: getInputElementById("cipher-input-mode-hex"),
  cipherOutput: getTextAreaElementById("cipher-output"),
  loadedLFSR: getTextAreaElementById("cipher-loaded-lfsr"),
  loadedNFSR: getTextAreaElementById("cipher-loaded-nfsr"),
  initLFSR: getTextAreaElementById("cipher-init-lfsr"),
  initNFSR: getTextAreaElementById("cipher-init-nfsr"),
  encryptButton: getButtonElementById("cipher-encrypt"),
  decryptButton: getButtonElementById("cipher-decrypt"),
  saveEncFileButton: getButtonElementById("cipher-save-enc-file"),
  loadEncFile: getInputElementById("cipher-plaintext-file"),
  plaintextOutputText: getTextAreaElementById(
    "cipher-plaintext-output-text",
  ),
};

export function getKey(): string {
  return els.key.value;
}

export function setKey(k: string) {
  els.key.value = k;
}

export function getPassword(): string {
  return els.password.value;
}

export function setPassword(p: string) {
  els.password.value = p;
}

export function getAD(): string {
  return els.ad.value;
}

export function setAD(a: string) {
  els.ad.value = a;
}

export function isADHex(): boolean {
  return els.adHex.checked;
}

export function getNonce(): string {
  return els.nonce.value;
}

export function setNonce(n: string) {
  els.nonce.value = n;
}

export function getCipherInput(): string {
  return els.cipherInput.value;
}

export function setCipherInput(ci: string) {
  els.cipherInput.value = ci;
}

export function isCipehrInputHex(): boolean {
  return els.cipherInputHex.checked;
}

export function getCipherOutput(): string {
  return els.cipherOutput.value;
}

export function setCipherOutput(co: string) {
  els.cipherOutput.value = co;
}

export function setLoadedLFSR(ll: string) {
  els.loadedLFSR.value = ll;
}

export function setLoadedNFSR(ln: string) {
  els.loadedNFSR.value = ln;
}

export function setInitLFSR(il: string) {
  els.initLFSR.value = il;
}

export function setInitNFSR(_in: string) {
  els.initNFSR.value = _in;
}

export function setPlaintextOutputText(pot: string) {
  els.plaintextOutputText.value = pot;
}

function getInputElementById(id: string): HTMLInputElement {
  return document.getElementById(id) as HTMLInputElement;
}

function getButtonElementById(id: string): HTMLButtonElement {
  return document.getElementById(id) as HTMLButtonElement;
}

function getTextAreaElementById(id: string): HTMLTextAreaElement {
  return document.getElementById(id) as HTMLTextAreaElement;
}
