import type { WrapKeyRequest, WrapKeyResp } from "../types/schema.js";

export function wrapKey(req: WrapKeyRequest): Promise<WrapKeyResp> {
  return post("/api/wrap-key", req);
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
