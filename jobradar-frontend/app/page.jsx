"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import { ArrowRight, Radar } from "lucide-react";

const EXAMPLE_QUERIES = [
  "kerja remote backend gaji minimal 8 juta",
  "data analyst fresh graduate di Bandung",
  "frontend developer Jakarta, hybrid",
  "magang UI/UX dekat Surabaya",
];

const STATS = [
  { value: "12.400+", label: "lowongan dipantau" },
  { value: "3", label: "sumber terintegrasi" },
  { value: "6 jam", label: "siklus update" },
];

export default function LandingPage() {
  const [i, setI] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => setI((x) => (x + 1) % EXAMPLE_QUERIES.length), 2800);
    return () => clearInterval(interval);
  }, []);

  return (
    <main>
      <section className="max-w-4xl mx-auto px-6 pt-20 pb-16 text-center">
        <span className="inline-flex items-center gap-1.5 text-xs font-medium text-coral-dark bg-coral/10 px-3 py-1 rounded-full mb-6">
          <Radar className="w-3.5 h-3.5" /> Pencarian kerja bertenaga AI
        </span>

        <h1 className="font-display font-bold text-4xl md:text-6xl tracking-tight leading-[1.05]">
          Berhenti scroll 10 platform.
          <br />
          Cukup bilang apa yang kamu cari.
        </h1>

        <p className="mt-5 text-lg text-ink/60 max-w-xl mx-auto">
          JobRadar memindai ribuan lowongan setiap hari, lalu memahami kalimat
          biasa — bukan cuma kata kunci.
        </p>

        <div className="mt-10 relative max-w-xl mx-auto">
          <div className="absolute left-5 top-1/2 -translate-y-1/2 w-3 h-3">
            <span className="radar-ring" />
            <span className="radar-ring" />
            <span className="radar-ring" />
            <span className="absolute inset-0 rounded-full bg-coral" />
          </div>
          <div className="flex items-center bg-white border border-line rounded-full pl-12 pr-2 py-2 shadow-sm">
            <span className="flex-1 text-left text-ink/50 truncate font-mono text-sm">
              {EXAMPLE_QUERIES[i]}
            </span>
            <Link
              href="/jobs"
              className="flex items-center gap-1.5 bg-ink text-white text-sm font-medium px-5 py-2.5 rounded-full hover:bg-ink/90 transition shrink-0"
            >
              Cari <ArrowRight className="w-4 h-4" />
            </Link>
          </div>
        </div>

        <div className="mt-14 flex justify-center gap-10 flex-wrap">
          {STATS.map((s) => (
            <div key={s.label}>
              <p className="font-display font-bold text-2xl">{s.value}</p>
              <p className="text-sm text-ink/50">{s.label}</p>
            </div>
          ))}
        </div>
      </section>

      <section className="max-w-4xl mx-auto px-6 py-16 border-t border-line">
        <h2 className="font-display font-bold text-2xl mb-10 text-center">Cara kerjanya</h2>
        <div className="grid md:grid-cols-3 gap-8">
          <Step n="01" title="Ceritakan kriterianya" description="Ketik dengan bahasamu sendiri — role, kota, gaji, tipe kerja." />
          <Step n="02" title="AI menyaring ribuan lowongan" description="Sistem memahami maksudmu dan mencocokkan ke data yang selalu segar." />
          <Step n="03" title="Lamar atau pasang alert" description="Klik langsung ke sumber asli, atau biarkan kami memberi tahu duluan." />
        </div>
      </section>
    </main>
  );
}

function Step({ n, title, description }) {
  return (
    <div>
      <p className="font-mono text-sm text-coral-dark mb-3">{n}</p>
      <h3 className="font-display font-semibold">{title}</h3>
      <p className="mt-2 text-sm text-ink/60">{description}</p>
    </div>
  );
}