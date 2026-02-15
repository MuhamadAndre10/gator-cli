# Belajar SQL Many-to-Many Relationship

## Apa itu Many-to-Many?

Many-to-Many adalah relasi dimana **satu baris di Tabel A bisa berhubungan dengan banyak baris di Tabel B**, dan sebaliknya.

**Contoh nyata:**

- Satu **User** bisa follow banyak **Feed**.
- Satu **Feed** bisa di-follow oleh banyak **User**.

## Bagaimana Cara Membuatnya?

Kita **tidak bisa** langsung menghubungkan 2 tabel secara Many-to-Many di SQL. Kita butuh **Junction Table** (tabel penghubung / tabel perantara) di tengahnya.

```
users ←──── feed_follow ────→ feed
(1 user)    (junction)     (1 feed)
            banyak baris
```

Jadi, Many-to-Many sebenarnya adalah **dua relasi One-to-Many** yang bertemu di junction table.

---

## Studi Kasus: Blog Aggregator

### Skema Tabel

```sql
-- Tabel Utama 1: users
CREATE TABLE users (
    id uuid PRIMARY KEY,
    name varchar(100) UNIQUE NOT NULL
);

-- Tabel Utama 2: feed
CREATE TABLE feed (
    id uuid PRIMARY KEY,
    name varchar(100) UNIQUE NOT NULL,
    url varchar(100) UNIQUE NOT NULL,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- Junction Table: feed_follow (penghubung users <-> feed)
CREATE TABLE feed_follow (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    feed_id uuid NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (feed_id) REFERENCES feed(id) ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)  -- Mencegah user follow feed yang sama 2x
);
```

### Elemen Penting pada Junction Table

| Elemen                     | Fungsi                                              |
| -------------------------- | --------------------------------------------------- |
| `user_id` (FK)             | Menghubungkan ke tabel `users`                      |
| `feed_id` (FK)             | Menghubungkan ke tabel `feed`                       |
| `UNIQUE(user_id, feed_id)` | Mencegah duplikasi relasi                           |
| `ON DELETE CASCADE`        | Jika user/feed dihapus, record follow ikut terhapus |
| `INDEX` pada FK            | Mempercepat query JOIN                              |

---

## Query CRUD pada Many-to-Many

### 1. CREATE — Follow Feed

```sql
INSERT INTO feed_follow(id, user_id, feed_id)
VALUES($1, $2, $3)
RETURNING *;
```

### 2. READ — Tampilkan Feed yang Di-follow User

Query ini menggunakan **JOIN melalui junction table** untuk menghubungkan `users` dan `feed`.

```sql
SELECT f.name AS feed_name,
    f.url AS feed_url,
    u.name AS user_name
FROM feed_follow ff
    JOIN users u ON ff.user_id = u.id
    JOIN feed f ON ff.feed_id = f.id
WHERE ff.user_id = $1;  -- Filter berdasarkan user tertentu
```

**Cara baca query ini:**

1. Mulai dari junction table `feed_follow` (alias `ff`).
2. JOIN ke `users` untuk dapat nama user.
3. JOIN ke `feed` untuk dapat nama dan URL feed.
4. Filter hanya untuk user tertentu (`WHERE ff.user_id = $1`).

### 3. DELETE — Unfollow Feed

```sql
DELETE FROM feed_follow
WHERE user_id = $1
    AND feed_id = $2;
```

Perlu **dua kondisi** (`user_id` DAN `feed_id`) karena kita menghapus satu relasi spesifik.

---

## Perbedaan One-to-Many vs Many-to-Many

| Aspek            | One-to-Many                             | Many-to-Many                                                     |
| ---------------- | --------------------------------------- | ---------------------------------------------------------------- |
| **FK di mana?**  | Langsung di tabel anak (`feed.user_id`) | Di junction table (`feed_follow.user_id`, `feed_follow.feed_id`) |
| **Jumlah tabel** | 2 tabel                                 | 3 tabel (2 utama + 1 junction)                                   |
| **Contoh**       | User membuat Feed                       | User mem-follow Feed                                             |
| **Arti**         | Kepemilikan / Pencipta                  | Hubungan / Langganan                                             |

---

## Tips Penting

1. **Selalu tambahkan `UNIQUE` constraint** pada kombinasi FK di junction table untuk mencegah data duplikat.
2. **Buat INDEX** pada kolom FK di junction table (`user_id`, `feed_id`) agar query JOIN lebih cepat.
3. **Gunakan `ON DELETE CASCADE`** agar data di junction table otomatis terhapus saat data utama dihapus.
4. Junction table bisa punya **kolom tambahan** seperti `created_at` untuk menyimpan kapan relasi dibuat.
