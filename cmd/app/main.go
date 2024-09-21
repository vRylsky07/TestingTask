package main

import (
	"fmt"

	cfg "github.com/Ra1nz0r/iteco-1/internal/config"
	"github.com/Ra1nz0r/iteco-1/internal/player"
	"github.com/Ra1nz0r/iteco-1/internal/session"

	"github.com/dariubs/percent"
)

func main() {
	cfg.Flags()

	// Выбираем вариант со случайным выбором номера шкатулок.
	var mode player.PlayerType = player.WithRandom

	res, errRun := Run(mode, cfg.DefSize, cfg.DefAttemptsPerPlayer, cfg.DefSessionsCount)
	if errRun != nil {
		panic(errRun)
	}
	fmt.Printf("Процент побед при случайном выборе, сессия из %d игр: %.0f%%.\n", cfg.DefSessionsCount, res)

	// Переключаем на режим, где игроки договорились о способе выбора.
	mode = player.WithOrder

	res, errRun = Run(mode, cfg.DefSize, cfg.DefAttemptsPerPlayer, cfg.DefSessionsCount)
	if errRun != nil {
		panic(errRun)
	}

	fmt.Printf("Процент побед при договорённости между игроками, сессия из %d игр: %.0f%%\n", cfg.DefSessionsCount, res)

}

func Run(p player.PlayerType, size, attemptsPerPlayer, sessionsCount int) (float64, error) {
	if attemptsPerPlayer > size || size <= 0 {
		return 0, fmt.Errorf("incorrect number of players or chances of attempts are greater than players")
	}

	var playersArr []player.Unit

	successedCount := 0
	// Запускаем цикл игровых сессий.
	for i := 0; i < sessionsCount; i++ {

		// В зависимости от Enum PlayerType выбираем реализацию интерфейса Unit и инициализируем переменную.
		switch p {
		case player.WithRandom:
			playersArr = player.CreatePlayersWithRandom(size, attemptsPerPlayer)
		case player.WithOrder:
			playersArr = player.CreatePlayersWithOrder(size, attemptsPerPlayer)
		}

		session, errSess := session.NewGameSession(size, playersArr)
		if errSess != nil {
			return 0, fmt.Errorf("%w", errSess)
		}

		ok, errPS := session.PlaySession()
		if errPS != nil {
			return 0, fmt.Errorf("%w", errPS)
		}

		if ok {
			successedCount++
		}
	}

	return percent.PercentOf(successedCount, sessionsCount), nil
}
