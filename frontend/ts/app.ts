import { RegisterKeyManagementEvents } from "./key-management.js";

async function loadComponent(id: string, file: string) {
  const res = await fetch(file);
  const el = document.getElementById(id);
  if (el) {
    el.innerHTML = await res.text();
  }
}

async function main() {
  await loadComponent(
    "key-management",
    "components/key-management.html",
  );

  RegisterKeyManagementEvents();
}

main();
