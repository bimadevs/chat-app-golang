# Chat App Golang

Chat App Golang adalah aplikasi chatting real-time sederhana yang dibangun menggunakan **Go** untuk backend dan **HTML/JavaScript dengan Tailwind CSS** untuk frontend. Aplikasi ini mendukung autentikasi pengguna sederhana, pesan real-time melalui WebSocket, dan status online/offline pengguna.


[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8.svg)](https://golang.org/)


## Fitur
- **Autentikasi Pengguna**: Registrasi, login, dan profil pengguna dengan JWT.
- **Real-Time Chat**: Mengirim dan menerima pesan secara instan melalui WebSocket.
- **Status Online/Offline**: Menampilkan daftar pengguna yang sedang online dan notifikasi saat pengguna bergabung/keluar.
- **Hapus Riwayat Chat**: Fitur untuk mengosongkan semua pesan dari database.
- **UI Modern**: Desain responsif dan estetis menggunakan Tailwind CSS.

## Teknologi
- **Backend**: Go, GORM (SQLite), Gorilla WebSocket, JWT, bcrypt
- **Frontend**: HTML, JavaScript, Tailwind CSS
- **Database**: SQLite

## Struktur Folder
```
chat-app/
├── cmd/
│   └── server/
│       └── main.go         # Entry point aplikasi
├── internal/
│   ├── auth/              # Logika autentikasi (JWT)
│   ├── config/            # Konfigurasi aplikasi
│   ├── database/          # Koneksi database
│   ├── handlers/          # Handler HTTP
│   ├── models/            # Definisi model data
│   └── websocket/         # Logika WebSocket
├── views/
│   └── index.html         # File frontend
├── go.mod                 # Dependensi Go
└── README.md              # Dokumentasi ini
```

## Prasyarat
- [Go](https://golang.org/dl/) (versi 1.16 atau lebih baru)
- Browser modern (Chrome, Firefox, dll.)

## Instalasi
1. Clone repository ini:
   ```bash
   git clone https://github.com/bimadevs/chat-app-golang.git
   cd chat-app-golang
   ```

2. Instal dependensi Go:
   ```bash
   go mod tidy
   ```

3. Jalankan aplikasi:
   ```bash
   go run cmd/server/main.go
   ```

4. Buka browser di `http://localhost:8080`.

## Penggunaan
1. **Registrasi**: Masukkan username dan password di form, lalu klik "Register".
2. **Login**: Gunakan kredensial yang sama untuk login.
3. **Chat**: Setelah login, Anda dapat mengirim pesan, melihat pengguna online, dan menghapus riwayat chat.
4. **Logout**: Klik tombol "Logout" untuk keluar.

## Lisensi
Proyek ini dilisensikan di bawah [MIT License](LICENSE).

## Kontak
Jika ada pertanyaan atau saran, silakan buka issue atau hubungi saya di [bimaj0206@gmail.com](mailto:bimaj0206@gmail.com).

atau dapat hubungi saya melalui whatsapp [082254044783][https://wa.me/6282254044783]
