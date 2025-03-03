package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
	splited := strings.Split(data, ",")
	var err error = nil
	if len(splited) != 3 {
		return 0, "", 0, errors.New("формат строки не соответствует")
	}
	steps, errint := strconv.Atoi(splited[0])
	duration, errtime := time.ParseDuration(splited[2])
	if errint != nil {
		return 0, "", 0, errint
	}
	if errtime != nil {
		return 0, "", 0, errtime
	}
	return steps, splited[1], duration, err

}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func Distance(steps int) float64 {
	// ваш код ниже
	return float64(steps) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration == 0 {
		return 0
	}
	speed := float64(steps) * float64(lenStep) / 1000 / (duration.Hours())
	return speed
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, style, duration, err := parseTraining(data)
	var resultLine string = ""
	if err != nil {
		return ""
	}
	switch style {
	case "Бег":
		//call := WalkingSpentCalories(steps, weight, height, duration)
		//call = +1
		resultLine = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", style, duration.Hours(), Distance(steps), meanSpeed(steps, duration), RunningSpentCalories(steps, weight, duration))
	case "Ходьба":
		call := WalkingSpentCalories(steps, weight, height, duration)
		resultLine = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", style, duration.Hours(), Distance(steps), meanSpeed(steps, duration), call)
	}
	return resultLine
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed(steps, duration)) - runningCaloriesMeanSpeedShift) * weight
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed(steps, duration)*meanSpeed(steps, duration)/height)*walkingSpeedHeightMultiplier) * duration.Hours() * float64(minInH)
}
