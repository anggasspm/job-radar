import Link from "next/link";
import { ArrowLeft, MapPin, Briefcase, GraduationCap, Clock } from "lucide-react";
import StatePanel from "../../components/StatePanel";

async function getJob(id) {
  const base = process.env.NEXT_PUBLIC_API_BASE_URL;
  if (!base) return { job: null, failed: true };

  try {
    const res = await fetch(`${base}/jobs/${id}`, { cache: "no-store" });
    if (res.status === 404) return { job: null, failed: false };
    if (!res.ok) return { job: null, failed: true };
    const job = await res.json();
    return { job, failed: false };
  } catch {
    return { job: null, failed: true };
  }
}

function formatSalary(min, max, currency) {
  if (!min && !max) return null;
  const symbol = currency === "USD" ? "$" : "Rp";
  if (min && max && min !== max) {
    return `${symbol} ${min.toLocaleString("id-ID")} – ${symbol} ${max.toLocaleString("id-ID")}`;
  }
  return `${symbol} ${(min || max).toLocaleString("id-ID")}`;
}

export default async function JobDetailPage({ params }) {
  const { id } = await params;
  const { job, failed } = await getJob(id);

  if (failed) {
    return (
      <main className="max-w-2xl mx-auto px-6 py-16">
        <StatePanel
          variant="error"
          title="Tidak bisa memuat lowongan ini"
          description="Server sedang tidak bisa dihubungi. Coba muat ulang halaman dalam beberapa saat."
          action={
            <Link
              href="/jobs"
              className="text-sm font-medium bg-ink text-white px-5 py-2.5 rounded-full hover:bg-ink/90 transition inline-block"
            >
              Kembali cari lowongan
            </Link>
          }
        />
      </main>
    );
  }

  if (!job) {
    return (
      <main className="max-w-2xl mx-auto px-6 py-16">
        <StatePanel
          variant="empty"
          title="Lowongan tidak ditemukan"
          description="Lowongan ini mungkin sudah ditutup atau tautannya salah."
          action={
            <Link
              href="/jobs"
              className="text-sm font-medium bg-ink text-white px-5 py-2.5 rounded-full hover:bg-ink/90 transition inline-block"
            >
              Kembali cari lowongan
            </Link>
          }
        />
      </main>
    );
  }

  const salary = formatSalary(job.salaryMin, job.salaryMax, job.currency);

  return (
    <main className="max-w-2xl mx-auto px-6 py-12">
      <Link
        href="/jobs"
        className="inline-flex items-center gap-1.5 text-sm text-ink-soft hover:text-ink mb-6 transition"
      >
        <ArrowLeft className="w-4 h-4" /> Kembali
      </Link>

      <div className="bg-surface border border-line rounded-2xl p-8">
        {!job.isActive && (
          <div className="mb-4 text-xs font-mono bg-amber/10 text-amber inline-block px-2.5 py-1 rounded-md">
            Lowongan ini sudah tidak aktif
          </div>
        )}

        <h1 className="font-display font-bold text-2xl text-ink">
          {job.title}
        </h1>
        <p className="text-ink-soft mt-1">{job.company}</p>

        <div className="mt-4 flex items-center gap-4 text-sm text-ink-soft flex-wrap">
          <span className="flex items-center gap-1.5">
            <MapPin className="w-4 h-4" /> {job.location}
          </span>
          {job.category && (
            <span className="flex items-center gap-1.5">
              <Briefcase className="w-4 h-4" /> {job.category}
            </span>
          )}
          {job.education && (
            <span className="flex items-center gap-1.5">
              <GraduationCap className="w-4 h-4" /> {job.education}
            </span>
          )}
          {(job.minExp || job.maxExp) && (
            <span className="flex items-center gap-1.5">
              <Clock className="w-4 h-4" /> {job.minExp}–{job.maxExp} tahun
            </span>
          )}
        </div>

        {salary && (
          <p className="mt-4 font-mono text-teal-dark text-sm bg-teal/10 inline-block px-3 py-1.5 rounded-full">
            {salary}
          </p>
        )}

        <div className="mt-6 text-ink/80 leading-relaxed whitespace-pre-line">
          {job.description || "Deskripsi lowongan belum tersedia untuk sumber ini."}
        </div>

        <a
          href={job.rawUrl}
          target="_blank"
          rel="noopener noreferrer"
          className="mt-8 inline-block bg-coral text-white px-6 py-3 rounded-full font-medium hover:bg-coral-dark transition"
        >
          Lamar lowongan ini
        </a>
      </div>
    </main>
  );
}
