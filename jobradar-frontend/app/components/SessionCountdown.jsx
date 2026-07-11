"use client";

import { useSyncExternalStore } from "react";

// "Sekarang" adalah sumber data eksternal yang berubah di luar kendali
// React. useSyncExternalStore adalah primitif yang tepat untuk membacanya:
// Date.now() dipanggil di dalam getSnapshot/subscribe, bukan di badan
// render maupun di badan effect secara langsung.
function subscribe(callback) {
  const interval = setInterval(callback, 15000);
  return () => clearInterval(interval);
}

function getNow() {
  return Date.now();
}

function getServerNow() {
  return 0;
}

export default function SessionCountdown({ expiresAt }) {
  const now = useSyncExternalStore(subscribe, getNow, getServerNow);

  if (!expiresAt) return null;
  const minutesLeft = Math.max(0, Math.round((expiresAt - now) / 60000));
  if (minutesLeft > 2) return null;

  return (
    <span
      title="Sesi akan berakhir dalam waktu dekat"
      className="text-[10px] font-mono text-amber bg-amber/10 px-1.5 py-0.5 rounded-md shrink-0"
    >
      {minutesLeft}m
    </span>
  );
}
