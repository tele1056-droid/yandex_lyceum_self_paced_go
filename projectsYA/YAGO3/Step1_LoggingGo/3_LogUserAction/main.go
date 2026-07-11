package main

import (
	"log/slog"
)

//т.е. сюда передают объект *slog.Logger - это логгер, который предоставляет методы для записи структурированных логов. Например: logger := slog.Default() // возвращает *slog.Logger
func LogUserAction(logger *slog.Logger, user string, action string) {
	//применяем к переданому объекту *slog.Logger формат Info 
	logger.Info("user action", "User", user, "Action", action)
}

/*
БАЗОВЫЙ СИНТАКСИС
logger.Info("Сообщение", "ключ", "значение")

logger.Info("Пользователь вошёл", "name", "Alice")
// time=... level=INFO msg="Пользователь вошёл" name=Alice
*/

func main() {
	user := "Kate"
	action := "Order"
	logger := slog.Default()

	LogUserAction(logger,user, action)
}