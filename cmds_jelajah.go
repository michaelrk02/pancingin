package main

import (
    "database/sql"
    "fmt"
    "strings"
)

func cmd_jelajah_ikanTangkapan() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdNelayan, Nama FROM NELAYAN")
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdn int
    fmt.Printf("Masukkan KdNelayan : ")
    fmt.Sscanf(input(lineInput), "%d", &kdn)

    qb.Reset()
    qb.WriteString("SELECT N.Nama Nelayan, I.KdIkan, I.Nama, (I.HargaJual + NMI.Keuntungan) Harga, NMI.Persediaan\n")
    qb.WriteString("FROM IKAN I\n")
    qb.WriteString("JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdIkan = I.KdIkan\n")
    qb.WriteString("JOIN NELAYAN N ON N.KdNelayan = NMI.KdNelayan\n")
    qb.WriteString("WHERE N.KdNelayan = @p1")
    data, err = db.Query(qb.String(), kdn)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_jelajah_ikanTambak() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdPetani, Nama FROM PETANI")
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdPetani : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    qb.Reset()
    qb.WriteString("SELECT P.Nama Petani, I.KdIkan, I.Nama, (I.HargaJual + PMI.BiayaPerSatuan + PMI.Keuntungan) Harga, PMI.StokDewasa\n")
    qb.WriteString("FROM IKAN I\n")
    qb.WriteString("JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdIkan = I.KdIkan\n")
    qb.WriteString("JOIN PETANI P ON P.KdPetani = PMI.KdPetani\n")
    qb.WriteString("WHERE P.KdPetani = @p1")
    data, err = db.Query(qb.String(), kdp)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_jelajah_produkPabrik() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdPabrik, Nama FROM PABRIK")
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdPabrik : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    qb.Reset()
    qb.WriteString("SELECT PA.Nama Pabrik, P.KdProduk, P.Nama, (I.HargaJual + PMP.BiayaProduksi + PMP.Keuntungan) Harga, PMP.Stok\n")
    qb.WriteString("FROM PRODUK P\n")
    qb.WriteString("JOIN PABRIK_MENGHASILKAN_PRODUK PMP ON PMP.KdProduk = P.KdProduk\n")
    qb.WriteString("JOIN PABRIK PA ON PA.KdPabrik = PMP.KdPabrik\n")
    qb.WriteString("JOIN IKAN I ON I.KdIkan = P.KdIkan\n")
    qb.WriteString("WHERE PA.KdPabrik = @p1")
    data, err = db.Query(qb.String(), kdp)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_jelajah_produkSupplier() error {
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
    qb.WriteString("SELECT S.Nama Supplier, P.KdProduk, P.Nama, (I.HargaJual + SMP.BiayaOperasional + SMP.Keuntungan) Harga, SMP.Stok\n")
    qb.WriteString("FROM PRODUK P\n")
    qb.WriteString("JOIN SUPPLIER_MENSUPLAI_PRODUK SMP ON SMP.KdProduk = P.KdProduk\n")
    qb.WriteString("JOIN SUPPLIER S ON S.KdSupplier = SMP.KdSupplier\n")
    qb.WriteString("JOIN IKAN I ON I.KdIkan = P.KdIkan\n")
    qb.WriteString("WHERE S.KdSupplier = @p1")
    data, err = db.Query(qb.String(), kds)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_jelajah_produkSupplierTermurah() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdProduk, Nama FROM PRODUK")
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdProduk : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    qb.Reset()
    qb.WriteString("SELECT P.Nama, S.KdSupplier, S.Nama NamaSupplier, (I.HargaJual + SMP.BiayaOperasional + SMP.Keuntungan) Harga, SMP.Stok\n")
    qb.WriteString("FROM PRODUK P\n")
    qb.WriteString("JOIN IKAN I ON I.KdIkan = P.KdIkan\n")
    qb.WriteString("JOIN SUPPLIER_MENSUPLAI_PRODUK SMP ON SMP.KdProduk = P.KdProduk\n")
    qb.WriteString("JOIN SUPPLIER S ON S.KdSupplier = SMP.KdSupplier\n")
    qb.WriteString("WHERE P.KdProduk = @p1\n")
    qb.WriteString("ORDER BY Harga ASC")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String(), kdp)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_jelajah_produkPabrikTermurah() error {
    var err error
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT KdProduk, Nama FROM PRODUK")
    data, err = db.Query(qb.String())
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdProduk : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    qb.Reset()
    qb.WriteString("SELECT P.Nama, PA.KdPabrik, PA.Nama NamaPabrik, (I.HargaJual + PMP.BiayaProduksi + PMP.Keuntungan) Harga, PMP.Stok\n")
    qb.WriteString("FROM PRODUK P\n")
    qb.WriteString("JOIN IKAN I ON I.KdIkan = P.KdIkan\n")
    qb.WriteString("JOIN PABRIK_MENGHASILKAN_PRODUK PMP ON PMP.KdProduk = P.KdProduk\n")
    qb.WriteString("JOIN PABRIK PA ON PA.KdPabrik = PMP.KdPabrik\n")
    qb.WriteString("WHERE P.KdProduk = @p1\n")
    qb.WriteString("ORDER BY Harga ASC")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String(), kdp)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_jelajah_ikanTambakTermurah() error {
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
    qb.WriteString("SELECT I.Nama, P.KdPetani, P.Nama NamaPetani, (I.HargaJual + PMI.BiayaPerSatuan + PMI.Keuntungan) Harga, PMI.StokDewasa\n")
    qb.WriteString("FROM IKAN I\n")
    qb.WriteString("JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdIkan = I.KdIkan\n")
    qb.WriteString("JOIN PETANI P ON P.KdPetani = PMI.KdPetani\n")
    qb.WriteString("WHERE I.KdIkan = @p1\n")
    qb.WriteString("ORDER BY Harga ASC")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String(), kdi)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_jelajah_ikanTangkapanTermurah() error {
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
    qb.WriteString("SELECT I.Nama, N.KdNelayan, N.Nama NamaNelayan, (I.HargaJual + NMI.Keuntungan) Harga, NMI.Persediaan\n")
    qb.WriteString("FROM IKAN I\n")
    qb.WriteString("JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdIkan = I.KdIkan\n")
    qb.WriteString("JOIN NELAYAN N ON N.KdNelayan = NMI.KdNelayan\n")
    qb.WriteString("WHERE I.KdIkan = @p1\n")
    qb.WriteString("ORDER BY Harga ASC")
    reportSQLStatement(qb.String())
    data, err = db.Query(qb.String(), kdi)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

