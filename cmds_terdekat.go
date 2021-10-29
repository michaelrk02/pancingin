package main

import (
    "database/sql"
    "fmt"
    "strings"
)

func cmd_terdekat_supplier() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdPelanggan, Nama FROM PELANGGAN")
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdPelanggan : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    qb.Reset()
    qb.WriteString("SELECT P.Nama NamaPelanggan, S.Lokasi, S.Nama NamaSupplier\n")
    qb.WriteString("FROM SUPPLIER S\n")
    qb.WriteString("JOIN PELANGGAN P ON P.Lokasi = S.Lokasi\n")
    qb.WriteString("WHERE P.KdPelanggan = @p1");
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String(), kdp)
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
    var qb strings.Builder

    qb.WriteString("SELECT KdSupplier, Nama FROM SUPPLIER")
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kds int
    fmt.Printf("Masukkan KdSupplier : ")
    fmt.Sscanf(input(lineInput), "%d", &kds)

    qb.Reset()
    qb.WriteString("SELECT S.Nama NamaSupplier, P.Lokasi, P.KdPabrik, P.Nama NamaPabrik\n")
    qb.WriteString("FROM PABRIK P\n")
    qb.WriteString("JOIN SUPPLIER S ON S.Lokasi = P.Lokasi\n")
    qb.WriteString("WHERE S.KdSupplier = @p1")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String(), kds)
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
    var qb strings.Builder

    qb.WriteString("SELECT KdSupplier, Nama FROM SUPPLIER")
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kds int
    fmt.Printf("Masukkan KdSupplier : ")
    fmt.Sscanf(input(lineInput), "%d", &kds)

    qb.Reset()
    qb.WriteString("SELECT S.Nama NamaSupplier, P.Lokasi, P.KdPabrik, P.Nama NamaPabrik\n")
    qb.WriteString("FROM PABRIK P\n")
    qb.WriteString("JOIN SUPPLIER S ON S.Lokasi = P.Lokasi\n")
    qb.WriteString("WHERE S.KdSupplier = @p1")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String(), kds)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

