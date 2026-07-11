"use client";

import { useCallback, useSyncExternalStore } from "react";
import { getUser, getTokenExpiry, clearSession, isAuthenticated } from "./auth";

const SERVER_SNAPSHOT = { user: null, expiresAt: null, ready: false };

let cachedSnapshot = null;

function readSession() {
  const next = isAuthenticated()
    ? { user: getUser(), expiresAt: getTokenExpiry(), ready: true }
    : { user: null, expiresAt: null, ready: true };

  if (
    cachedSnapshot &&
    cachedSnapshot.ready === next.ready &&
    cachedSnapshot.expiresAt === next.expiresAt &&
    JSON.stringify(cachedSnapshot.user) === JSON.stringify(next.user)
  ) {
    // Isi tidak berubah — kembalikan referensi lama supaya Object.is sama.
    return cachedSnapshot;
  }

  cachedSnapshot = next;
  return cachedSnapshot;
}

function subscribeSession(callback) {
  window.addEventListener("jobradar-auth-change", callback);
  window.addEventListener("storage", callback);
  // Polling ringan agar UI tahu saat token 15 menit kedaluwarsa secara
  // alami (tanpa aksi pengguna yang memicu 401 lebih dulu).
  const interval = setInterval(callback, 15000);
  return () => {
    window.removeEventListener("jobradar-auth-change", callback);
    window.removeEventListener("storage", callback);
    clearInterval(interval);
  };
}

export function useSession() {
  const session = useSyncExternalStore(
    subscribeSession,
    readSession,
    () => SERVER_SNAPSHOT
  );

  const logout = useCallback(() => {
    clearSession();
  }, []);

  return {
    user: session.user,
    expiresAt: session.expiresAt,
    ready: session.ready,
    isAuthenticated: Boolean(session.user),
    logout,
  };
}
