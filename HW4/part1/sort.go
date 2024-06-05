package main

import (
    "fmt"
    "sort"
)

type item struct {
    id int
}

func uniqueItems(items []item) []item {
    seen := make(map[int]struct{})
    var unique []item
    for _, itm := range items {
        if _, found := seen[itm.id]; !found {
            seen[itm.id] = struct{}{}
            unique = append(unique, itm)
        }
    }
    sort.Slice(unique, func(i, j int) bool {
        return unique[i].id < unique[j].id
    })
    return unique
}

func main() {
    items := []item{{id: 3}, {id: 2}, {id: 1}, {id: 2}}
    unique := uniqueItems(items)
    fmt.Println("Unique Items sorted by ID:")
    for _, itm := range unique {
        fmt.Printf("{ID: %d} ", itm.id)
    }
    fmt.Println()
}