# Warehouse Management System

## Deskripsi Proyek

Warehouse Management System adalah aplikasi web yang memungkinkan pengguna untuk mengelola inventaris produk dengan mudah. Aplikasi ini dibangun menggunakan **Vue.js** untuk front-end dan **Go** untuk back-end. Pengguna dapat melakukan login untuk mengakses dashboard yang menampilkan total stok, jumlah produk, dan item dengan stok rendah. 

Aplikasi ini memiliki sistem peran yang membatasi akses ke menu produk. Hanya pengguna dengan peran **Admin** dan **Manager** yang dapat mengakses menu produk untuk melakukan operasi seperti menambahkan, memperbarui, dan menghapus produk. Pengguna dengan peran lain hanya dapat mengakses dashboard.

## Fitur

- **Login**: Pengguna dapat masuk ke aplikasi untuk mengakses dashboard.
- **Dashboard**: Menampilkan total stok, jumlah produk, dan item dengan stok rendah.
- **Manajemen Produk** (Hanya untuk Admin dan Manager):
  - **Daftar Produk**: Menampilkan semua produk dalam inventaris.
  - **Menambahkan Produk**: Memungkinkan pengguna untuk menambahkan produk baru.
  - **Memperbarui Produk**: Memungkinkan pengguna untuk memperbarui detail produk yang ada.
  - **Menghapus Produk**: Memungkinkan pengguna untuk menghapus produk dari inventaris.
  - **Ekspor Data**: Mengizinkan pengguna untuk mengekspor data produk ke format CSV.
  - **Unduh Barcode**: Menghasilkan dan mengunduh barcode berdasarkan SKU produk.
- **Sistem Peran**: Mengelola akses pengguna berdasarkan peran (Admin, Manager, dan pengguna biasa).

## Setup Proyek

### 1. Persiapan Database SQL

Sebelum menjalankan aplikasi, Anda perlu menyiapkan database SQL. Berikut adalah langkah-langkah untuk membuat database dan tabel yang diperlukan.

#### a. **Buat Database**

```sql
CREATE DATABASE warehouse;
USE warehouse;
```

#### b. **Buat Tabel Produk**

```sql
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(100) NOT NULL UNIQUE,
    qty INT NOT NULL,
    location VARCHAR(255) NOT NULL,
    status ENUM('available', 'out_of_stock') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

#### c. **Buat Tabel users**

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role ENUM('admin', 'manager', 'user') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

#### d. **Insert Data Awal**

```sql
-- Insert data ke tabel products (Opsional)
INSERT INTO products (name, sku, qty, location, status) VALUES
('Product A', 'SKU123', 10, 'Warehouse 1', 'available'),
('Product B', 'SKU124', 5, 'Warehouse 2', 'available'),
('Product C', 'SKU125', 0, 'Warehouse 1', 'out_of_stock');

-- Insert data ke tabel users
INSERT INTO users (username, password, role) VALUES
('admin', 'admin_password', 'admin'),
('manager', 'manager_password', 'manager'),
('user', 'user_password', 'user');
```

### 2. Persiapan Back-End (Go)
1. **Clone Repository**:
   ```bash
   git clone <repository-url>
   cd warehouse-management-system/backend
   ```
2. **Inisialisasi Modul Go:**:
   ```bash
   go mod init warehouse
   ```
3. **Jalankan Server:**:
   ```bash
   go run main.go
   ```

### 2. Persiapan Front-End (Vue.js)
1. **Navigasi ke Folder Front-End:**:
   ```bash
    cd warehouse-management-system/frontend
   ```
2. **Instalasi Dependensi:**:
   ```bash
   npm install
   ```
3. **Jalankan Server:**:
   ```bash
    npm run serve
   ```
