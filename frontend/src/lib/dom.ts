// Register a click event.
export function onClick(el: HTMLElement, handler: () => void) {
  el.addEventListener("click", handler);
}

// Register a change event.
export function onChange(el: HTMLInputElement, handler: () => void) {
  el.addEventListener("change", handler);
}
