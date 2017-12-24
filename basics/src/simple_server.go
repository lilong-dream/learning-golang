// http://localhost:8080/?id=1

package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Author struct{
    Id      string
    Name    string
}

var authors = map[string]string{
    "1": "二师兄",
    "2": "罗总",
    "3": "昭哥",
    "4": "小温",
}

func handler(w http.ResponseWriter, r *http.Request) {
    ids, ok := r.URL.Query()["id"]

    if !ok || len(ids) < 1 {
        fmt.Println("url param 'id' is missing")
        return
    }

    id := ids[0]

    result := Author{Id: id, Name: "not found"}

    if author, ok := authors[id]; ok {
        result.Name = author
    }

    json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
