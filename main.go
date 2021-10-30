package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/url"
    "os"

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

    var cfg DBConfig
    func() {
        var f *os.File
        f, err = os.Open("config.json")
        if err != nil {
            panic(err)
        }
        defer f.Close()

        dec := json.NewDecoder(f)
        err = dec.Decode(&cfg)
        if err != nil {
            panic(err)
        }
    }()

    db, err = sql.Open("sqlserver", fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=15", url.QueryEscape(cfg.DBUser), url.QueryEscape(cfg.DBPass), cfg.DBHost, url.QueryEscape(cfg.DBName)))
    if err != nil {
        panic(err)
    }
    defer db.Close()

    cmds = make(map[string]operation)

    // main commands
    cmds["help"] = operation{desc: "tampilkan daftar perintah", handler: cmd_help}
    cmds["query"] = operation{desc: "jalankan perintah SQL", handler: cmd_query}
    cmds["tables"] = operation{desc: "lihat daftar table", handler: cmd_tables}
    cmds["views"] = operation{desc: "lihat daftar view", handler: cmd_views}
    cmds["exit"] = operation{desc: "keluar dari program", handler: cmd_exit}

    // "list" commands
    cmds["list:nelayan"] = operation{desc: "tampilkan daftar nelayan", handler: cmd_list_nelayan}
    cmds["list:kapal"] = operation{desc: "tampilkan daftar kapal beserta lokasinya di laut", handler: cmd_list_kapal}
    cmds["list:laut"] = operation{desc: "tampilkan daftar laut", handler: cmd_list_laut}
    cmds["list:ikan"] = operation{desc: "tampilkan daftar ikan", handler: cmd_list_ikan}
    cmds["list:petani"] = operation{desc: "tampilkan daftar petani tambak", handler: cmd_list_petani}
    cmds["list:pabrik"] = operation{desc: "tampilkan daftar pabrik", handler: cmd_list_pabrik}
    cmds["list:produk"] = operation{desc: "tampilkan daftar produk beserta ikan asalnya", handler: cmd_list_produk}
    cmds["list:supplier"] = operation{desc: "tampilkan daftar supplier", handler: cmd_list_supplier}
    cmds["list:pelanggan"] = operation{desc: "tampilkan daftar pelanggan", handler: cmd_list_pelanggan}

    // "ikan" commands
    cmds["ikan:dibudidayakan"] = operation{desc: "tampilkan daftar ikan yang dibudidayakan", handler: cmd_ikan_dibudidayakan}
    cmds["ikan:di_laut"] = operation{desc: "tampilkan sebaran spesies ikan di laut", handler: cmd_ikan_diLaut}
    cmds["ikan:jumlah_per_laut"] = operation{desc: "tampilkan sebaran jumlah spesies ikan per laut", handler: cmd_ikan_jumlahPerLaut}

    // "kapal" commands
    cmds["kapal:penggunaan"] = operation{desc: "tampilkan jumlah penggunaan kapal", handler: cmd_kapal_penggunaan}
    cmds["kapal:overkapasitas"] = operation{desc: "tampilkan pada tanggal berapa suatu kapal over kapasitas", handler: cmd_kapal_overkapasitas}
    cmds["kapal:penggunaan_tanggal"] = operation{desc: "tampilkan daftar nelayan yang berlayar pada tanggal tertentu", handler: cmd_kapal_penggunaanTanggal}
    cmds["kapal:digunakan_tanggal"] = operation{desc: "tampilkan daftar nelayan yang menggunakan kapal tertentu pada tanggal tertentu", handler: cmd_kapal_digunakanTanggal}

    // "terdekat" commands
    cmds["terdekat:supplier"] = operation{desc: "tampilkan supplier terdekat dengan pelanggan", handler: cmd_terdekat_supplier}
    cmds["terdekat:pabrik"] = operation{desc: "tampilkan pabrik terdekat dengan supplier", handler: cmd_terdekat_pabrik}
    cmds["terdekat:tambak"] = operation{desc: "tampilkan tambak terdekat dengan pabrik", handler: cmd_terdekat_tambak}

    // "jelajah" commands
    cmds["jelajah:ikan_tangkapan"] = operation{desc: "tampilkan daftar ikan beserta harga jualnya yang ditangkap nelayan", handler: cmd_jelajah_ikanTangkapan}
    cmds["jelajah:ikan_tambak"] = operation{desc: "tampilkan daftar ikan beserta harga jualnya yang dibudidayakan petani tambak", handler: cmd_jelajah_ikanTambak}
    cmds["jelajah:produk_pabrik"] = operation{desc: "tampilkan daftar produk beserta harga jualnya yang dihasilkan pabrik", handler: cmd_jelajah_produkPabrik}
    cmds["jelajah:produk_supplier"] = operation{desc: "tampilkan daftar produk beserta harga jualnya yang disediakan supplier", handler: cmd_jelajah_produkSupplier}
    cmds["jelajah:produk_supplier_termurah"] = operation{desc: "tampilkan harga (termurah sampai termahal) dan stok produk dari supplier yang dapat dibeli pelanggan", handler: cmd_jelajah_produkSupplierTermurah}
    cmds["jelajah:produk_pabrik_termurah"] = operation{desc: "tampilkan harga (termurah sampai termahal) dan stok produk dari pabrik yang dapat dibeli supplier", handler: cmd_jelajah_produkPabrikTermurah}
    cmds["jelajah:ikan_tambak_termurah"] = operation{desc: "tampilkan harga (termurah sampai termahal) dan stok dewasa ikan dari tambak yang dapat dibeli pabrik", handler: cmd_jelajah_ikanTambakTermurah}
    cmds["jelajah:ikan_tangkapan_termurah"] = operation{desc: "tampilkan harga (termurah sampai termahal) dan persediaan ikan dari nelayan yang dapat dibeli pabrik dan petani tambak", handler: cmd_jelajah_ikanTangkapanTermurah}

    // "stats" commands
    cmds["stats:rank_supplier"] = operation{desc: "tampilkan ranking supplier berdasarkan nilai seluruh stok produk yang disediakan", handler: cmd_stats_rankSupplier}
    cmds["stats:rank_pabrik"] = operation{desc: "tampilkan ranking pabrik berdasarkan nilai seluruh stok produk yang dihasilkan", handler: cmd_stats_rankPabrik}
    cmds["stats:rank_petani"] = operation{desc: "tampilkan ranking petani tambak berdasarkan nilai seluruh stok dewasa ikan", handler: cmd_stats_rankPetani}
    cmds["stats:rank_nelayan"] = operation{desc: "tampilkan ranking nelayan berdasarkan nilai seluruh persediaan tangkapan ikan", handler: cmd_stats_rankNelayan}

    // "buku" commands
    cmds["buku:belanja_pelanggan"] = operation{desc: "tampilkan total belanja produk dari supplier oleh pelanggan", handler: cmd_buku_belanjaPelanggan}
    cmds["buku:belanja_supplier"] = operation{desc: "tampilkan total belanja produk dari pabrik oleh supplier", handler: cmd_buku_belanjaSupplier}
    cmds["buku:belanja_pabrik"] = operation{desc: "tampilkan total belanja ikan dari nelayan oleh pabrik", handler: cmd_buku_belanjaPabrik}
    cmds["buku:belanja_petani"] = operation{desc: "tampilkan total belanja ikan dari nelayan oleh petani tambak", handler: cmd_buku_belanjaPetani}
    cmds["buku:pendapatan_nelayan"] = operation{desc: "tampilkan total pendapatan nelayan dalam penjualan ikan ke pabrik", handler: cmd_buku_pendapatanNelayan}
    cmds["buku:pendapatan_petani"] = operation{desc: "tampilkan total pendapatan petani tambak dalam penjualan ikan ke pabrik", handler: cmd_buku_pendapatanPetani}
    cmds["buku:pendapatan_pabrik"] = operation{desc: "tampilkan total pendapatan pabrik dalam penjualan produk ke supplier", handler: cmd_buku_pendapatanPabrik}
    cmds["buku:pendapatan_supplier"] = operation{desc: "tampilkan total pendapatan supplier dalam penjualan produk ke pelanggan", handler: cmd_buku_pendapatanSupplier}

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
