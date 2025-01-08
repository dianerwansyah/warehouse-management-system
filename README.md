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
    status VARCHAR(20) NOT NULL,
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
    role VARCHAR(20) NOT NULL,
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
a. **Clone Repository**:
   ```bash
   git clone <repository-url>
   cd warehouse-management-system/backend
   ```
b. **Inisialisasi Modul Go:**:
   ```bash
   go mod init warehouse
   ```
c. **Jalankan Server:**:
   ```bash
   go run main.go
   ```

### 3. Persiapan Front-End (Vue.js)
a. **Navigasi ke Folder Front-End:**:
   ```bash
    cd warehouse-management-system/frontend
   ```
b. **Instalasi Dependensi:**:
   ```bash
   npm install
   ```
c. **Jalankan Server:**:
   ```bash
    npm run serve
   ```

### Dokumentasi API
1. **Login**:
Endpoint: /login
Method: POST
Request Body:
 ```bash
  {
    "username": "your_username",
    "password": "your_password"
  }
 ```
Response:
 ```bash
  {
    "token": "your_jwt_token",
    "user": "data_user",
  }
 ```

2. **Get Products**:
Endpoint: /products/gets
Method: GET
Response:
 ```bash
   [
    {
      "id": 1,
      "name": "Product A",
      "sku": "SKU123",
      "qty": 10,
      "location": "Warehouse 1",
      "status": "available"
    },
    ...
  ]
 ```

3. **Get Product by ID**:
Endpoint: /products/getbyid/{id}
Method: GET
Response:
 ```bash
    {
      "id": 1,
      "name": "Product A",
      "sku": "SKU123",
      "qty": 10,
      "location": "Warehouse 1",
      "status": "available"
    },
 ```

4. **Add Product**:
Endpoint: /products/add
Method: POST
Request Body:
 ```bash
   {
    "name": "Product A",
    "sku": "SKU123",
    "qty": 10,
    "location": "Warehouse 1",
    "status": "available"
  }
 ```

5. **Update Product**:
Endpoint: /products/update/{id}
Method: PUT
Request Body:
 ```bash
  {
    "name": "Updated Product A",
    "sku": "SKU123",
    "qty": 15,
    "location": "Warehouse 1",
    "status": "available"
  }
 ```

6. **Delete Product**:
Endpoint: /products/delete/{id}

Method: DELETE
Response:
 ```bash
  {
    "message": "Product deleted successfully"
  }
 ```

#### **Screenshots**

#### **Index/Login Page**
![image](https://github.com/user-attachments/assets/f4d87ea6-c65c-4922-91f4-d8454267d306)

#### **Dashboard**
![image](https://github.com/user-attachments/assets/d2217a02-a3ad-404a-b5ba-e523e07b496d)

#### **List Product**
![image](https://github.com/user-attachments/assets/91d771ef-e3f3-4da3-b633-7b212799a6a7)

#### **Add Product**
![image](https://github.com/user-attachments/assets/568d390d-e7a1-45cb-94c5-e42f8152a9bc)

#### **Edit Product**
![image](https://github.com/user-attachments/assets/9f008305-ebc7-4fd7-a85b-38f0bba92dfa)

#### **Delete Product**
![image](https://github.com/user-attachments/assets/babe4936-f869-4203-b176-31b4f05b1bda)



