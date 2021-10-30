package main

import (
    "database/sql"
    "fmt"
)

func cmd_terdekat_supplier() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT KdPelanggan, Nama FROM PELANGGAN")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdPelanggan : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    data, err = db.Query("SELECT * FROM V_TERDEKAT_Supplier WHERE KdPelanggan = @p1", kdp)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_terdekat_pabrik() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT KdSupplier, Nama FROM SUPPLIER")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kds int
    fmt.Printf("Masukkan KdSupplier : ")
    fmt.Sscanf(input(lineInput), "%d", &kds)

    data, err = db.Query("SELECT * FROM V_TERDEKAT_Pabrik WHERE KdSupplier = @p1", kds)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_terdekat_tambak() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT KdPabrik, Nama FROM PABRIK")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdPabrik : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    data, err = db.Query("SELECT * FROM  V_TERDEKAT_Tambak WHERE KdPabrik = @p1", kdp)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

