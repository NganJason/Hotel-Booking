import { checkAvailabilityHandler } from "./handler.module.js";

let checkAvailabilityButton = document.getElementById(
  "check-availability-button"
);

checkAvailabilityButton.addEventListener("click", checkAvailabilityHandler);
