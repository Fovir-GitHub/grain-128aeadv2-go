function generateKey(): string {
  const bytes = new Uint8Array(16);
  crypto.getRandomValues(bytes);
  return Array.from(bytes)
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("");
}

function handleGenerateKeyClick() {
  const input = document.getElementById(
    "key-management-key-input",
  ) as HTMLInputElement;
  const key = "0x" + generateKey();
  input.value = key;
}

export function RegisterGenerateKeyClick() {
  const btn = document.getElementById(
    "key-management-generate-key",
  ) as HTMLButtonElement;
  btn.addEventListener("click", handleGenerateKeyClick);
}
