package main

import (
    "database/sql"
)

func cmd_list_nelayan() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_LIST_Nelayan")
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

    data, err = db.Query("SELECT * FROM V_LIST_Kapal")
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

    data, err = db.Query("SELECT * FROM V_LIST_Laut")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_list_ikan() error {
    var err error
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_LIST_Ikan")
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

    data, err = db.Query("SELECT * FROM V_LIST_Petani")
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

    data, err = db.Query("SELECT * FROM V_LIST_Pabrik")
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

    data, err = db.Query("SELECT * FROM V_LIST_Produk")
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

    data, err = db.Query("SELECT * FROM V_LIST_Supplier")
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

    data, err = db.Query("SELECT * FROM V_LIST_Pelanggan")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

