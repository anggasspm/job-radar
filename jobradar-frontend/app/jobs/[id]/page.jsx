import Link from "next/link";
import { ArrowLeft, MapPin, Briefcase } from "lucide-react";

async function getJob(id) {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/jobs/${id}`, { cache: "no-store" });
  if (!res.ok) return null;
  return res.json();
}

export default async function JobDetailPage({ params }) {
  const job = await getJob(params.id);

  if (!job) {
    return (
      <main className="max-w-2xl mx-auto px-6 py-20 text-center">
        <p className="text-ink/50">Lowongan tidak ditemukan.</p>
        <Link href="/jobs" className="text-coral-dark font-medium mt-2 inline-block">Kembali cari lowongan</Link>
      </main>
    );
  }

  return (
    <main className="max-w-2xl mx-auto px-6 py-12">
      <Link href="/jobs" className="inline-flex items-center gap-1.5 text-sm text-ink/50 hover:text-ink mb-6 transition">
        <ArrowLeft className="w-4 h-4" /> Kembali
      </Link>

      <div className="bg-white border border-line rounded-2xl p-8">
        <h1 className="font-display font-bold text-2xl">{job.title}</h1>
        <p className="text-ink/60 mt-1">{job.company}</p>

        <div className="mt-4 flex items-center gap-4 text-sm text-ink/60 flex-wrap">
          <span className="flex items-center gap-1"><MapPin className="w-4 h-4" /> {job.location}</span>
          <span className="flex items-center gap-1"><Briefcase className="w-4 h-4" /> {job.category}</span>
        </div>

        {job.salary_min && (
          <p className="mt-4 font-mono text-teal text-sm bg-teal/10 inline-block px-3 py-1.5 rounded-full">
            Rp {job.salary_min.toLocaleString("id-ID")} – Rp {job.salary_max?.toLocaleString("id-ID")}
          </p>
        )}

        <div className="mt-6 text-ink/80 leading-relaxed whitespace-pre-line">{job.description}</div>

        <a
          href={job.raw_url}
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