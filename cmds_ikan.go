package main

import (
    "database/sql"
    "fmt"
)

func cmd_ikan_dibudidayakan() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_IKAN_Dibudidayakan")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_ikan_diLaut() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT KdIkan, Nama FROM IKAN")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdi int
    fmt.Printf("Masukkan KdIkan : ")
    fmt.Sscanf(input(lineInput), "%d", &kdi)

    data, err = db.Query("SELECT * FROM V_IKAN_DiLaut WHERE KdIkan = @p1", kdi)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_ikan_jumlahPerLaut() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_IKAN_JumlahPerLaut")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

