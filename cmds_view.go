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

func cmd_pelanggan_supplierTerdekat() error {
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

func cmd_pelanggan_telusurProduk() error {
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

func cmd_supplier_pabrikTerdekat() error {
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

func cmd_supplier_telusurProduk() error {
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

func cmd_supplier_ranking() error {
    var err error 
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT S.KdSupplier, S.Nama, SUM(I.HargaJual * SMP.Stok) Nilai\n")
    qb.WriteString("FROM SUPPLIER S\n")
    qb.WriteString("JOIN SUPPLIER_MENSUPLAI_PRODUK SMP ON SMP.KdSupplier = S.KdSupplier\n")
    qb.WriteString("JOIN PRODUK PR ON PR.KdProduk = SMP.KdProduk\n")
    qb.WriteString("JOIN IKAN I ON I.KdIkan = PR.KdIkan\n")
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

func cmd_pabrik_tambakTerdekat() error {
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

func cmd_pabrik_telusurTambak() error {
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

func cmd_pabrik_ranking() error {
    var err error 
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT P.KdPabrik, P.Nama, SUM(I.HargaJual * PMP.Stok) Nilai\n")
    qb.WriteString("FROM PABRIK P\n")
    qb.WriteString("JOIN PABRIK_MENGHASILKAN_PRODUK PMP ON PMP.KdPabrik = P.KdPabrik\n")
    qb.WriteString("JOIN PRODUK PR ON PR.KdProduk = PMP.KdProduk\n")
    qb.WriteString("JOIN IKAN I ON I.KdIkan = PR.KdIkan\n")
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

func cmd_petani_ranking() error {
    var err error 
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT P.KdPetani, P.Nama, SUM(I.HargaJual * PMI.StokDewasa) Nilai\n")
    qb.WriteString("FROM PETANI P\n")
    qb.WriteString("JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdPetani = P.KdPetani\n")
    qb.WriteString("JOIN IKAN I ON I.KdIkan = PMI.KdIkan\n")
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

func cmd_pabrikPetani_telusurIkan() error {
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

func cmd_nelayan_ranking() error {
    var err error 
    var data *sql.Rows
    var qb strings.Builder

    qb.WriteString("SELECT N.KdNelayan, N.Nama, SUM(I.HargaJual * NMI.Persediaan) Nilai\n")
    qb.WriteString("FROM NELAYAN N\n")
    qb.WriteString("JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdNelayan = N.KdNelayan\n")
    qb.WriteString("JOIN IKAN I ON I.KdIkan = NMI.KdIkan\n")
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

