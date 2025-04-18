# 💧 Sistem Rekomendasi Galon Terbaik di Wilayah Baubau
Menggunakan **Metode Simple Additive Weighting (SAW)**

Sistem ini membantu merekomendasikan depot galon terbaik di wilayah Baubau berdasarkan berbagai kriteria menggunakan metode SAW. Dibangun menggunakan framework [goarchi](https://github.com/zayn1510/goarchi) berbasis Golang.

---

## 🚀 Fitur

- ✅ Manajemen data kriteria
- ✅ Manajemen data depot
- ✅ Manajemen data user
- ✅ Manajemen data lokasi user
- ✅ Manajemen data kecamatan
- ✅ Perhitungan metode Simple Additive Weighting (SAW)

---

## 📦 Instalasi

### 1. Clone Repository

```bash
git clone https://github.com/zayn1510/galon_metode_saw.git
cd galon_metode_saw
```

### 2. Setup Environment

Salin file `.env.example` menjadi `.env`:

```bash
cp .env.example .env
```

Lalu sesuaikan konfigurasi database:

```
DB_NAME=db_metode_saw_depot
DB_HOST=localhost
DB_USER=root       # ubah sesuai user MySQL kamu
DB_PASS=1234       # ubah sesuai password MySQL kamu
DB_PORT=3329       # port lokal
DB_PREFIX=tbl
```

---


### 3. Jalankan dengan Docker

Edit file `docker-compose.yml`, sesuaikan environment MySQL:

```yaml
environment:
  - DB_NAME=db_metode_saw_depot
  - DB_HOST=db
  - DB_USER=root
  - DB_PASS=1234
  - DB_PORT=3306
  - DB_PREFIX=tbl
```

> **Catatan:** Jangan lupa sesuaikan juga `user` dan `password` pada service MySQL.

### 4. Build & Jalankan

#### 🐧 Linux / macOS

```bash
./.goarchi build
docker compose build
docker compose up -d
```

#### 🪟 Windows

```bash
go run cli/main.go build
docker compose build
docker compose up -d
```

---

### 5. Migrasi Database

#### 🐧 Linux / macOS

```bash
./.goarchi make migrate up
```

#### 🪟 Windows

```bash
goarchi.exe make migrate up
```

---

## 🌐 API Endpoint

> Sesuaikan dengan port yang digunakan pada `docker-compose`, default: `localhost:8024`

### 📘 Welcome
| Method | Endpoint            | Keterangan          |
|--------|---------------------|---------------------|
| GET    | `/api/v1/welcome`   | Cek koneksi API     |

### 📘 Kriteria
| Method | Endpoint                | Keterangan           |
|--------|-------------------------|----------------------|
| GET    | `/api/v1/kriteria`      | Ambil semua data     |
| POST   | `/api/v1/kriteria`      | Buat data            |
| PUT    | `/api/v1/kriteria/{id}` | Update data          |
| DELETE | `/api/v1/kriteria/{id}` | Hapus data           |

### 📘 Depot
| Method | Endpoint             | Keterangan           |
|--------|----------------------|----------------------|
| GET    | `/api/v1/depot`      | Ambil semua data     |
| POST   | `/api/v1/depot`      | Buat data            |
| PUT    | `/api/v1/depot/{id}` | Update data          |
| DELETE | `/api/v1/depot/{id}` | Hapus data           |

### 📘 User
| Method | Endpoint            | Keterangan           |
|--------|---------------------|----------------------|
| GET    | `/api/v1/user`      | Ambil semua data     |
| POST   | `/api/v1/user`      | Buat data            |
| PUT    | `/api/v1/user/{id}` | Update data          |
| DELETE | `/api/v1/user/{id}` | Hapus data           |

### 📘 Lokasi User
| Method | Endpoint                        | Keterangan           |
|--------|----------------------------------|----------------------|
| GET    | `/api/v1/user-locations`        | Ambil semua data     |
| POST   | `/api/v1/user-locations`        | Buat data            |
| PUT    | `/api/v1/user-locations/{id}`   | Update data          |
| DELETE | `/api/v1/user-locations/{id}`   | Hapus data           |

### 📘 Kecamatan
| Method | Endpoint               | Keterangan           |
|--------|------------------------|----------------------|
| GET    | `/api/v1/kecamatan`    | Ambil semua data     |
| POST   | `/api/v1/kecamatan`    | Buat data            |
| PUT    | `/api/v1/kecamatan/{id}` | Update data        |
| DELETE | `/api/v1/kecamatan/{id}` | Hapus data         |

---

## 🧪 Koleksi Postman

Untuk mempermudah pengujian API, kamu bisa menggunakan file koleksi Postman berikut:

📥 [Download Koleksi Postman](postman/galon-api.postman_collection.json)

**Cara Import:**
1. Buka Postman.
2. Klik tombol `Import`.
3. Pilih file `.json` dari folder `postman/`.
4. Jalankan request sesuai endpoint yang tersedia.

## 📄 Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).

---

## 🙌 Terima Kasih

Framework ini dibangun dengan semangat open-source dan kolaborasi.  
Semoga bermanfaat untuk proyek-proyek kamu! ✨