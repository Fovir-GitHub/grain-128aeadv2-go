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

function getInputElementById(id: string): HTMLInputElement {
  return document.getElementById(id) as HTMLInputElement;
}

function getButtonElementById(id: string): HTMLButtonElement {
  return document.getElementById(id) as HTMLButtonElement;
}
