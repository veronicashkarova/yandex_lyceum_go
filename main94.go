package main

func CountingSort(contacts []string) map[string]int {
    dom := make(map [string] int )

    for _, contacts := range contacts {
        dom[contacts]++
    }

    return dom
}