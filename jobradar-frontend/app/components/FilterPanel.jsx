"use client";

import { SlidersHorizontal, X } from "lucide-react";

export default function FilterPanel({
  categories,
  locations,
  filters,
  onChange,
  onReset,
  activeCount,
}) {
  return (
    <div className="bg-surface border border-line rounded-xl p-5">
      <div className="flex items-center justify-between mb-4">
        <h2 className="flex items-center gap-2 font-display font-semibold text-sm text-ink">
          <SlidersHorizontal className="w-4 h-4" />
          Filter
        </h2>
        {activeCount > 0 && (
          <button
            onClick={onReset}
            className="flex items-center gap-1 text-xs text-ink-soft hover:text-coral-dark transition"
          >
            <X className="w-3 h-3" />
            Reset ({activeCount})
          </button>
        )}
      </div>

      <div className="space-y-5">
        <div>
          <label className="block text-xs font-medium text-ink-soft mb-1.5">
            Kategori
          </label>
          <select
            value={filters.category}
            onChange={(e) => onChange({ ...filters, category: e.target.value })}
            className="w-full text-sm border border-line rounded-lg px-3 py-2 bg-white focus:outline-none focus:ring-2 focus:ring-coral/30 focus:border-coral"
          >
            <option value="">Semua kategori</option>
            {categories.map((c) => (
              <option key={c} value={c}>
                {c}
              </option>
            ))}
          </select>
        </div>

        <div>
          <label className="block text-xs font-medium text-ink-soft mb-1.5">
            Lokasi
          </label>
          <select
            value={filters.location}
            onChange={(e) => onChange({ ...filters, location: e.target.value })}
            className="w-full text-sm border border-line rounded-lg px-3 py-2 bg-white focus:outline-none focus:ring-2 focus:ring-coral/30 focus:border-coral"
          >
            <option value="">Semua lokasi</option>
            {locations.map((l) => (
              <option key={l} value={l}>
                {l}
              </option>
            ))}
          </select>
        </div>

        <div>
          <label className="block text-xs font-medium text-ink-soft mb-1.5">
            Gaji minimum
          </label>
          <select
            value={filters.minSalary}
            onChange={(e) =>
              onChange({ ...filters, minSalary: e.target.value })
            }
            className="w-full text-sm border border-line rounded-lg px-3 py-2 bg-white focus:outline-none focus:ring-2 focus:ring-coral/30 focus:border-coral"
          >
            <option value="0">Berapa saja</option>
            <option value="5000000">Rp 5 jt+</option>
            <option value="8000000">Rp 8 jt+</option>
            <option value="10000000">Rp 10 jt+</option>
            <option value="15000000">Rp 15 jt+</option>
          </select>
        </div>

        <fieldset>
          <legend className="block text-xs font-medium text-ink-soft mb-1.5">
            Pengalaman
          </legend>
          <div className="flex flex-wrap gap-2">
            {[
              { value: "", label: "Semua" },
              { value: "0-1", label: "Fresh grad" },
              { value: "1-3", label: "1–3 thn" },
              { value: "3-5", label: "3–5 thn" },
              { value: "5-99", label: "5+ thn" },
            ].map((opt) => (
              <button
                key={opt.value}
                type="button"
                onClick={() => onChange({ ...filters, experience: opt.value })}
                className={`text-xs px-3 py-1.5 rounded-full border transition ${
                  filters.experience === opt.value
                    ? "bg-ink text-white border-ink"
                    : "border-line text-ink-soft hover:border-ink-soft/50"
                }`}
              >
                {opt.label}
              </button>
            ))}
          </div>
        </fieldset>
      </div>
    </div>
  );
}
