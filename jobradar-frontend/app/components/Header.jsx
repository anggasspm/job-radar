import Link from "next/link";
import Logo from "./Logo";

export default function Header() {
  return (
    <header className="sticky top-0 z-50 w-full bg-canvas/80 backdrop-blur-md border-b border-line">
      <div className="max-w-6xl mx-auto px-6 h-16 flex items-center justify-between gap-6">
        <Link href="/" className="shrink-0">
          <Logo />
        </Link>

        <nav className="hidden md:flex items-center gap-8 text-sm font-medium text-ink/70 flex-1 justify-center">
          <Link href="/jobs" className="hover:text-ink transition-colors">
            Cari Lowongan
          </Link>
        </nav>

        <div className="flex items-center gap-3 shrink-0">
          <Link href="/login" className="text-sm font-medium text-ink/70 hover:text-ink transition-colors">
            Masuk
          </Link>
          <Link
            href="/register"
            className="text-sm font-medium bg-ink text-white px-4 py-2 rounded-full hover:bg-ink/90 transition-colors"
          >
            Daftar
          </Link>
        </div>
      </div>
    </header>
  );
}