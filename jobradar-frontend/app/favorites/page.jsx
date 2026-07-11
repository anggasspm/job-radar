"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import { MapPin, ArrowUpRight, Bookmark, Loader2 } from "lucide-react";
import StatePanel from "../components/StatePanel";
import { useSession } from "../lib/useSession";
import { fetchFavorites, ApiError } from "../lib/api";

function formatSalary(min, max) {
  if (!min && !max) return null;
  const fmt = (n) => Math.round(n / 1_000_000);
  if (min && max && min !== max) return `Rp ${fmt(min)}–${fmt(max)} jt`;
  return `Rp ${fmt(min || max)} jt`;
}

export default function FavoritesPage() {
  const { ready, isAuthenticated } = useSession();
  const [favorites, setFavorites] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!ready) return;
    let cancelled = false;

    async function load() {
      if (!isAuthenticated) {
        if (!cancelled) setLoading(false);
        return;
      }
      if (!cancelled) setLoading(true);
      try {
        const data = await fetchFavorites();
        if (!cancelled) setFavorites(Array.isArray(data) ? data : []);
      } catch (err) {
        if (!cancelled) {
          setError(
            err instanceof ApiError
              ? err.message
              : "Gagal memuat daftar favorit."
          );
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    }

    load();
    return () => {
      cancelled = true;
    };
  }, [ready, isAuthenticated]);

  if (!ready || loading) {
    return (
      <main className="max-w-3xl mx-auto px-6 py-16 flex justify-center">
        <Loader2 className="w-5 h-5 animate-spin text-ink-soft" />
      </main>
    );
  }

  if (!isAuthenticated) {
    return (
      <main className="max-w-2xl mx-auto px-6 py-16">
        <StatePanel
          variant="empty"
          title="Masuk untuk melihat favoritmu"
          description="Lowongan yang kamu simpan akan muncul di sini setelah kamu masuk ke akunmu."
          action={
            <Link
              href="/login"
              className="text-sm font-medium bg-ink text-white px-5 py-2.5 rounded-full hover:bg-ink/90 transition inline-block"
            >
              Masuk
            </Link>
          }
        />
      </main>
    );
  }

  return (
    <main className="max-w-3xl mx-auto px-6 py-12">
      <h1 className="font-display font-bold text-2xl text-ink mb-1">
        Lowongan favorit
      </h1>
      <p className="text-sm text-ink-soft mb-8">
        Fitur menyimpan lowongan baru sedang dalam pengembangan di sisi
        server — untuk saat ini halaman ini hanya menampilkan yang sudah
        tersimpan.
      </p>

      {error && (
        <StatePanel variant="error" title="Gagal memuat" description={error} />
      )}

      {!error && favorites.length === 0 && (
        <StatePanel
          variant="empty"
          title="Belum ada lowongan tersimpan"
          description="Saat fitur simpan aktif, lowongan yang kamu tandai akan muncul di sini."
          action={
            <Link
              href="/jobs"
              className="text-sm font-medium bg-ink text-white px-5 py-2.5 rounded-full hover:bg-ink/90 transition inline-block"
            >
              Jelajahi lowongan
            </Link>
          }
        />
      )}

      <div className="space-y-3">
        {favorites.map((fav) => {
          const salary = formatSalary(fav.salary_min, fav.salary_max);
          return (
            <Link
              key={fav.id}
              href={`/jobs/${fav.job_id}`}
              className="group block bg-surface border border-line rounded-xl px-5 py-4 hover:border-ink-soft/40 hover:shadow-sm transition"
            >
              <div className="flex items-start gap-3">
                <Bookmark className="w-4 h-4 text-coral mt-1 shrink-0" />
                <div className="flex-1 min-w-0">
                  <div className="flex items-start justify-between gap-3">
                    <div className="min-w-0">
                      <h3 className="font-display font-semibold text-ink truncate">
                        {fav.title}
                      </h3>
                      <p className="text-sm text-ink-soft mt-0.5 truncate">
                        {fav.company}
                      </p>
                    </div>
                    <ArrowUpRight className="w-4 h-4 text-ink-soft/40 group-hover:text-coral transition shrink-0 mt-1" />
                  </div>
                  <div className="mt-2.5 flex items-center gap-3 flex-wrap font-mono text-xs">
                    <span className="flex items-center gap-1 text-ink-soft">
                      <MapPin className="w-3.5 h-3.5" />
                      {fav.location}
                    </span>
                    {salary && (
                      <span className="text-teal-dark bg-teal/10 px-2 py-0.5 rounded-md">
                        {salary}
                      </span>
                    )}
                  </div>
                </div>
              </div>
            </Link>
          );
        })}
      </div>
    </main>
  );
}
