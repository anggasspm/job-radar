"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import Logo from "../components/Logo";

const API_BASE = process.env.NEXT_PUBLIC_API_BASE_URL;

export default function RegisterPage() {
  const router = useRouter();
  const [form, setForm] = useState({ email: "", password: "" });
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e) {
    e.preventDefault();
    setLoading(true);
    setError(null);
    try {
      const res = await fetch(`${API_BASE}/auth/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(form),
      });
      if (!res.ok) throw new Error("Pendaftaran gagal, coba email lain");
      router.push("/login");
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }

  return (
    <main className="min-h-[calc(100vh-4rem)] flex items-center justify-center px-6">
      <form onSubmit={handleSubmit} className="w-full max-w-sm">
        <div className="flex justify-center mb-8"><Logo /></div>
        <div className="bg-white border border-line rounded-2xl p-8">
          <h1 className="font-display font-bold text-xl mb-6">Buat akun</h1>
          {error && <p className="text-red-600 text-sm mb-4">{error}</p>}

          <label className="block text-sm font-medium text-ink/70 mb-1">Email</label>
          <input
            type="email" required value={form.email}
            onChange={(e) => setForm({ ...form, email: e.target.value })}
            className="w-full mb-4 px-4 py-2.5 rounded-xl border border-line focus:outline-none focus:ring-2 focus:ring-coral/40 focus:border-coral"
          />

          <label className="block text-sm font-medium text-ink/70 mb-1">Password</label>
          <input
            type="password" required value={form.password}
            onChange={(e) => setForm({ ...form, password: e.target.value })}
            className="w-full mb-6 px-4 py-2.5 rounded-xl border border-line focus:outline-none focus:ring-2 focus:ring-coral/40 focus:border-coral"
          />

          <button
            type="submit" disabled={loading}
            className="w-full bg-ink text-white py-2.5 rounded-xl font-medium hover:bg-ink/90 transition disabled:opacity-50"
          >
            {loading ? "Memproses…" : "Daftar"}
          </button>
        </div>
        <p className="mt-5 text-sm text-ink/60 text-center">
          Sudah punya akun? <Link href="/login" className="text-coral-dark font-medium">Masuk</Link>
        </p>
      </form>
    </main>
  );
}