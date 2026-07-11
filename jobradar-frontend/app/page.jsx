"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { ArrowRight, Search } from "lucide-react";

const EXAMPLE_QUERIES = [
  "kerja remote backend gaji minimal 8 juta",
  "data analyst fresh graduate di Bandung",
  "frontend developer Jakarta, hybrid",
  "magang UI/UX dekat Surabaya",
];

const SOURCES = ["Glints", "Tech in Asia", "We Work Remotely"];

export default function LandingPage() {
  const router = useRouter();
  const [query, setQuery] = useState("");

  function handleSubmit(e) {
    e.preventDefault();
    const params = query.trim()
      ? `?q=${encodeURIComponent(query.trim())}`
      : "";
    router.push(`/jobs${params}`);
  }

  return (
    <main>
      <section className="max-w-4xl mx-auto px-6 pt-20 pb-16 text-center">
        <span className="inline-flex items-center gap-1.5 text-xs font-mono text-coral-dark bg-coral/10 px-3 py-1 rounded-full mb-6">
          <span className="w-1.5 h-1.5 rounded-full bg-coral signal-dot" />
          memindai {SOURCES.length} sumber lowongan
        </span>

        <h1 className="font-display font-bold text-4xl md:text-6xl tracking-tight leading-[1.05] text-ink">
          Berhenti scroll 10 platform.
          <br />
          Cukup bilang apa yang kamu cari.
        </h1>

        <p className="mt-5 text-lg text-ink-soft max-w-xl mx-auto">
          JobRadar memindai lowongan dari {SOURCES.join(", ")}, lalu
          memahami kalimat biasa — bukan cuma kata kunci.
        </p>

        <form
          onSubmit={handleSubmit}
          className="mt-10 relative max-w-xl mx-auto"
        >
          <div className="flex items-center bg-white border border-line rounded-full pl-5 pr-2 py-2 shadow-sm focus-within:border-coral focus-within:ring-2 focus-within:ring-coral/20 transition">
            <Search className="w-4 h-4 text-ink-soft/50 shrink-0 mr-2" />
            <input
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              placeholder={EXAMPLE_QUERIES[0]}
              className="flex-1 min-w-0 text-left text-sm bg-transparent focus:outline-none placeholder:text-ink-soft/50"
            />
            <button
              type="submit"
              className="flex items-center gap-1.5 bg-ink text-white text-sm font-medium px-5 py-2.5 rounded-full hover:bg-ink/90 transition shrink-0"
            >
              Cari <ArrowRight className="w-4 h-4" />
            </button>
          </div>
          <div className="mt-3 flex flex-wrap justify-center gap-x-4 gap-y-1 text-xs text-ink-soft/70 font-mono">
            {EXAMPLE_QUERIES.slice(1).map((q) => (
              <button
                key={q}
                type="button"
                onClick={() => setQuery(q)}
                className="hover:text-coral-dark transition"
              >
                “{q}”
              </button>
            ))}
          </div>
        </form>
      </section>

      <section className="max-w-4xl mx-auto px-6 py-16 border-t border-line">
        <h2 className="font-display font-bold text-2xl mb-10 text-center text-ink">
          Cara kerjanya
        </h2>
        <div className="grid md:grid-cols-3 gap-8">
          <Step
            title="Ceritakan kriterianya"
            description="Ketik dengan bahasamu sendiri — peran, kota, gaji, tipe kerja."
          />
          <Step
            title="Kami menyaring lowongan aktif"
            description="Setiap lowongan dicocokkan ke judul, lokasi, kategori, dan perusahaan sekaligus."
          />
          <Step
            title="Lamar langsung ke sumber asli"
            description="Setiap hasil terhubung langsung ke halaman lamaran resminya."
          />
        </div>
      </section>
    </main>
  );
}

function Step({ title, description }) {
  return (
    <div className="border-l-2 border-line pl-4">
      <h3 className="font-display font-semibold text-ink">{title}</h3>
      <p className="mt-2 text-sm text-ink-soft">{description}</p>
    </div>
  );
}
