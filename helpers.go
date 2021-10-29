package main

import (
    "bufio"
    "database/sql"
    "fmt"
    "io"
    "os"
    "strings"
)

type inputType int
const (
    lineInput inputType = iota
    textInput
)

func input(typ inputType) string {
    var err error
    var out strings.Builder

    in := bufio.NewReader(os.Stdin)
    for {
        var c byte

        c, err = in.ReadByte()
        if err != nil {
            if typ == lineInput {
                break
            } else {
                if err == io.EOF {
                    err = nil
                    break
                } else {
                    break
                }
            }
        }

        if typ == lineInput {
            if c == '\n' {
                break
            }
        }
        out.WriteByte(c)
    }

    if err != nil {
        out.Reset()
    }

    return out.String()
}

func reportSQLStatement(s string) {
    fmt.Printf("=== BEGIN SQL STATEMENT ===\n");
    fmt.Printf("%s\n", s)
    fmt.Printf("=== END SQL STATEMENT ===\n")
}

func displaySQLResult(data *sql.Rows) {
    var cols []string
    cols, _ = data.Columns()
    fmt.Printf("[")
    for i := range cols {
        fmt.Printf("%s", cols[i])
        if i < len(cols) - 1 {
            fmt.Printf(", ")
        }
    }
    fmt.Printf("]\n")

    for data.Next() {
        var tuple = make([]string, len(cols))
        var tuplePtr = make([]interface{}, len(cols))
        for i := range tuplePtr {
            tuplePtr[i] = &tuple[i]
        }
        data.Scan(tuplePtr...)

        fmt.Printf("(")
        for i := range tuple {
            var cell = tuple[i]
            cell = strings.ReplaceAll(cell, "\\", "\\\\")
            cell = strings.ReplaceAll(cell, "\"", "\\\"")
            fmt.Printf("\"%s\"", cell)
            if i < len(cols) - 1 {
                fmt.Printf(", ")
            }
        }
        fmt.Printf(")\n")
    }
}

