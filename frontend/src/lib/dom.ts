export function onClick(el: HTMLElement, handler: () => void) {
  el.addEventListener("click", handler);
}
