package main

import (
	"fmt"
	"net/http"
)

//также делаем нахождение фибоначи как и в GO2_Step7_01, но в глобал. count храним колич вызовов
var count int = 0

func fib(n int) int {
    if n <= 1 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}

func handler(w http.ResponseWriter, r *http.Request) {
    // Увеличиваем счётчик при каждом запросе, и тут сначала подсчитываем фибоначи, а потом увелич счетчик
	result := fib(count)
    count++
    
    fmt.Fprint(w, result)
}

func main() {
    http.HandleFunc("/", handler)
    //fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}