package main

import (
    "database/sql"
    "fmt"
)

func cmd_kapal_penggunaan() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_KAPAL_Penggunaan")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_kapal_overkapasitas() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_KAPAL_Overkapasitas")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_kapal_penggunaanTanggal() error {
    var err error
    var data *sql.Rows

    fmt.Printf("Masukkan tanggal (YYYY-MM-DD) : ")
    tgl := input(lineInput)

    data, err = db.Query("SELECT * FROM V_KAPAL_PenggunaanTanggal WHERE Tanggal = @p1", tgl)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_kapal_digunakanTanggal() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT KdKapal, Nama FROM KAPAL")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdk int
    fmt.Printf("Masukkan KdKapal : ")
    fmt.Sscanf(input(lineInput), "%d", &kdk)

    fmt.Printf("Masukkan tanggal (YYYY-MM-DD) : ")
    tgl := input(lineInput)

    data, err = db.Query("SELECT * FROM V_KAPAL_DigunakanTanggal WHERE KdKapal = @p1 AND Tanggal = @p2", kdk, tgl)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

