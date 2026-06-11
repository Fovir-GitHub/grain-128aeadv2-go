import { RegisterKeyManagementEvents } from "./features/key-management.js";
import { RegisterNonceManagementEvents } from "./features/nonce-management.js";

function main() {
  RegisterKeyManagementEvents();
  RegisterNonceManagementEvents();
}

main();
