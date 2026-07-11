"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import { Bookmark, LogOut } from "lucide-react";
import Logo from "./Logo";
import SessionCountdown from "./SessionCountdown";
import { useSession } from "../lib/useSession";

export default function Header() {
  const { user, expiresAt, ready, isAuthenticated, logout } = useSession();
  const router = useRouter();

  function handleLogout() {
    logout();
    router.push("/");
  }

  return (
    <header className="sticky top-0 z-50 w-full bg-canvas/90 backdrop-blur-md border-b border-line">
      <div className="max-w-6xl mx-auto px-6 h-16 flex items-center justify-between gap-6">
        <Link href="/" className="shrink-0">
          <Logo />
        </Link>

        <nav className="hidden md:flex items-center gap-8 text-sm font-medium text-ink-soft flex-1 justify-center">
          <Link href="/jobs" className="hover:text-ink transition-colors">
            Cari lowongan
          </Link>
          {isAuthenticated && (
            <Link
              href="/favorites"
              className="hover:text-ink transition-colors flex items-center gap-1.5"
            >
              <Bookmark className="w-3.5 h-3.5" />
              Favorit
            </Link>
          )}
        </nav>

        <div className="flex items-center gap-3 shrink-0 min-w-[140px] justify-end">
          {!ready ? null : isAuthenticated ? (
            <>
              <span className="hidden sm:flex items-center gap-1.5 text-sm text-ink-soft truncate max-w-[160px]">
                {user?.name || user?.email}
                <SessionCountdown expiresAt={expiresAt} />
              </span>
              <button
                onClick={handleLogout}
                className="flex items-center gap-1.5 text-sm font-medium text-ink-soft hover:text-ink transition-colors px-3 py-2 rounded-full hover:bg-white"
              >
                <LogOut className="w-3.5 h-3.5" />
                Keluar
              </button>
            </>
          ) : (
            <>
              <Link
                href="/login"
                className="text-sm font-medium text-ink-soft hover:text-ink transition-colors"
              >
                Masuk
              </Link>
              <Link
                href="/register"
                className="text-sm font-medium bg-ink text-white px-4 py-2 rounded-full hover:bg-ink/90 transition-colors"
              >
                Daftar
              </Link>
            </>
          )}
        </div>
      </div>
    </header>
  );
}
