import {
  DownloadBlobFile,
  GenerateHex,
  ReadTextFileContent,
} from "./utils.js";

// Handle the click event on `Generate Key` button.
function handleGenerateKeyClick() {
  const input = document.getElementById(
    "key-management-key-input",
  ) as HTMLInputElement;
  const key = "0x" + GenerateHex(16);
  input.value = key;
}

// Handle the click event on `Wrap & Save .key File` button.
async function handleWrapSaveKeyClick() {
  // A tool function.
  const getInputValueById = (id: string) => {
    return (document.getElementById(id) as HTMLInputElement).value;
  };

  // Get required DOM values.
  const key = getInputValueById("key-management-key-input");
  const password = getInputValueById("key-management-password");
  const ad = getInputValueById("key-management-ad");
  const isHex = (
    document.getElementById(
      "key-management-input-mode-hex",
    ) as HTMLInputElement
  ).checked;

  // Communicate with backend.
  const resp = await fetch("/api/wrap-key", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ key, password, ad, isHex }),
  });

  if (!resp.ok) {
    const err = await resp.json();
    alert(err.msg);
    return;
  }

  // Create blob and download it.
  const blob = await resp.blob();
  DownloadBlobFile(blob, "wrapped.key");
}

async function handleLoadKeyFileClick() {
  const getInputValueById = (id: string) => {
    return (document.getElementById(id) as HTMLInputElement).value;
  };

  const input = document.getElementById(
    "key-management-load-key-file",
  ) as HTMLInputElement;

  const b64 = await ReadTextFileContent(input);
  const passphrase = getInputValueById("key-management-password");
  const ad = getInputValueById("key-management-ad");
  if (!passphrase) {
    return;
  }

  const resp = await fetch("/api/unwrap-key", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ b64, passphrase, ad }),
  });
  if (!resp.ok) {
    const err = await resp.json();
    alert(err.msg);
    return;
  }

  const content = await resp.json();
  const keyField = document.getElementById(
    "key-management-key-input",
  ) as HTMLInputElement;
  keyField.value = content.key;
}

function registerLoadKeyFileChange() {
  const input = document.getElementById(
    "key-management-load-key-file",
  ) as HTMLInputElement;
  input.addEventListener("change", handleLoadKeyFileClick);
}

function registerGenerateKeyClick() {
  const btn = document.getElementById(
    "key-management-generate-key",
  ) as HTMLButtonElement;
  btn.addEventListener("click", handleGenerateKeyClick);
}

function registerWrapSaveKeyClick() {
  const btn = document.getElementById(
    "key-management-wrap-save-key",
  ) as HTMLButtonElement;
  btn.addEventListener("click", handleWrapSaveKeyClick);
}

// Register all key management events.
export function RegisterKeyManagementEvents() {
  registerGenerateKeyClick();
  registerWrapSaveKeyClick();
  registerLoadKeyFileChange();
}
