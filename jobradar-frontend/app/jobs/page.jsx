"use client";

import { useState, useEffect, useMemo, useCallback, Suspense } from "react";
import { useSearchParams } from "next/navigation";
import { Search, Loader2 } from "lucide-react";
import JobCard from "../components/JobCard";
import JobCardSkeleton from "../components/JobCardSkeleton";
import StatePanel from "../components/StatePanel";
import FilterPanel from "../components/FilterPanel";
import { fetchJobs, searchJobs, ApiError } from "../lib/api";

const EMPTY_FILTERS = { category: "", location: "", minSalary: "0", experience: "" };

function JobsPageInner() {
  const searchParams = useSearchParams();
  const initialQuery = searchParams.get("q") || "";
  const [inputValue, setInputValue] = useState(initialQuery);
  const [query, setQuery] = useState(initialQuery);
  const [jobs, setJobs] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [filters, setFilters] = useState(EMPTY_FILTERS);
  const [showFilters, setShowFilters] = useState(false);

  const runSearch = useCallback(async (q) => {
    setLoading(true);
    setError(null);
    try {
      const data = q ? await searchJobs(q) : await fetchJobs();
      setJobs(Array.isArray(data) ? data : []);
    } catch (err) {
      setError(
        err instanceof ApiError
          ? err.message
          : "Terjadi kesalahan yang tidak terduga."
      );
      setJobs([]);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    let cancelled = false;
    async function load() {
      try {
        const data = initialQuery
          ? await searchJobs(initialQuery)
          : await fetchJobs();
        if (!cancelled) {
          setJobs(Array.isArray(data) ? data : []);
          setError(null);
        }
      } catch (err) {
        if (!cancelled) {
          setError(
            err instanceof ApiError
              ? err.message
              : "Terjadi kesalahan yang tidak terduga."
          );
          setJobs([]);
        }
      } finally {
        if (!cancelled) setLoading(false);
      }
    }
    load();
    return () => {
      cancelled = true;
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  function handleSubmit(e) {
    e.preventDefault();
    setQuery(inputValue);
    runSearch(inputValue);
  }

  const categories = useMemo(
    () => [...new Set(jobs.map((j) => j.category).filter(Boolean))].sort(),
    [jobs]
  );
  const locations = useMemo(
    () => [...new Set(jobs.map((j) => j.location).filter(Boolean))].sort(),
    [jobs]
  );

  const filteredJobs = useMemo(() => {
    return jobs.filter((job) => {
      if (filters.category && job.category !== filters.category) return false;
      if (filters.location && job.location !== filters.location) return false;
      if (Number(filters.minSalary) > 0) {
        const jobMax = job.salaryMax || job.salaryMin || 0;
        if (jobMax < Number(filters.minSalary)) return false;
      }
      if (filters.experience) {
        const [lo, hi] = filters.experience.split("-").map(Number);
        const jobMin = job.minExp ?? 0;
        if (jobMin < lo || jobMin > hi) return false;
      }
      return true;
    });
  }, [jobs, filters]);

  const activeFilterCount = Object.entries(filters).filter(
    ([k, v]) => v && v !== EMPTY_FILTERS[k]
  ).length;

  return (
    <main className="max-w-6xl mx-auto px-6 py-10">
      <div className="mb-6">
        <h1 className="font-display font-bold text-2xl md:text-3xl text-ink">
          Cari lowongan
        </h1>
        <p className="text-sm text-ink-soft mt-1">
          Ketik peran, kota, atau kata kunci apa pun — kami cocokkan ke judul,
          perusahaan, lokasi, dan kategori sekaligus.
        </p>
      </div>

      <form onSubmit={handleSubmit} className="flex gap-2 mb-6">
        <div className="relative flex-1">
          <Search className="absolute left-4 top-1/2 -translate-y-1/2 w-4 h-4 text-ink-soft/60" />
          <input
            type="text"
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
            placeholder="mis. backend Jakarta gaji 8 juta"
            className="w-full pl-11 pr-4 py-3.5 rounded-full border border-line bg-white text-sm focus:outline-none focus:ring-2 focus:ring-coral/30 focus:border-coral"
          />
        </div>
        <button
          type="submit"
          disabled={loading}
          className="bg-ink text-white px-6 py-3.5 rounded-full font-medium text-sm hover:bg-ink/90 transition disabled:opacity-50 shrink-0 flex items-center gap-2"
        >
          {loading && <Loader2 className="w-4 h-4 animate-spin" />}
          Cari
        </button>
        <button
          type="button"
          onClick={() => setShowFilters((s) => !s)}
          className="md:hidden border border-line px-4 py-3.5 rounded-full text-sm font-medium text-ink-soft shrink-0"
        >
          Filter{activeFilterCount > 0 ? ` (${activeFilterCount})` : ""}
        </button>
      </form>

      {!loading && !error && (
        <div className="mb-6">
          <div className="flex items-center justify-between text-xs font-mono text-ink-soft mb-2">
            <span>
              {filteredJobs.length} sinyal ditemukan
              {query && <> untuk “{query}”</>}
              {jobs.length !== filteredJobs.length && (
                <> · {jobs.length} total sebelum filter</>
              )}
            </span>
          </div>
          <div className="scan-line" />
        </div>
      )}

      <div className="grid md:grid-cols-[240px_1fr] gap-6">
        <aside className={`${showFilters ? "block" : "hidden"} md:block`}>
          <div className="md:sticky md:top-24">
            <FilterPanel
              categories={categories}
              locations={locations}
              filters={filters}
              onChange={setFilters}
              onReset={() => setFilters(EMPTY_FILTERS)}
              activeCount={activeFilterCount}
            />
          </div>
        </aside>

        <div className="space-y-3 min-w-0">
          {loading &&
            Array.from({ length: 6 }).map((_, i) => (
              <JobCardSkeleton key={i} />
            ))}

          {!loading && error && (
            <StatePanel
              variant="error"
              title="Gagal mengambil data lowongan"
              description={error}
              action={
                <button
                  onClick={() => runSearch(query)}
                  className="text-sm font-medium bg-ink text-white px-5 py-2.5 rounded-full hover:bg-ink/90 transition"
                >
                  Coba lagi
                </button>
              }
            />
          )}

          {!loading && !error && filteredJobs.length === 0 && (
            <StatePanel
              variant="empty"
              title="Tidak ada lowongan yang cocok"
              description="Coba longgarkan filter atau ganti kata kunci pencarian."
              action={
                activeFilterCount > 0 ? (
                  <button
                    onClick={() => setFilters(EMPTY_FILTERS)}
                    className="text-sm font-medium text-coral-dark hover:underline"
                  >
                    Hapus semua filter
                  </button>
                ) : null
              }
            />
          )}

          {!loading &&
            !error &&
            filteredJobs.map((job) => <JobCard key={job.id} job={job} />)}
        </div>
      </div>
    </main>
  );
}

export default function JobsPage() {
  return (
    <Suspense fallback={null}>
      <JobsPageInner />
    </Suspense>
  );
}
