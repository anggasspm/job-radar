"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { Loader2 } from "lucide-react";
import Logo from "../components/Logo";
import { login, ApiError } from "../lib/api";
import { saveSession } from "../lib/auth";

export default function LoginPage() {
  const router = useRouter();
  const [form, setForm] = useState({ email: "", password: "" });
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e) {
    e.preventDefault();
    setLoading(true);
    setError(null);
    try {
      const res = await login(form.email, form.password);
      saveSession(res.data.access_token, res.data.user);
      router.push("/jobs");
    } catch (err) {
      setError(
        err instanceof ApiError
          ? err.message
          : "Gagal masuk. Coba lagi sebentar lagi."
      );
    } finally {
      setLoading(false);
    }
  }

  return (
    <main className="min-h-[calc(100vh-4rem)] flex items-center justify-center px-6">
      <form onSubmit={handleSubmit} className="w-full max-w-sm">
        <div className="flex justify-center mb-8">
          <Logo />
        </div>
        <div className="bg-surface border border-line rounded-2xl p-8">
          <h1 className="font-display font-bold text-xl mb-6 text-ink">
            Masuk
          </h1>
          {error && (
            <p
              role="alert"
              className="text-coral-dark text-sm mb-4 bg-coral/10 rounded-lg px-3 py-2"
            >
              {error}
            </p>
          )}

          <label className="block text-sm font-medium text-ink-soft mb-1">
            Email
          </label>
          <input
            type="email"
            required
            autoComplete="email"
            value={form.email}
            onChange={(e) => setForm({ ...form, email: e.target.value })}
            className="w-full mb-4 px-4 py-2.5 rounded-xl border border-line focus:outline-none focus:ring-2 focus:ring-coral/30 focus:border-coral"
          />

          <label className="block text-sm font-medium text-ink-soft mb-1">
            Kata sandi
          </label>
          <input
            type="password"
            required
            autoComplete="current-password"
            value={form.password}
            onChange={(e) => setForm({ ...form, password: e.target.value })}
            className="w-full mb-6 px-4 py-2.5 rounded-xl border border-line focus:outline-none focus:ring-2 focus:ring-coral/30 focus:border-coral"
          />

          <button
            type="submit"
            disabled={loading}
            className="w-full bg-ink text-white py-2.5 rounded-xl font-medium hover:bg-ink/90 transition disabled:opacity-50 flex items-center justify-center gap-2"
          >
            {loading && <Loader2 className="w-4 h-4 animate-spin" />}
            {loading ? "Memproses…" : "Masuk"}
          </button>
        </div>
        <p className="mt-5 text-sm text-ink-soft text-center">
          Belum punya akun?{" "}
          <Link href="/register" className="text-coral-dark font-medium">
            Daftar
          </Link>
        </p>
      </form>
    </main>
  );
}
