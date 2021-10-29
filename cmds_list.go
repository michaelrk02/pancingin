package main

import (
    "database/sql"
    "strings"
)

func cmd_list_nelayan() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdNelayan, Nama, NoTelp FROM NELAYAN")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_list_kapal() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT K.KdKapal, K.Nama, K.Jenis, K.Kapasitas, L.Nama Lokasi\n")
    qb.WriteString("FROM KAPAL K\n")
    qb.WriteString("JOIN LAUT L ON L.KdLaut = K.KdLaut")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_list_laut() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdLaut, Nama, Lokasi FROM LAUT")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_list_petani() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdPetani, Nama, NoTelp FROM PETANI")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_list_pabrik() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdPabrik, Nama, Lokasi, NoTelp FROM PABRIK")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_list_produk() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT P.KdProduk, P.Nama, P.Deskripsi, I.Nama AsalIkan\n")
    qb.WriteString("FROM PRODUK P\n")
    qb.WriteString("JOIN IKAN I ON I.KdIkan = P.KdIkan")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_list_supplier() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdSupplier, Nama, Lokasi, NoTelp FROM SUPPLIER")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_list_pelanggan() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdPelanggan, Nama, Lokasi FROM PELANGGAN")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

