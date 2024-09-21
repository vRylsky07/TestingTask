package config

import (
	"flag"
)

var (
	DefSize              = 50    // стандартное количество игроков
	DefAttemptsPerPlayer = 25    // стандартное количество попыток
	DefSessionsCount     = 10000 // стандартный размер сессии
)

// Создаёт флаги для изменения стандартных параметров.
func Flags() {
	flag.IntVar(&DefSize, "size", DefSize, "changes the number of size")
	flag.IntVar(&DefAttemptsPerPlayer, "attempts", DefAttemptsPerPlayer, "changes the number of attempts")
	flag.IntVar(&DefSessionsCount, "sessions", DefSessionsCount, "changes the number of sessions")
	flag.Parse()
}
