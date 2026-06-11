export function onClick(el: HTMLElement, handler: () => void) {
  el.addEventListener("click", handler);
}

export function onChange(el: HTMLInputElement, handler: () => void) {
  el.addEventListener("change", handler);
}
