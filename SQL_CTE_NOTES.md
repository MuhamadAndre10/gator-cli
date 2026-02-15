# Belajar SQL Common Table Expressions (CTE)

## Apa itu CTE?

CTE (Common Table Expression) adalah hasil query sementara yang bisa kamu beri nama dan gunakan dalam query utama (`SELECT`, `INSERT`, `UPDATE`, atau `DELETE`). Pikirkan CTE seperti membuat variables atau tabel temporary yang hanya hidup selama eksekusi satu query tersebut.

Sintaks dasarnya menggunakan kata kunci `WITH`.

**Contoh Dasar:**

```sql
WITH sales_januari AS (
    SELECT product_id, amount
    FROM sales
    WHERE month = 'january'
)
SELECT sum(amount) FROM sales_januari;
```

## Mengapa Menggunakan CTE?

1. **Keterbacaan (Readability):** Memecah query yang kompleks dan panjang menjadi bagian-bagian kecil yang logis. Lebih mudah dibaca daripada _nested subqueries_.
2. **DRY (Don't Repeat Yourself):** Kamu bisa mendefinisikan CTE sekali dan mereferensikannya berkali-kali dalam query utama.
3. **Data-Modifying (Yang kamu gunakan):** Melakukan `INSERT`/`UPDATE` dan langsung menggunakan hasil kembaliannya di langkah berikutnya.

---

## Studi Kasus: `CreateFeedUser`

Dalam file `sql/queries/feed-users.sql`, kamu menggunakan teknik **Data-Modifying CTE**. Ini adalah fitur powerful di PostgreSQL.

```sql
WITH insert_feed_user AS (
    -- Langkah 1: Lakukan Insert
    INSERT INTO feed_follow(id, user_id, feed_id)
    VALUES($1, $2, $3)
    RETURNING *  -- 'RETURNING *' membuat INSERT ini mengembalikan data baris yang baru dibuat
)
-- Langkah 2: Gunakan hasil insert tersebut
SELECT
    ifu.*,               -- Ambil semua kolom dari baris baru (dari CTE)
    u.name AS user_name, -- Join untuk dapat nama user
    f.name AS feed_name  -- Join untuk dapat nama feed
FROM insert_feed_user ifu
    JOIN users u ON ifu.user_id = u.id
    JOIN feed f ON ifu.feed_id = f.id;
```

### Penjelasan Alur:

1. **Di dalam `WITH`:** Database mengeksekusi `INSERT`. Karena ada `RETURNING *`, data baris yang baru saja masuk ke tabel `feed_follow` "ditangkap" dan disimpan sementara dalam CTE bernama `insert_feed_user`.
2. **Di dalam `SELECT` Utama:** Kita memperlakukan `insert_feed_user` seolah-olah itu adalah tabel biasa.
3. **Langsung JOIN:** Kita bisa langsung men-`JOIN` baris yang baru saja dibuat tersebut dengan tabel `users` dan `feeds`.

### Keuntungan Cara Ini:

Tanpa CTE, kamu mungkin harus melakukan 2 langkah terpisah (2 round-trip ke database):

1. `INSERT INTO ... RETURNING id` (simpan ID di aplikasi Go).
2. `SELECT ... FROM feed_follow JOIN ... WHERE id = $1`.

Dengan CTE, kamu melakukan Insert + Select detail tambahannya hanya dalam **satu query tunggal**. Ini lebih efisien (lebih sedikit overhead jaringan) dan atomik.
