package main

import (
    "database/sql"
)

func cmd_stats_rankSupplier() error {
    var err error 
    var data *sql.Rows

    data, err = db.Query("SELECT * FROM V_STATS_RankSupplier ORDER BY Nilai DESC")
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

    data, err = db.Query("SELECT * FROM V_STATS_RankPabrik ORDER BY Nilai DESC")
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

    data, err = db.Query("SELECT * FROM V_STATS_RankPetani ORDER BY Nilai DESC")
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

    data, err = db.Query("SELECT * FROM V_STATS_RankNelayan ORDER BY Nilai DESC")
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

