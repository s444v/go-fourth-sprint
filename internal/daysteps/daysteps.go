package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/s444v/go-fourth-sprint/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	splited := strings.Split(data, ",")
	if len(splited) != 2 {
		return 0, 0, errors.New("формат строки не соответствует")
	}
	steps, err := strconv.Atoi(splited[0])
	if err != nil {
		return 0, 0, err
	}
	duration, err := time.ParseDuration(splited[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, time, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, spentcalories.Distance(steps), spentcalories.WalkingSpentCalories(steps, weight, height, time))
}
