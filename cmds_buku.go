package main

import (
    "database/sql"
)

func cmd_buku_belanjaPelanggan() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_BUKU_BelanjaPelanggan")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_buku_belanjaSupplier() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_BUKU_BelanjaSupplier")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_buku_belanjaPabrik() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_BUKU_BelanjaPabrik")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_buku_belanjaPetani() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_BUKU_BelanjaPetani")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_buku_pendapatanNelayan() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_BUKU_PendapatanNelayan")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_buku_pendapatanPetani() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_BUKU_PendapatanPetani")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_buku_pendapatanPabrik() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_BUKU_PendapatanPabrik")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_buku_pendapatanSupplier() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_BUKU_PendapatanSupplier")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

