"use client";

import { useState } from "react";
import Link from "next/link";
import { Search, MapPin, Briefcase, ArrowRight } from "lucide-react";

const API_BASE = process.env.NEXT_PUBLIC_API_BASE_URL;

export default function JobsPage() {
  const [query, setQuery] = useState("");
  const [jobs, setJobs] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [searched, setSearched] = useState(false);

  async function handleSearch(e) {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSearched(true);
    try {
      const res = await fetch(`${API_BASE}/jobs/search?q=${encodeURIComponent(query)}`);
      if (!res.ok) throw new Error("Gagal mengambil data lowongan");
      const data = await res.json();
      setJobs(data || []); // backend balikin array langsung, bukan { results: [...] }
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }

  return (
    <main className="max-w-3xl mx-auto px-6 py-12">
      <h1 className="font-display font-bold text-3xl mb-6">Cari lowongan</h1>

      <form onSubmit={handleSearch} className="flex gap-2 mb-10">
        <div className="relative flex-1">
          <Search className="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-ink/40" />
          <input
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            placeholder="kerja remote backend gaji minimal 8 juta di Jakarta"
            className="w-full pl-11 pr-4 py-3.5 rounded-full border border-line bg-white font-mono text-sm focus:outline-none focus:ring-2 focus:ring-coral/40 focus:border-coral"
          />
        </div>
        <button
          type="submit"
          disabled={loading}
          className="bg-ink text-white px-6 py-3.5 rounded-full font-medium hover:bg-ink/90 transition disabled:opacity-50 shrink-0"
        >
          {loading ? "Mencari…" : "Cari"}
        </button>
      </form>

      {error && <p className="text-red-600 mb-6">{error}</p>}

      <div className="space-y-3">
        {!searched && (
          <p className="text-ink/40 text-center py-16">Ketik kriteria pekerjaanmu di atas untuk mulai.</p>
        )}
        {searched && !loading && jobs.length === 0 && !error && (
          <p className="text-ink/40 text-center py-16">Tidak ada lowongan yang cocok. Coba kriteria lain.</p>
        )}
        {jobs.map((job) => (
          <Link
            key={job.id}
            href={`/jobs/${job.id}`}
            className="group block bg-white border border-line rounded-2xl p-5 hover:border-coral/50 hover:shadow-md transition"
          >
            <div className="flex items-start justify-between gap-4">
              <div>
                <h3 className="font-display font-semibold">{job.title}</h3>
                <p className="text-sm text-ink/60 mt-0.5">{job.company}</p>
              </div>
              <ArrowRight className="w-4 h-4 text-ink/30 group-hover:text-coral group-hover:translate-x-0.5 transition shrink-0 mt-1" />
            </div>
            <div className="mt-3 flex items-center gap-3 flex-wrap">
              <Tag icon={<MapPin className="w-3.5 h-3.5" />}>{job.location}</Tag>
              <Tag icon={<Briefcase className="w-3.5 h-3.5" />}>{job.category}</Tag>
              {job.salaryMin > 0 && (
                <span className="font-mono text-xs text-teal bg-teal/10 px-2.5 py-1 rounded-full">
                  Rp {Math.round(job.salaryMin / 1_000_000)}–{Math.round(job.salaryMax / 1_000_000)} jt
                </span>
              )}
            </div>
          </Link>
        ))}
      </div>
    </main>
  );
}

function Tag({ icon, children }) {
  return (
    <span className="flex items-center gap-1 text-xs text-ink/60">
      {icon} {children}
    </span>
  );
}