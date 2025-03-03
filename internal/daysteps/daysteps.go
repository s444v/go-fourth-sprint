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
	var err error = nil
	if len(splited) != 2 {
		return 0, 0, errors.New("формат строки не соответствует")
	}
	steps, errint := strconv.Atoi(splited[0])
	duration, errtime := time.ParseDuration(splited[1])
	if errint != nil {
		return 0, 0, errint
	}
	if errtime != nil {
		return 0, 0, errtime
	}
	return steps, duration, err
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
	resultLine := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, spentcalories.Distance(steps), spentcalories.WalkingSpentCalories(steps, weight, height, time))
	if err != nil {
		resultLine = ""
	}
	return resultLine
}
