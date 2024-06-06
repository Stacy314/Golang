/*Create a structure type that contains one field 
(for example, `ID`). Write a function that takes 
as input a slice with elements of the created type,
and returns a slice of the same type with only 
unique values (structures with duplicate field 
values are discarded). The result of the function 
should be sorted in ascending order of the structure 
field values. Additional conditions: Do not use 
libraries to search for unique values. Use the 
capabilities of the standard `sort` library for 
sorting. */

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