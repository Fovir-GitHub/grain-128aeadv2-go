// Generate a random hex string with a given length in byte.
export function GenerateHex(byte: number): string {
  const bytes = new Uint8Array(byte);
  crypto.getRandomValues(bytes);
  return Array.from(bytes)
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("");
}
