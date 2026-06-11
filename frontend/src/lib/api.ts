import type {
  DecryptionRequest,
  DecryptionResp,
  EncryptionRequest,
  EncryptionResp,
  UnwrapKeyRequest,
  UnwrapKeyResp,
  WrapKeyRequest,
  WrapKeyResp,
} from "../types/schema.js";

export function wrapKey(req: WrapKeyRequest): Promise<WrapKeyResp> {
  return post("/api/wrap-key", req);
}

export function unwrapKey(
  req: UnwrapKeyRequest,
): Promise<UnwrapKeyResp> {
  return post("/api/unwrap-key", req);
}

export function encrypt(
  req: EncryptionRequest,
): Promise<EncryptionResp> {
  return post("/api/encrypt", req);
}

export function decrypt(
  req: DecryptionRequest,
): Promise<DecryptionResp> {
  return post("/api/decrypt", req);
}

async function post<T>(path: string, body: unknown): Promise<T> {
  const resp = await fetch(path, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
  if (!resp.ok) {
    const err = await resp.json();
    alert(err.msg);
    throw new Error(err.msg);
  }
  return resp.json() as Promise<T>;
}
