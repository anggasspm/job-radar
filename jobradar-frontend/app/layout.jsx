import { Space_Grotesk, Inter, IBM_Plex_Mono } from "next/font/google";
import "./globals.css";
import Header from "./components/Header";

const display = Space_Grotesk({
  subsets: ["latin"],
  variable: "--font-display",
  weight: ["500", "700"],
});
const body = Inter({ subsets: ["latin"], variable: "--font-body" });
const mono = IBM_Plex_Mono({
  subsets: ["latin"],
  variable: "--font-mono",
  weight: ["400", "500"],
});

export const metadata = {
  title: "JobRadar — Pencarian kerja yang paham bahasamu",
  description:
    "JobRadar memindai ribuan lowongan dari Glints, Tech in Asia, dan We Work Remotely, lalu membantumu menemukan yang cocok lewat kalimat biasa.",
};

export default function RootLayout({ children }) {
  return (
    <html
      lang="id"
      className={`${display.variable} ${body.variable} ${mono.variable}`}
    >
      <body className="font-body bg-canvas text-ink antialiased min-h-screen flex flex-col">
        <Header />
        <div className="flex-1">{children}</div>
      </body>
    </html>
  );
}
