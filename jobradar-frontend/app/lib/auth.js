// Sesi login disimpan di localStorage (bukan cookie, backend tidak set cookie).
// Access token backend berumur pendek: 15 menit, dan tidak ada endpoint
// refresh yang aktif — jadi kita jujur ke pengguna soal ini alih-alih
// berpura-pura sesi bertahan lama.

const TOKEN_KEY = "jobradar_token";
const USER_KEY = "jobradar_user";
const EXPIRY_KEY = "jobradar_token_exp";

const TOKEN_LIFETIME_MS = 15 * 60 * 1000;

function isBrowser() {
  return typeof window !== "undefined";
}

export function saveSession(token, user) {
  if (!isBrowser()) return;
  const expiresAt = Date.now() + TOKEN_LIFETIME_MS;
  window.localStorage.setItem(TOKEN_KEY, token);
  window.localStorage.setItem(USER_KEY, JSON.stringify(user));
  window.localStorage.setItem(EXPIRY_KEY, String(expiresAt));
  window.dispatchEvent(new Event("jobradar-auth-change"));
}

export function clearSession() {
  if (!isBrowser()) return;
  window.localStorage.removeItem(TOKEN_KEY);
  window.localStorage.removeItem(USER_KEY);
  window.localStorage.removeItem(EXPIRY_KEY);
  window.dispatchEvent(new Event("jobradar-auth-change"));
}

export function getToken() {
  if (!isBrowser()) return null;
  const expiry = window.localStorage.getItem(EXPIRY_KEY);
  if (!expiry || Date.now() > Number(expiry)) {
    return null;
  }
  return window.localStorage.getItem(TOKEN_KEY);
}

export function getUser() {
  if (!isBrowser()) return null;
  if (!getToken()) return null;
  const raw = window.localStorage.getItem(USER_KEY);
  if (!raw) return null;
  try {
    return JSON.parse(raw);
  } catch {
    return null;
  }
}

export function getTokenExpiry() {
  if (!isBrowser()) return null;
  const expiry = window.localStorage.getItem(EXPIRY_KEY);
  return expiry ? Number(expiry) : null;
}

export function isAuthenticated() {
  return Boolean(getToken());
}
