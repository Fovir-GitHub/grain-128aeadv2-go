// Generate a random hex string with a given length in byte.
export function GenerateHex(byte: number): string {
  const bytes = new Uint8Array(byte);
  crypto.getRandomValues(bytes);
  return Array.from(bytes)
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("");
}

export async function ReadTextFileContent(
  input: HTMLInputElement,
): Promise<string> {
  const file = input.files?.[0];
  if (!file) {
    throw new Error("no file selected");
  }
  return await file.text();
}

export function DownloadBlobFile(blob: Blob, filename: string) {
  const url = URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = filename;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
}

export function Hex2String(hex: string): string {
  hex = hex.replace(/^0x/, "");

  const bytes = new Uint8Array(hex.length / 2);
  for (let i = 0; i < hex.length; i += 2) {
    bytes[i / 2] = parseInt(hex.slice(i, i + 2), 16);
  }

  return new TextDecoder().decode(bytes);
}

export function createBlobWithFileContent(content: string): Blob {
  return new Blob([content], { type: "text/plain" });
}
