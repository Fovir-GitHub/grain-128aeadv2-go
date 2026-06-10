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
