/**
 * Function that logs errors.
 * May in the future be overwriten to function that display error on screen.
 */
export const error = (s: string) => alert(s);

export const csrfToken =
  document?.querySelector('meta[name="csrf-token"]')?.getAttribute("content") ||
  "";
