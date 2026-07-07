package main

import (
	"strings"
	"time"
	"fmt"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	var correct int

	for i, _ := range questions {
		//тоесть в списке с вопросами соотносятся по порядку ответы из списка, и если в итерацию вопроса ответ польз. совпадает с ответом из списка, то увелич. счетчик на 1
		select {
		case answer := <-answerCh:
			//проверяем ответ пользователя который мы прочитали из канала в answer, используем strings.EqualFold чтобы сравнивать строки независимо от регистра
			if strings.EqualFold(strings.TrimSpace(answer), strings.TrimSpace(answers[i])) {
				correct++ 
			} //!!! тут был затуп, что отправлял на проверку без обработки strings.TrimSpace, и один ответ не проходил, нужно было обработать, как минимум, ответы пользователя на лишние пробелы (в начале и в конце)
			
		case <-time.After(1000 * time.Millisecond): //тут 1000 милисек. - это 1 сек.
			// тут если таймаут быстрее сработал, то время вышло и начнется следующая итерация вопросов
		}
	}
	return correct
}

func main () {
	// questions := []string{"Q1", "Q2", "Q3"}
	// answers := []string{"A1", "A2", "A3"}
	// answerCh := make(chan string)

	// go func() {
	// 	answerCh <- "A1"
	// 	answerCh <- "A2"
	// 	answerCh <- "A3"
	// }()

	// result := QuizRunner(questions, answers, answerCh)
	// if result != 3 {
	// 	fmt.Println("Expected 3 correct, got %d", result)
	// }
	questions := []string{"Q1", "Q2"}
    answers := []string{"GO", "Yes"}

    answerCh := make(chan string)

	go func() {
        answerCh <- "go"    // нижний регистр, и вот тут в автотесте был вариан ввода пользователя с пробелом, например "go ", и тут надо strings.TrimSpace()
        answerCh <- "YES"   // верхний регистр
    }()

	result := QuizRunner(questions, answers, answerCh)

	fmt.Println(result)
}

/*
тут по сигнатуре функции будут передавать канал с ответами пользователя, и нужно будет ответы польз. (из канала) сравнивать с ответами из слайса. Ну, то есть внешний код будет отправлять ответы пользователя, а функция QuizRunner будет читать из этого канала и проверять ответы. Т.е. моя функция не содаёт канал - она его принимает, моя функц. не закрывает канал - это делает внешний код, моя функц. читает из канала
*/