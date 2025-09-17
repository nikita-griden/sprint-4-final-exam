package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	splitString := strings.Split(data, ",") // Разделяем строку по , на слайс строк
	if len(splitString) != 2 {              // проверяем что длина слайса равна 2
		return 0, 0, errors.New("длина слайса не равна 2")
	}

	steps, err := strconv.Atoi(splitString[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errors.New("steps не может быть <= 0")
	}
	duration, err := time.ParseDuration(splitString[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, errors.New("продолжительность не может быть < или = 0")
	}
	return steps, duration, nil

}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	if weight <= 0 {
		return ""
	}
	if height <= 0 {
		return ""
	}
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("err = %v", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distM := float64(steps) * stepLength
	distKm := distM / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}

	result := fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, steps, distKm, calories)
	return result

}
