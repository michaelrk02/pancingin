package main

import (
    "database/sql"
    "fmt"
    "strings"
)

func cmd_kapal_penggunaan() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdKapal, Nama, COUNT(*) JmlHari FROM (\n")
    qb.WriteString("    SELECT K.KdKapal, K.Nama, NMK.Tanggal, COUNT(*) AS JmlTerpakai\n")
    qb.WriteString("    FROM KAPAL K\n")
    qb.WriteString("    JOIN NELAYAN_MENGGUNAKAN_KAPAL NMK ON NMK.KdKapal = K.KdKapal\n")
    qb.WriteString("    GROUP BY K.KdKapal, K.Nama, NMK.Tanggal\n")
    qb.WriteString(") TGLPENGGUNAAN GROUP BY KdKapal, Nama")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
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
    var qb strings.Builder

    qb.WriteString("SELECT NMK.Tanggal, K.KdKapal, K.Nama, K.Kapasitas, COUNT(*) JmlAwak\n")
    qb.WriteString("FROM KAPAL K\n")
    qb.WriteString("JOIN NELAYAN_MENGGUNAKAN_KAPAL NMK ON NMK.KdKapal = K.KdKapal\n")
    qb.WriteString("GROUP BY K.KdKapal, K.Nama, K.Kapasitas, NMK.Tanggal HAVING COUNT(*) > K.Kapasitas")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
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
    var qb strings.Builder

    fmt.Printf("Masukkan tanggal (YYYY-MM-DD) : ")
    tgl := input(lineInput)

    qb.WriteString("SELECT NMK.Tanggal, K.KdKapal, K.Nama Kapal, N.KdNelayan, N.Nama Nelayan\n")
    qb.WriteString("FROM NELAYAN_MENGGUNAKAN_KAPAL NMK\n")
    qb.WriteString("JOIN NELAYAN N ON N.KdNelayan = NMK.KdNelayan\n")
    qb.WriteString("JOIN KAPAL K ON K.KdKapal = NMK.KdKapal\n")
    qb.WriteString("WHERE NMK.Tanggal = @p1")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String(), tgl)
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
    var qb strings.Builder

    qb.WriteString("SELECT KdKapal, Nama FROM KAPAL");
    data, err = db.Query(qb.String())
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

    qb.Reset()
    qb.WriteString("SELECT K.Nama Kapal, NMK.Tanggal, N.KdNelayan, N.Nama Nelayan\n")
    qb.WriteString("FROM NELAYAN_MENGGUNAKAN_KAPAL NMK\n")
    qb.WriteString("JOIN NELAYAN N ON N.KdNelayan = NMK.KdNelayan\n")
    qb.WriteString("JOIN KAPAL K ON K.KdKapal = NMK.KdKapal\n")
    qb.WriteString("WHERE NMK.KdKapal = @p1 AND NMK.Tanggal = @p2")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String(), kdk, tgl)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

