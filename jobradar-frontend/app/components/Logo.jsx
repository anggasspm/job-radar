export default function Logo({ className = "" }) {
  return (
    <div className={`flex items-center gap-2 ${className}`}>
      <svg width="22" height="22" viewBox="0 0 24 24" fill="none" aria-hidden="true">
        <circle cx="12" cy="12" r="3" fill="#FF6B4A" />
        <circle
          cx="12"
          cy="12"
          r="7"
          stroke="#14213D"
          strokeWidth="1.5"
          strokeOpacity="0.3"
        />
        <circle
          cx="12"
          cy="12"
          r="11"
          stroke="#14213D"
          strokeWidth="1.5"
          strokeOpacity="0.15"
        />
      </svg>
      <span className="font-display font-bold text-lg text-ink">JobRadar</span>
    </div>
  );
}
