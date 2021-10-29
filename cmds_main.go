package main

import (
    "database/sql"
    "fmt"
    "sort"
)

func cmd_help() error {
    sorted := make([]string, len(cmds))
    i := 0
    for cmd := range cmds {
        sorted[i] = cmd
        i++
    }
    sort.Strings(sorted)
    for i := range sorted {
        fmt.Printf(" %s - %s\n", sorted[i], cmds[sorted[i]].desc)
    }

    return nil
}

func cmd_query() error {
    var err error

    fmt.Printf("Masukkan SQL query di bawah kemudian akhiri dengan EOF:\n")

    fmt.Printf("=== BEGIN SQL STATEMENT ===\n");
    qstr := input(textInput)
    fmt.Printf("=== END SQL STATEMENT ===\n")

    fmt.Printf("Executing ...\n")

    var data *sql.Rows
    data, err = db.Query(qstr)
    if err != nil {
        return err
    }
    displaySQLResult(data)
    data.Close()

    return err
}

func cmd_exit() error {
    running = false

    return nil
}

