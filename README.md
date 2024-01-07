# Pemesanan Ticket Wisata Backend API - Daerah Garut

## Deskripsi Proyek

Proyek ini adalah backend API untuk sistem pemesanan tiket wisata yang dibangun menggunakan teknologi Golang. Aplikasi ini dirancang untuk menyediakan layanan pemesanan tiket wisata secara efisien dan aman.

## Teknologi Utama

1. **Golang**: Digunakan sebagai bahasa pemrograman utama untuk mengembangkan backend API. Golang dipilih karena kinerjanya yang tinggi, kemudahan dalam pengelolaan kode, dan dukungan untuk pengembangan aplikasi skala besar.

2. **Gofiber**: Merupakan framework web yang ringan dan cepat untuk Golang. Gofiber digunakan untuk membangun endpoint-endpoint API dengan kinerja yang optimal.

3. **Postgres**: Database relasional Postgres digunakan untuk menyimpan dan mengelola data terkait pemesanan tiket wisata. Ini memberikan keandalan dan fleksibilitas dalam pengelolaan data.

4. **Midtrans - Core**: Diperintahkan untuk menangani proses pembayaran. Midtrans adalah gateway pembayaran yang terintegrasi, memungkinkan aplikasi menerima pembayaran dengan berbagai metode pembayaran.

5. **Cloudinary**: Digunakan sebagai cloud storage untuk mengelola dan menyimpan file media, seperti gambar terkait destinasi wisata.

6. **JWT (JSON Web Token)**: Digunakan untuk otentikasi dan otorisasi pengguna. JWT memberikan cara aman untuk mentransmisikan informasi otentikasi antara pihak-pihak yang terlibat.

## Cara Menjalankan Proyek

1. Pastikan Golang sudah terinstal di sistem Anda.
2. Buat file `.env` di direktori proyek, serta sesuaikan konfigurasi database Postgres dan API Midtrans.
3. Install dependensi dengan menggunakan perintah `go mod tidy`.
4. Sesuaikan konfigurasi database Postgres dan API Midtrans di file konfigurasi.
5. Jalankan aplikasi dengan perintah `go run main.go`.
6. Backend API akan berjalan pada `http://localhost:8080` secara default.


## Documentation

### Authentication
![Screenshot from 2024-01-08 01-34-36](https://github.com/RianIhsan/go-backend-ulinan/assets/93025581/0c7f2a17-2551-4540-be35-fe0cb3288260)

### Dashboard page
![Screenshot from 2024-01-08 01-34-26](https://github.com/RianIhsan/go-backend-ulinan/assets/93025581/0091009e-5aa8-456c-bba4-8e63dddbc546)

### Admin page
![Screenshot from 2024-01-08 01-35-48](https://github.com/RianIhsan/go-backend-ulinan/assets/93025581/affcf7cb-ba31-493c-9c68-8d5ee3f6f7aa)

### Payment page
![Screenshot from 2024-01-08 01-36-50](https://github.com/RianIhsan/go-backend-ulinan/assets/93025581/e2695dae-bf77-4ba9-8988-fe730ec1ad70)



