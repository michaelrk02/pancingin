package main

import (
    "database/sql"
    "strings"
)

func cmd_stats_rankSupplier() error {
    var err error 
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT S.KdSupplier, S.Nama, ISNULL(SUM(I.HargaJual * SMP.Stok), 0) Nilai\n")
    qb.WriteString("FROM SUPPLIER S\n")
    qb.WriteString("LEFT JOIN SUPPLIER_MENSUPLAI_PRODUK SMP ON SMP.KdSupplier = S.KdSupplier\n")
    qb.WriteString("LEFT JOIN PRODUK PR ON PR.KdProduk = SMP.KdProduk\n")
    qb.WriteString("LEFT JOIN IKAN I ON I.KdIkan = PR.KdIkan\n")
    qb.WriteString("GROUP BY S.KdSupplier, S.Nama\n")
    qb.WriteString("ORDER BY Nilai DESC")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_stats_rankPabrik() error {
    var err error 
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT P.KdPabrik, P.Nama, ISNULL(SUM(I.HargaJual * PMP.Stok), 0) Nilai\n")
    qb.WriteString("FROM PABRIK P\n")
    qb.WriteString("LEFT JOIN PABRIK_MENGHASILKAN_PRODUK PMP ON PMP.KdPabrik = P.KdPabrik\n")
    qb.WriteString("LEFT JOIN PRODUK PR ON PR.KdProduk = PMP.KdProduk\n")
    qb.WriteString("LEFT JOIN IKAN I ON I.KdIkan = PR.KdIkan\n")
    qb.WriteString("GROUP BY P.KdPabrik, P.Nama\n")
    qb.WriteString("ORDER BY Nilai DESC")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_stats_rankPetani() error {
    var err error 
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT P.KdPetani, P.Nama, ISNULL(SUM(I.HargaJual * PMI.StokDewasa), 0) Nilai\n")
    qb.WriteString("FROM PETANI P\n")
    qb.WriteString("LEFT JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdPetani = P.KdPetani\n")
    qb.WriteString("LEFT JOIN IKAN I ON I.KdIkan = PMI.KdIkan\n")
    qb.WriteString("GROUP BY P.KdPetani, P.Nama\n")
    qb.WriteString("ORDER BY Nilai DESC")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_stats_rankNelayan() error {
    var err error 
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT N.KdNelayan, N.Nama, ISNULL(SUM(I.HargaJual * NMI.Persediaan), 0) Nilai\n")
    qb.WriteString("FROM NELAYAN N\n")
    qb.WriteString("LEFT JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdNelayan = N.KdNelayan\n")
    qb.WriteString("LEFT JOIN IKAN I ON I.KdIkan = NMI.KdIkan\n")
    qb.WriteString("GROUP BY N.KdNelayan, N.Nama\n")
    qb.WriteString("ORDER BY Nilai DESC")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

