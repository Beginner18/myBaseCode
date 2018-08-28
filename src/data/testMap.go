package main

import (
    "fmt"
    //"sort"
)

func main() {
   /* var m = map[string]int{
        "unix":         0,
        "python":       1,
        "go":           2,
        "javascript":   3,
        "testing":      4,
        "philosophy":   5,
        "startups":     6,
        "productivity": 7,
        "hn":           8,
        "reddit":       9,
        "C++":          10,
    }
    var keys []string
    //put all key of map m in []string keys
    for k := range m {
        keys = append(keys, k)
    }

    stringnum := sort.SearchStrings(keys,"python")
    fmt.Println( stringnum)
    fmt.Println(keys[stringnum])
    //sort string keys
    sort.Strings(keys)
    //loop all element of keys as keyvalue for map m
    for _, k := range keys {
        fmt.Println("Key:", k, "Value:", m[k])
    }*/
    try1Slice := []string{"11","222","333"}
    fmt.Println(try1Slice)
    try1Slice = try1Slice[1:]
    fmt.Println(try1Slice)


}
