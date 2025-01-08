# Warehouse Management System

## Deskripsi Proyek

Warehouse Management System adalah aplikasi web yang memungkinkan pengguna untuk mengelola inventaris produk dengan mudah. Aplikasi ini dibangun menggunakan **Vue3** untuk front-end dan **Go** untuk back-end. Pengguna dapat melakukan login untuk mengakses dashboard yang menampilkan total stok, jumlah produk, dan item dengan stok rendah. Selain itu, pengguna dapat menambahkan, memperbarui, menghapus produk, serta mengekspor data ke CSV dan mengunduh barcode berdasarkan SKU.

## Fitur

- **Login**: Pengguna dapat masuk ke aplikasi untuk mengakses dashboard.
- **Dashboard**: Menampilkan total stok, jumlah produk, dan item dengan stok rendah.
- **Manajemen Produk**:
  - **Daftar Produk**: Menampilkan semua produk dalam inventaris.
  - **Menambahkan Produk**: Memungkinkan pengguna untuk menambahkan produk baru.
  - **Memperbarui Produk**: Memungkinkan pengguna untuk memperbarui detail produk yang ada.
  - **Menghapus Produk**: Memungkinkan pengguna untuk menghapus produk dari inventaris.
  - **Ekspor Data**: Mengizinkan pengguna untuk mengekspor data produk ke format CSV.
  - **Unduh Barcode**: Menghasilkan dan mengunduh barcode berdasarkan SKU produk.

## Setup Proyek

### 1. Persiapan Back-End (Go)

1. **Clone Repository**:
   ```bash
   git clone <repository-url>
   cd warehouse-management-system/backend
2. **Inisialisasi Modul Go:**:
   ```bash
   go mod init warehouse-management-system
3. **Jalankan Server:**:
   ```bash
   go run main.go
