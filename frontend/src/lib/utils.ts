// Generate a random hex string with a given length in byte.
export function generateHex(byte: number): string {
  const bytes = new Uint8Array(byte);
  crypto.getRandomValues(bytes);
  return Array.from(bytes)
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("");
}

// Read file content as text when the user uploads a file.
export async function readTextFileContent(
  input: HTMLInputElement,
): Promise<string> {
  const file = input.files?.[0];
  if (!file) {
    throw new Error("no file selected");
  }
  return await file.text();
}

// Download a blob file named `filename`.
export function downloadBlobFile(blob: Blob, filename: string) {
  const url = URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = filename;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
}

// Convert hex format string to text string.
export function hex2string(hex: string): string {
  hex = hex.replace(/^0x/, "");

  const bytes = new Uint8Array(hex.length / 2);
  for (let i = 0; i < hex.length; i += 2) {
    bytes[i / 2] = parseInt(hex.slice(i, i + 2), 16);
  }

  return new TextDecoder().decode(bytes);
}

// Create a blob object with plain text format content.
export function createBlobWithFileContent(content: string): Blob {
  return new Blob([content], { type: "text/plain" });
}

export function prepend0xToHex(s: string): string {
  if (!s.toLowerCase().startsWith("0x")) {
    s = "0x" + s;
  }
  return s;
}
