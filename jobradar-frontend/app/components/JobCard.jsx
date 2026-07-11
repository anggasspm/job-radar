import Link from "next/link";
import { MapPin, ArrowUpRight } from "lucide-react";

// Setiap sumber punya warna sinyal sendiri — bukan tag generik,
// tapi identitas kecil yang konsisten di seluruh daftar.
function sourceColor(sourceId) {
  const bySourceId = { 1: "#FF6B4A", 2: "#2EC4B6", 3: "#14213D" };
  return bySourceId[sourceId] || "#8A93A6";
}

function formatSalary(min, max, currency) {
  if (!min && !max) return null;
  const symbol = currency === "USD" ? "$" : "Rp";
  const fmt = (n) =>
    currency === "USD"
      ? Math.round(n / 1000)
      : Math.round(n / 1_000_000);
  const unit = currency === "USD" ? "rb" : "jt";
  if (min && max && min !== max) {
    return `${symbol} ${fmt(min)}–${fmt(max)} ${unit}`;
  }
  return `${symbol} ${fmt(min || max)} ${unit}`;
}

export default function JobCard({ job }) {
  const salary = formatSalary(job.salaryMin, job.salaryMax, job.currency);
  const dotColor = sourceColor(job.sourceId);

  return (
    <Link
      href={`/jobs/${job.id}`}
      className="group block bg-surface border border-line rounded-xl px-5 py-4 hover:border-ink-soft/40 hover:shadow-sm transition-all"
    >
      <div className="flex items-start gap-4">
        <span
          className="mt-2 w-1.5 h-1.5 rounded-full shrink-0 signal-dot"
          style={{ backgroundColor: dotColor }}
          aria-hidden="true"
        />
        <div className="flex-1 min-w-0">
          <div className="flex items-start justify-between gap-3">
            <div className="min-w-0">
              <h3 className="font-display font-semibold text-ink truncate">
                {job.title}
              </h3>
              <p className="text-sm text-ink-soft mt-0.5 truncate">
                {job.company}
              </p>
            </div>
            <ArrowUpRight className="w-4 h-4 text-ink-soft/40 group-hover:text-coral group-hover:translate-x-0.5 group-hover:-translate-y-0.5 transition shrink-0 mt-1" />
          </div>

          <div className="mt-3 flex items-center gap-x-4 gap-y-1.5 flex-wrap font-mono text-xs">
            <span className="flex items-center gap-1 text-ink-soft">
              <MapPin className="w-3.5 h-3.5" />
              {job.location}
            </span>
            {job.category && (
              <span className="text-ink-soft">{job.category}</span>
            )}
            {(job.minExp || job.maxExp) ? (
              <span className="text-ink-soft">
                {job.minExp}–{job.maxExp} thn
              </span>
            ) : null}
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
}
