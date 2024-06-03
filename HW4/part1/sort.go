package main

import (
    "fmt"
    "sort"
)

type Item struct {
    ID int
}

func uniqueItems(items []Item) []Item {
    seen := make(map[int]bool)
    var unique []Item
    for _, item := range items {
        if !seen[item.ID] {
            seen[item.ID] = true
            unique = append(unique, item)
        }
    }
    sort.Slice(unique, func(i, j int) bool {
        return unique[i].ID < unique[j].ID
    })
    return unique
}

func main() {
    items := []Item{{ID: 3}, {ID: 2}, {ID: 1}, {ID: 2}}
    unique := uniqueItems(items)
    fmt.Println("Unique Items sorted by ID:")
    for _, item := range unique {
        fmt.Printf("{ID: %d} ", item.ID)
    }
    fmt.Println()
}