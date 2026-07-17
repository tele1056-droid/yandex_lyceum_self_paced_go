package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
HTTP-сервер уже использует горутины,
Каждый HTTP-запрос в Go обрабатывается в отдельной горутине:
http.HandleFunc("/", handler) // handler запускается в новой горутине на каждый запрос
Это значит, что несколько запросов могут выполняться одновременно.

Мьютекс нужен всегда, когда несколько горутин читают/пишут общие данные.

Общая переменная	Нужен мьютекс
*/

/*
			НАПРИМЕР
Только критическая секция (между Lock и Unlock) выполняется последовательно
Всё, что вне мьютекса, выполняется параллельно

Всё, что вне мьютекса, выполняется параллельно
func handler(w http.ResponseWriter, r *http.Request) {
    //  Эта часть выполняется параллельно (все запросы)
    name := r.URL.Query().Get("name")

    mu.Lock()
    //  Эта часть выполняется последовательно (по очереди)
    counter++
    count := counter
    mu.Unlock()

    //  Эта часть снова параллельно
    fmt.Fprintf(w, "%s: %d", name, count)
}
*/

//тут по заданию нужно добавить два хендлера, запросы по адресу которых / и /metrics будут возвращать занчения, и для хендлера /metrics подсчет числа запросов реализовать через middleware

var (
    countFib int
    countReq int
    mu sync.Mutex
)

//          ВСПОМОГАТ. ФУНКЦИЯ (FIB)
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


//          MIDDLEWARE
//чтобы реализовать подсчет числа запросов через MIDDLEWARE, мы тут считаем count, и в main обвернем fibHandler в этот мидлваре, т.е. при запросе на хендлер fibHandler, сначала запрос пройдет Metrics, сделает подсчет, и выполнит обработчик fibHandler. А при запросе на /metrick этот хедлер будет отправлять на клиент коллич. этих вызовов fibHandler
func Metrics(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        //и нужно использовать мьютекс, т.к. запросы - это горутины. Т.е. countReq - это общая переменная, к которой могут обращаться несколько горутин одновременно

        //и правило: каждый раз когда читаешь или пишешь общую переменную, например как тут countReq, нужно использовать мьютекс

        //конкретно тут мы записываем
        mu.Lock()
        countReq++
        mu.Unlock()

        next.ServeHTTP(w, r)
    })
}

//          HANDLERS
func fibHandler(w http.ResponseWriter, r *http.Request) {
    // Увеличиваем счётчик при каждом запросе, и тут сначала подсчитываем фибоначи, а потом увелич счетчик
	result := fib(countFib)
    countFib++
    
    fmt.Fprint(w, result)
}

func metrickHandler(w http.ResponseWriter, r *http.Request) {
	//countReq := +count

    //а тут читаем из общ. перемен. countReq,и тоже нужно читать с мьютексом
    mu.Lock()
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", countReq)
    mu.Unlock()
}

// ВАЖНО: Всегда используй один и тот же мьютекс для защиты одной переменной!!!
/*
То есть вот так нельзя:
var (
    counter int
    mu1     sync.Mutex
    mu2     sync.Mutex
)

// ❌ Запись через mu1
mu1.Lock()
counter++
mu1.Unlock()

// ❌ Чтение через mu2
mu2.Lock()
fmt.Println(counter) // ❌ не защищено!
mu2.Unlock()

Не работало бы, потому что mu1 и mu2 — разные блокировки
*/

func main () {
    mux := http.NewServeMux()

    handlerFib := Metrics(http.HandlerFunc(fibHandler))

    mux.Handle("/", handlerFib)
    mux.Handle("/metrics", http.HandlerFunc(metrickHandler))

    fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", mux)
}