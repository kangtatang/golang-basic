## CONTOH SEDERHANA APLIKASI BACKEND

**Deskripsi**: contoh sederhana backend untuk manajemen karyawan (CRUD)

Tools:

- Golang: 1.21.5
- gin v1.10.0
- gorm v1.9.16
- godotenv v1.5.

**Installed but not yet used**

- golang-jwt/jwt/v4 v4.5.0
- jwt-go v3.2.0+incompatible

### Instalasi dependency jika anda clone repo ini

```bash

go get -u ./...

```

### Endpoint

`POST /employees`

### Deskripsi

Endpoint ini digunakan untuk membuat karyawan baru dalam database.

### Parameter Permintaan

| Nama         | Tipe   | Wajib | Deskripsi              |
| ------------ | ------ | ----- | ---------------------- |
| name         | String | Ya    | Nama karyawan          |
| email        | String | Ya    | Email karyawan         |
| phone_number | String | Tidak | Nomor telepon karyawan |
| join_date    | Date   | Ya    | Tanggal masuk karyawan |
| status       | String | Tidak | Status karyawan        |
| address      | String | Tidak | Alamat karyawan        |
| nik          | String | Ya    | Nomor Induk Karyawan   |
| gender       | String | Tidak | Jenis kelamin karyawan |
| position     | String | Tidak | Jabatan karyawan       |

### Contoh Permintaan

```json
{
  "name": "Jane Grey",
  "email": "jane.grey@example.com",
  "phone_number": "5550123456",
  "join_date": "2023-11-01T00:00:00Z",
  "status": "active",
  "address": "808 Poplar St, Anyplace, USA",
  "nik": "0123456789",
  "gender": "female",
  "position": "HR Manager"
}
```

### Respons

json

Salin kode

```json
{
  "id": 11,
  "name": "Jane Grey",
  "email": "jane.grey@example.com",
  "phone_number": "5550123456",
  "join_date": "2023-11-01T00:00:00Z",
  "status": "active",
  "address": "808 Poplar St, Anyplace, USA",
  "nik": "0123456789",
  "gender": "female",
  "position": "HR Manager",
  "created_at": "2024-05-27T20:12:47.5674808+07:00",
  "updated_at": "2024-05-27T20:12:47.5674808+07:00"
}
```

## Mendapatkan Semua Karyawan

### Endpoint

`GET /employees`

### Deskripsi

Endpoint ini digunakan untuk mendapatkan semua karyawan yang tersimpan dalam database.

### Respons

```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "join_date": "2024-05-30",
    "nik": "123456789",
    "created_at": "2024-05-30T12:00:00Z",
    "updated_at": "2024-05-30T12:00:00Z"
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane@example.com",
    "join_date": "2024-06-01",
    "nik": "987654321",
    "created_at": "2024-06-01T12:00:00Z",
    "updated_at": "2024-06-01T12:00:00Z"
  }
]
```

nama database: `be_karyawan`
table: `employees`

**create DB in postgre:**

```bash
psql -U postgres
CREATE DATABASE be_karyawan;
\l
\q
```

Contoh penggunaan:

```bash

$ psql -U postgres
Password: (masukkan password Anda)
psql (12.5)
Type "help" for help.

postgres=# CREATE DATABASE be_karyawan;
CREATE DATABASE
postgres=# \l
                                  List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
-----------+----------+----------+------------+------------+-----------------------
 be_karyawan | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
(4 rows)

postgres=# \q


```

Create table employee sql schema:

```sql

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(50),
    join_date TIMESTAMP,
    status VARCHAR(50),
    address TEXT,
    nik VARCHAR(50) NOT NULL UNIQUE,
    gender VARCHAR(50),
    position VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


```

## GET ALL

`GET /employees`

## VIEW

`GET /employees/:id`

## DELETE

`DELETE /employees/:id`

## UPDATE

`PUT /employees/:id`

## URL

`http://localhost:8080/`
