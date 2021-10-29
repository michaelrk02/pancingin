package main

import (
    "database/sql"
    "fmt"
    "net/url"

    _ "github.com/denisenkom/go-mssqldb"
)

var running = true

var cmds map[string]operation

type operation struct {
    desc string
    handler func () error
}

func main() {
    var err error
    var ok bool

    db, err = sql.Open("sqlserver", fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=15", url.QueryEscape(dbuser), url.QueryEscape(dbpass), dbhost, url.QueryEscape(dbname)))
    if err != nil {
        panic(err)
    }
    defer db.Close()

    cmds = make(map[string]operation)

    // main commands
    cmds["help"] = operation{desc: "tampilkan daftar perintah", handler: cmd_help}
    cmds["query"] = operation{desc: "jalankan perintah SQL", handler: cmd_query}
    cmds["exit"] = operation{desc: "keluar dari program", handler: cmd_exit}

    // view commands
    cmds["kapal:penggunaan"] = operation{desc: "tampilkan jumlah penggunaan kapal", handler: cmd_kapal_penggunaan}
    cmds["kapal:overkapasitas"] = operation{desc: "tampilkan pada tanggal berapa suatu kapal over kapasitas", handler: cmd_kapal_overkapasitas}
    cmds["pelanggan:supplier_terdekat"] = operation{desc: "tampilkan supplier terdekat dengan pelanggan", handler: cmd_pelanggan_supplierTerdekat}
    cmds["pelanggan:telusur_produk"] = operation{desc: "tampilkan harga (termurah sampai termahal) dan stok produk dari supplier yang dapat dibeli pelanggan", handler: cmd_pelanggan_telusurProduk}
    cmds["supplier:pabrik_terdekat"] = operation{desc: "tampilkan supplier terdekat dengan pelanggan", handler: cmd_supplier_pabrikTerdekat}
    cmds["supplier:telusur_produk"] = operation{desc: "tampilkan harga (termurah sampai termahal) dan stok produk dari pabrik yang dapat dibeli supplier", handler: cmd_supplier_telusurProduk}
    cmds["supplier:ranking"] = operation{desc: "tampilkan ranking supplier berdasarkan nilai seluruh stok produk yang disediakan", handler: cmd_supplier_ranking}
    cmds["pabrik:tambak_terdekat"] = operation{desc: "tampilkan tambak terdekat dengan pabrik", handler: cmd_pabrik_tambakTerdekat}
    cmds["pabrik:telusur_tambak"] = operation{desc: "tampilkan harga (termurah sampai termahal) dan stok dewasa ikan dari tambak yang dapat dibeli pabrik", handler: cmd_pabrik_telusurTambak}
    cmds["pabrik:ranking"] = operation{desc: "tampilkan ranking pabrik berdasarkan nilai seluruh stok produk yang dihasilkan", handler: cmd_pabrik_ranking}
    cmds["petani:ranking"] = operation{desc: "tampilkan ranking petani tambak berdasarkan nilai seluruh stok dewasa ikan", handler: cmd_petani_ranking}
    cmds["pabrik_petani:telusur_ikan"] = operation{desc: "tampilkan harga (termurah sampai termahal) dan persediaan ikan dari nelayan yang dapat dibeli pabrik dan petani tambak", handler: cmd_pabrikPetani_telusurIkan}
    cmds["ikan:dibudidayakan"] = operation{desc: "tampilkan daftar ikan yang dibudidayakan", handler: cmd_ikan_dibudidayakan}
    cmds["nelayan:ranking"] = operation{desc: "tampilkan ranking nelayan berdasarkan nilai seluruh persediaan tangkapan ikan", handler: cmd_nelayan_ranking}

    fmt.Printf("Selamat datang di interface PancingIN v1.0. Ketik `help` untuk melihat bantuan\n");
    for running {
        fmt.Printf("$ ")
        cmd := input(lineInput)

        var op operation
        if op, ok = cmds[cmd]; !ok {
            fmt.Printf("Perintah tidak ditemukan :(\n")
            continue
        }

        fmt.Printf("Menjalankan : %s\n", op.desc)

        err = op.handler()
        if err != nil {
            fmt.Printf("ERROR: %s\n", err)
            continue
        }
    }
}
