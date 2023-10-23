package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Buat koneksi ke database MySQL
    db, err := sql.Open("mysql", "wulan:wulwul21@tcp(localhost:3306)/db_mahasiswa")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Tes koneksi ke database
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    // Query SQL untuk membaca data dari tabel
    rows, err := db.Query("SELECT id, nama FROM tb_mahasiswa")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Baca data hasil query
    for rows.Next() {
        var id int
        var nama string
        err = rows.Scan(&id, &nama)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s\n", id, nama)
    }
}
