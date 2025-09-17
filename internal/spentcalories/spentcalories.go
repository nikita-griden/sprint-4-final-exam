package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	stringsSlice := strings.Split(data, ",")
	if len(stringsSlice) != 3 {
		return 0, " ", 0, errors.New("длина слайсов не равна трем")
	}
	steps, err := strconv.Atoi(stringsSlice[0])
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, errors.New("шаги не могут быть < или = 0")
	}
	duration, err := time.ParseDuration(stringsSlice[2])
	if err != nil {
		return 0, "", 0, err
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("duration <=0")
	}
	return steps, stringsSlice[1], duration, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distanceResult := (float64(steps) * stepLength) / mInKm
	return distanceResult
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distanceForSpeed := distance(steps, height)
	avgSpead := distanceForSpeed / duration.Hours()
	return avgSpead
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	var (
		distanceForTraining  float64
		meanSpeedForTraining float64
		trainingType         string
	)

	// TODO: реализовать функцию
	if weight <= 0 {
		return "", errors.New("вес не может быть <= 0")
	}
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if steps <= 0 {
		return "", errors.New("шаги не могут быть < или = 0")
	}
	distanceForTraining = distance(steps, height)
	meanSpeedForTraining = meanSpeed(steps, height, duration)
	switch trainingType {
	case "Бег":
		caloriesRun, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", errors.New("err")

		}
		runResult := fmt.Sprintf(`Тип тренировки: %s
Длительность: %0.2f ч.
Дистанция: %0.2f км.
Скорость: %0.2f км/ч
Сожгли калорий: %0.2f
`, trainingType, duration.Hours(), distanceForTraining, meanSpeedForTraining, caloriesRun)
		return runResult, nil
	case "Ходьба":
		caloriesWalk, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err

		}
		walkResult := fmt.Sprintf(`Тип тренировки: %s
Длительность: %0.2f ч.
Дистанция: %0.2f км.
Скорость: %0.2f км/ч
Сожгли калорий: %0.2f
`, trainingType, duration.Hours(), distanceForTraining, meanSpeedForTraining, caloriesWalk)
		return walkResult, nil
	default:
		return "", errors.New("неизвестный тип тренировки")

	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("ошибка при рассчете: пройдено 0 шагов")
	}
	if weight <= 0 {
		return 0, errors.New("вес не может быть <=0")
	}
	if height <= 0 {
		return 0, errors.New("рост не может быть <= 0 ")
	}
	if duration <= 0 {
		return 0, errors.New("duration <= 0")
	}
	meanSpeedForRunning := meanSpeed(steps, height, duration)
	caloriesForRunning := (weight * meanSpeedForRunning * float64(duration.Minutes())) / minInH
	return caloriesForRunning, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("steps не может быть <= 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight не может быть <= 0")
	}
	if height <= 0 {
		return 0, errors.New("height не может быть <= 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration не может быть <= 0")
	}
	meanSpeedForWalking := meanSpeed(steps, height, duration)
	timeInMinutes := duration / time.Minute
	caloriesForWalking := (weight * meanSpeedForWalking * float64(timeInMinutes)) / minInH
	caloriesWithCoefficient := caloriesForWalking * walkingCaloriesCoefficient
	return caloriesWithCoefficient, nil

}
