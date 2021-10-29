package main

import (
    "database/sql"
    "fmt"
    "strings"
)

func cmd_ikan_dibudidayakan() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT I.KdIkan, I.Nama, COUNT(*) AS JmlTambak\n")
    qb.WriteString("FROM IKAN I\n")
    qb.WriteString("JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdIkan = I.KdIkan\n")
    qb.WriteString("GROUP BY I.KdIkan, I.Nama")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
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
    var qb strings.Builder

    qb.WriteString("SELECT KdIkan, Nama FROM IKAN")
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdi int
    fmt.Printf("Masukkan KdIkan : ")
    fmt.Sscanf(input(lineInput), "%d", &kdi)

    qb.Reset()
    qb.WriteString("SELECT I.Nama, L.KdLaut, L.Nama\n")
    qb.WriteString("FROM IKAN I\n")
    qb.WriteString("JOIN LAUT_MEMILIKI_IKAN LMI ON LMI.KdIkan = I.KdIkan\n")
    qb.WriteString("JOIN LAUT L ON L.KdLaut = LMI.KdLaut\n")
    qb.WriteString("WHERE I.KdIkan = @p1")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String(), kdi)
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
    var qb strings.Builder

    qb.WriteString("SELECT L.KdLaut, L.Nama, COUNT(*) JmlSpesies\n")
    qb.WriteString("FROM LAUT L\n")
    qb.WriteString("JOIN LAUT_MEMILIKI_IKAN LMI ON LMI.KdLaut = L.KdLaut\n")
    qb.WriteString("GROUP BY L.KdLaut, L.Nama")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

