package main

import (
    "database/sql"
    "fmt"
)

func cmd_jelajah_ikanTangkapan() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT KdNelayan, Nama FROM NELAYAN")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdn int
    fmt.Printf("Masukkan KdNelayan : ")
    fmt.Sscanf(input(lineInput), "%d", &kdn)

    data, err = db.Query("SELECT * FROM V_JELAJAH_IkanTangkapan WHERE KdNelayan = @p1", kdn)
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

    data, err = db.Query("SELECT KdPetani, Nama FROM PETANI")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdPetani : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    data, err = db.Query("SELECT * FROM V_JELAJAH_IkanTambak WHERE KdPetani = @p1", kdp)
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

    data, err = db.Query("SELECT KdPabrik, Nama FROM PABRIK")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdPabrik : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    data, err = db.Query("SELECT * FROM V_JELAJAH_ProdukPabrik WHERE KdPabrik = @p1", kdp)
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

    data, err = db.Query("SELECT KdSupplier, Nama FROM SUPPLIER")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kds int
    fmt.Printf("Masukkan KdSupplier : ")
    fmt.Sscanf(input(lineInput), "%d", &kds)

    data, err = db.Query("SELECT * FROM V_JELAJAH_ProdukSupplier WHERE KdSupplier = @p1", kds)
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

    data, err = db.Query("SELECT KdProduk, Nama FROM PRODUK")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdProduk : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    data, err = db.Query("SELECT * FROM V_JELAJAH_ProdukSupplierTermurah WHERE KdProduk = @p1 ORDER BY Harga ASC", kdp)
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

    data, err = db.Query("SELECT KdProduk, Nama FROM PRODUK")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdp int
    fmt.Printf("Masukkan KdProduk : ")
    fmt.Sscanf(input(lineInput), "%d", &kdp)

    data, err = db.Query("SELECT * FROM V_JELAJAH_ProdukPabrikTermurah WHERE KdProduk = @p1 ORDER BY Harga ASC", kdp)
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

    data, err = db.Query("SELECT KdIkan, Nama FROM IKAN")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdi int
    fmt.Printf("Masukkan KdIkan : ")
    fmt.Sscanf(input(lineInput), "%d", &kdi)

    data, err = db.Query("SELECT * FROM V_JELAJAH_IkanTambakTermurah WHERE KdIkan = @p1 ORDER BY Harga ASC", kdi)
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

    data, err = db.Query("SELECT KdIkan, Nama FROM IKAN")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    var kdi int
    fmt.Printf("Masukkan KdIkan : ")
    fmt.Sscanf(input(lineInput), "%d", &kdi)

    data, err = db.Query("SELECT * FROM V_JELAJAH_IkanTangkapanTermurah WHERE KdIkan = @p1 ORDER BY Harga ASC", kdi)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

