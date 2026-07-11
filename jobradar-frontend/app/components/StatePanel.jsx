import { AlertTriangle, RadarIcon, SearchX } from "lucide-react";

const ICONS = {
  error: AlertTriangle,
  empty: SearchX,
  radar: RadarIcon,
};

export default function StatePanel({
  variant = "empty",
  title,
  description,
  action,
}) {
  const Icon = ICONS[variant] || SearchX;
  const iconColor = variant === "error" ? "text-coral-dark" : "text-ink-soft/50";

  return (
    <div className="flex flex-col items-center text-center py-16 px-6">
      <Icon className={`w-8 h-8 mb-4 ${iconColor}`} strokeWidth={1.5} />
      <p className="font-display font-semibold text-ink">{title}</p>
      {description && (
        <p className="text-sm text-ink-soft mt-1.5 max-w-sm">{description}</p>
      )}
      {action && <div className="mt-5">{action}</div>}
    </div>
  );
}
