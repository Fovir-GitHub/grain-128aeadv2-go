import { DownloadBlobFile, ReadTextFileContent } from "./utils.js";

export function RegisterCipherEvents() {
  registerEncryptClick();
  registerPlaintextUploadChange();
  registerSaveEncFileClick();
}

async function handleEncryptClick() {
  const getInputValueById = (id: string) => {
    return (document.getElementById(id) as HTMLInputElement).value;
  };

  const key = getInputValueById("key-management-key-input");
  const nonce = getInputValueById("nonce-management-input");
  const plaintext = getInputValueById("cipher-input");
  const isInputHex = (
    document.getElementById("cipher-input-mode-hex") as HTMLInputElement
  ).checked;

  const resp = await fetch("/api/encrypt", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ key, nonce, plaintext, isInputHex }),
  });

  if (!resp.ok) {
    const err = await resp.json();
    alert(err.msg);
    return;
  }

  const content = await resp.json();

  const output = content.output;
  const loadedLFSR = content.loadedLFSR;
  const loadedNFSR = content.loadedNFSR;
  const initLFSR = content.initLFSR;
  const initNFSR = content.initNFSR;

  updateResult(output, loadedLFSR, loadedNFSR, initLFSR, initNFSR);
}

function registerEncryptClick() {
  const btn = document.getElementById(
    "cipher-encrypt",
  ) as HTMLButtonElement;
  btn.addEventListener("click", handleEncryptClick);
}

function updateResult(
  output: string,
  loadedLFSR: string,
  loadedNFSR: string,
  initLFSR: string,
  initNFSR: string,
) {
  const updateLiContent = (id: string, content: string) => {
    const li = document.getElementById(id) as HTMLDataListElement;
    li.innerText = content;
  };

  const idContentMap = new Map<string, string>([
    ["cipher-output", output],
    ["cipher-loaded-lfsr", loadedLFSR],
    ["cipher-loaded-nfsr", loadedNFSR],
    ["cipher-init-lfsr", initLFSR],
    ["cipher-init-nfsr", initNFSR],
  ]);

  for (const [k, v] of idContentMap) {
    updateLiContent(k, v);
  }
}

async function handlePlaintextFileUpload() {
  const input = document.getElementById(
    "cipher-plaintext-file",
  ) as HTMLInputElement;
  const plaintextInput = document.getElementById(
    "cipher-input",
  ) as HTMLInputElement;

  const plaintext = await ReadTextFileContent(input);
  plaintextInput.value = plaintext;
}

function registerPlaintextUploadChange() {
  const input = document.getElementById(
    "cipher-plaintext-file",
  ) as HTMLInputElement;

  input.addEventListener("change", handlePlaintextFileUpload);
}

function handleSaveEncFileClick() {
  const input = document.getElementById(
    "cipher-output",
  ) as HTMLInputElement;
  const content = input.value;
  if (!content) {
    alert("Empty output");
    return;
  }

  const blob = new Blob([content], { type: "text/plain" });
  DownloadBlobFile(blob, "ciphertext.enc");
}

function registerSaveEncFileClick() {
  const btn = document.getElementById(
    "cipher-save-enc-file",
  ) as HTMLButtonElement;
  btn.addEventListener("click", handleSaveEncFileClick);
}
