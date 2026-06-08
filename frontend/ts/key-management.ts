// Generate a random 128-bit (16 bytes) key.
function generateKey(): string {
  const bytes = new Uint8Array(16);
  crypto.getRandomValues(bytes);
  return Array.from(bytes)
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("");
}

// Handle the click event on `Generate Key` button.
function handleGenerateKeyClick() {
  const input = document.getElementById(
    "key-management-key-input",
  ) as HTMLInputElement;
  const key = "0x" + generateKey();
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
  const url = URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = "wrapped.key";
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
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
}
