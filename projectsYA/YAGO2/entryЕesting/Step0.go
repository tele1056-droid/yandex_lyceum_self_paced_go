package main

import (
	"fmt"
	"time"
    "strings"
)

type Ticket struct {
    Ticket string
    User   string
    Status string
    Date   time.Time
}



func GetTasks(text string, user *string, status *string) []Ticket {
    //и вот эту переменную лучше располагать внутри функции
    var tickets []Ticket

    //сплитим текст на слайс из строк, по разделителю перенос строки ("\n")
    lines := strings.Split(text, "\n")
    
    //тут в for range берем каждую строку из lines = ["строка1","строка2" и тд], и сплитим строки по разделителю ("_"), чтобы получить отдельные элементы строки parts = ["TICKET-12345", "Паша Попов", "Готово", "2024-01-01"]. И аппендим в слайс Тикетов каждый parts
    for _, line := range lines {
        //убираем пробелы в начале, в конце строки и пустые строки
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }


        //fmt.Printf("line: %q\n", line)

        //ОБРАБОТКА ОШИБКИ (проверка валидности) на "TICKET", strings.HasPrefix - проверяет начинается ли строка с опред. подстроки, и если есть то вернет true, ести нет то false, и регистр важен. И в случае если возвращ. false (то есть нет подстроки) то не false (!false) - это true и мы заходим в if и завершаем эту итерацию, т.е. код дальше пойдет, а начнет новую итерацию
        if !strings.HasPrefix(line, "TICKET") {
            continue
        }

        //тут сплитим каждую строку
        parts := strings.Split(line, "_")

        //и проверяем колич. элем. в parts
        if len(parts) != 4 {
            continue // пропескаем строку, если не 4 части
        }

        //ОБРАБОТКА ОШИБКИ (проверка валидности даты) тут time.Parse возвращает дату, либо ошибку если формат не совпадает или дата невозможна(напр. 32 января), и мы обрабатываем ошибку err, если она не равна nil то пропускаем 
        timeStamp, err := time.Parse("2006-01-02", parts[3])
        if err != nil {
            continue
        }

        //ОБРАБОТКА ОШИБКИ(проверка валид. статуса) валид = «Готово», «В работе», «Не будет сделано».
        //То есть тут мы обращаем к мапе () по ключу (который лежит в parts[2]), мапа возвращает true если ключ который мы передаем есть в мапе, если нет то false. И используем в if логическое НЕ, т.е. если ключ который мы передаем есть в мапе - вернется true, а !true - это false и код будет выполнятся дальше, если false - то будем заходит в if завершать эту итерацию
        if strings.TrimSpace(parts[2]) == "" {
            continue
        }
        
        validStatuses := map[string]bool {
            "Готово": true,
            "В работе": true,
            "Не будет сделано": true,
        }

        if !validStatuses[parts[2]] {
            continue
        } 
        

        /*
        Можно было так реализовать: 
        status := parts[2]
            if status != "Готово" && status != "В работе" && status != "Не будет сделано" {
                continue
            } */

        //делаем фильтрацию относительно user и status
        if user != nil && parts[1] != *user {
            continue // не добаляем этот тикет, если имена не совпад
        }

        if status != nil && parts[2] != *status {
            continue //не добаляем этот тикет, если статус не совпад
        }

        t := Ticket{
            Ticket: parts[0],
            User: parts[1],
            Status: parts[2],
            Date: timeStamp,
        }

        tickets = append(tickets, t)
    }

    return tickets
}



func main() {
	// text := "TICKET-12345_Паша Попов_Готово_2024-01-01\nTICKET-12346_Иван Иванов_В работе_2024-01-02\nTICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03\nTICKET-12348_Паша Попов_В работе_2024-01-04"

//     chatHistory := `
// TICKET-12345_Паша Попов_Готово_2024-01-01
// TICKET-12346_Иван Иванов_В работе_2024-01-02
// TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
// TICKET-12348_Паша Попов_В работе_2024-01-04
// TICKET-12349_Паша Попов_Готово_2024-01-04
// TICKET-12353_Ещё Один Разраб_В работе_2024-01-08
// `

    chatHistory1 := `
    TICKET-12345_Паша Попов_Готово_2024-01-01
    TICKET-12346_Иван Иванов_В работе_2024-01-02
    TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
    TICKET-12348_Паша Попов_В работе_2024-01-04
    Это не задача
    TICKET-12353_Ещё Один Разраб_В работе_2024-01-08
    TICKET-12363_Иван Иванов_Готово_2024-01-16
    Ещё какой-то текст
    TICKET-12362_Иван Иванов_В работе_2024-01-15
    TICKET-12363_Иван Иванов_В работе_2024-01-16
    
    `

    user := "Иван Иванов"
    //status := "В работе"
    tasks := GetTasks(chatHistory1, &user, nil)
    //tasks1 := GetTasks(chatHistory, nil, &status)

    fmt.Println("tasks ",tasks)
    fmt.Println("--------------------------")
    //fmt.Println("tasks1 ",tasks1)

    //fmt.Println(GetTasks(text,nil,nil))
    fmt.Println("--------------------------")



    //fmt.Println(Ticket.Date.tickets[0])
    // fmt.Println(tickets)
    // fmt.Println(tickets[1].Date)


    //КЛЮЧЕВОЙ МОМЕНТ БЫЛ: СДЕЛАТЬ trings.TrimSpace()
}