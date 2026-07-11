// Klien API tunggal untuk backend JobRadar (Go + Gin).
// Semua bentuk respons di sini mengikuti persis DTO backend
// (lihat internal/dto/*.go) supaya tidak ada tebakan bentuk data.

import { getToken, clearSession } from "./auth";

const API_BASE = process.env.NEXT_PUBLIC_API_BASE_URL;

class ApiError extends Error {
  constructor(message, status) {
    super(message);
    this.status = status;
  }
}

async function request(path, options = {}) {
  if (!API_BASE) {
    throw new ApiError(
      "NEXT_PUBLIC_API_BASE_URL belum diatur. Backend tidak bisa dihubungi.",
      0
    );
  }

  const token = getToken();
  const headers = {
    "Content-Type": "application/json",
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
    ...options.headers,
  };

  let res;
  try {
    res = await fetch(`${API_BASE}${path}`, { ...options, headers });
  } catch {
    // Ini biasanya cold-start Hugging Face Space, DNS gagal, atau CORS.
    throw new ApiError(
      "Tidak bisa menghubungi server. Server mungkin sedang bangun dari tidur (cold start) — coba lagi dalam beberapa detik.",
      0
    );
  }

  if (res.status === 401) {
    clearSession();
    throw new ApiError("Sesi berakhir. Silakan masuk kembali.", 401);
  }

  let body = null;
  const text = await res.text();
  if (text) {
    try {
      body = JSON.parse(text);
    } catch {
      body = null;
    }
  }

  if (!res.ok) {
    const message =
      (body && (body.error || body.message)) ||
      `Permintaan gagal (${res.status})`;
    throw new ApiError(message, res.status);
  }

  return body;
}

// ---- Jobs ----
// Backend selalu balikin array polos untuk /jobs dan /jobs/search.
export function fetchJobs() {
  return request("/jobs");
}

export function searchJobs(query) {
  const q = query?.trim() ?? "";
  const params = q ? `?q=${encodeURIComponent(q)}` : "";
  return request(`/jobs/search${params}`);
}

export function fetchJobById(id) {
  return request(`/jobs/${id}`);
}

// ---- Auth ----
export function login(email, password) {
  return request("/auth/login", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
}

export function register(name, email, password) {
  return request("/auth/register", {
    method: "POST",
    body: JSON.stringify({ name, email, password }),
  });
}

// ---- Favorites ----
// Catatan jujur: backend saat ini hanya punya endpoint GET untuk favorit.
// Menambah/menghapus favorit belum tersedia di sisi server, jadi UI
// menampilkan itu secara terbuka alih-alih berpura-pura berfungsi.
export function fetchFavorites() {
  return request("/favorite/");
}

export { ApiError };
