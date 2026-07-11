package main

import (
	"fmt"
)

// Order представляет информацию о заказе.
type Order struct {
    OrderNumber  int
    CustomerName string
    OrderAmount  float64
}

// OrderLogger представляет журнал заказов и хранит записи о заказах.
type OrderLogger struct{
	OrderRecording string
}

// NewOrderLogger создает новый экземпляр OrderLogger.
//тут создаем конструктор, который создаст экземпляр структуры и это будет не копия, а оригинал который можно изменять. 
func NewOrderLogger() *OrderLogger {
    return &OrderLogger{} // это создание экземпляра с полями значения которых по умолчанию пустые, и в дальнейшем поля этого экземпляра можно заполнять значениями
}


//нужно запис. в терминал
//так, тут вызывающий код будет передвать копию структуры Order - это значит что огригинал изменять нельзя, для этого нужно исп. указатель *Order
func (logger *OrderLogger) AddOrder(order Order) {
	logger.OrderRecording = fmt.Sprintf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)

	fmt.Print(logger.OrderRecording)
}

func main() {
	recording := NewOrderLogger()

	order := Order{
		OrderNumber: 101,
		CustomerName: "Анна",
		OrderAmount: 123.45,
	}

	recording.AddOrder(order)
}

