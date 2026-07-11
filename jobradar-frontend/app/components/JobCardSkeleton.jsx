export default function JobCardSkeleton() {
  return (
    <div className="bg-surface border border-line rounded-xl px-5 py-4 animate-pulse">
      <div className="flex items-start gap-4">
        <span className="mt-2 w-1.5 h-1.5 rounded-full shrink-0 bg-line" />
        <div className="flex-1 space-y-2.5">
          <div className="h-4 bg-line rounded w-2/5" />
          <div className="h-3 bg-line rounded w-1/4" />
          <div className="h-3 bg-line rounded w-3/5 mt-3" />
        </div>
      </div>
    </div>
  );
}
