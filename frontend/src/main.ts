import { RegisterCipherEvents } from "./features/cipehr.js";
import { RegisterKeyManagementEvents } from "./features/key-management.js";
import { RegisterNonceManagementEvents } from "./features/nonce-management.js";

// Entry point.
// Register all events.
function main() {
  RegisterKeyManagementEvents();
  RegisterNonceManagementEvents();
  RegisterCipherEvents();
}

main();
