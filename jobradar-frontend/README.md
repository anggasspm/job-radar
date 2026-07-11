# JobRadar Frontend

Frontend Next.js untuk JobRadar, disambungkan ke backend Go yang sudah
di-deploy di `https://anggasspm-jobradar.hf.space`.

## Menjalankan secara lokal

```bash
npm install
npm run dev
```

Buka `http://localhost:3000`. Variabel `NEXT_PUBLIC_API_BASE_URL` di
`.env.local` sudah diarahkan ke backend yang di-deploy — ganti kalau kamu
menjalankan backend sendiri di lokal (`http://localhost:8080`).

## Yang sudah tersambung dan bekerja

- **Cari & jelajah lowongan** (`/jobs`) — `GET /jobs` dan
  `GET /jobs/search?q=...`, dengan filter tambahan (kategori, lokasi, gaji,
  pengalaman) yang dihitung di sisi frontend dari hasil pencarian.
- **Detail lowongan** (`/jobs/:id`) — `GET /jobs/:id`, termasuk penanganan
  404 dan kegagalan koneksi yang jelas.
- **Daftar & masuk** (`/register`, `/login`) — `POST /auth/register` dan
  `POST /auth/login`, token disimpan di `localStorage` dengan pelacakan
  waktu kedaluwarsa.
- **Favorit** (`/favorites`) — `GET /favorite/` (perlu login).

## Keterbatasan yang jujur perlu diketahui

Ini bukan bug di frontend — ini kondisi nyata dari backend saat ini:

1. **Token akses berumur 15 menit** dan backend belum punya endpoint
   refresh token yang aktif (field `RefreshToken` di DTO backend memang
   sengaja tidak pernah dikirim ke klien — lihat `json:"-"` di
   `dto/userRequestDto.go`). Jadi setelah 15 menit, pengguna akan diminta
   masuk ulang. Frontend menampilkan penghitung waktu di header saat sesi
   akan berakhir, dan otomatis membersihkan sesi begitu server membalas
   401.
2. **Menambah/menghapus favorit belum berfungsi.** Handler
   `AddToFavorites` dan `DeleteFromFavorites` di backend
   (`favoriteHandler.go`) masih kosong, dan rute-nya bahkan belum
   didaftarkan di `favoriteRoute.go`. Halaman `/favorites` hanya bisa
   menampilkan daftar yang sudah ada lewat `GET /favorite/` — tombol
   "simpan lowongan" sengaja belum ditambahkan ke UI supaya tidak
   berpura-pura berfungsi. Setelah endpoint POST/DELETE dibuat dan
   didaftarkan di backend, tambahkan pemanggilannya di
   `app/lib/api.js` lalu tombolnya bisa ditambahkan ke `JobCard`.
3. **Cold start.** Backend berjalan di Hugging Face Space tier gratis,
   yang bisa "tidur" setelah tidak ada trafik. Permintaan pertama setelah
   idle lama bisa butuh beberapa detik — frontend menampilkan pesan yang
   sesuai ("server mungkin sedang bangun dari tidur") alih-alih pesan
   error generik saat ini terjadi.

## Struktur

```
app/
  lib/
    api.js         # satu-satunya tempat pemanggilan fetch ke backend
    auth.js         # penyimpanan sesi + kedaluwarsa token
    useSession.js   # hook reaktif untuk status login
  components/       # JobCard, FilterPanel, Header, dll — dipakai ulang
  jobs/             # pencarian + detail lowongan
  login/ register/  # autentikasi
  favorites/        # daftar favorit (read-only, lihat keterbatasan di atas)
```
