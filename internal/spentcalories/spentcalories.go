package spentcalories

import (
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
		log.Println("Длина слайсов не равна трем")
		return 0, " ", 0, fmt.Errorf("")
	}
	steps, err := strconv.Atoi(stringsSlice[0])
	if err != nil {
		log.Printf("%v", err)
		return 0, "", 0, err
	}

	if steps <= 0 {
		log.Println("Шаги < или = 0")
		return 0, "", 0, fmt.Errorf("")
	}
	duration, err := time.ParseDuration(stringsSlice[2])
	if err != nil {
		log.Printf("%v", err)
		return 0, "", 0, err
	}
	if duration <= 0 {
		log.Println("duration <= 0")
		return 0, "", 0, fmt.Errorf("")
	}
	return steps, stringsSlice[1], duration, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distance := (float64(steps) * stepLength) / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := distance(steps, height)
	avgSpead := distance / duration.Hours()
	return avgSpead
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	if weight <= 0 {
		log.Println("вес не может быть <= 0")
		return "", fmt.Errorf("")
	}
	var trainingType string
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	switch trainingType {
	case "Бег":
		distance := distance(steps, height)
		meanSpeed := meanSpeed(steps, height, duration)
		caloriesRun, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Printf("ошибка, %v", err)
			return "", err

		}
		runResult := fmt.Sprintf(`Тип тренировки: %s
Длительность: %0.2f ч.
Дистанция: %0.2f км.
Скорость: %0.2f км/ч
Сожгли калорий: %0.2f
`, trainingType, duration.Hours(), distance, meanSpeed, caloriesRun)
		return runResult, nil
	case "Ходьба":
		distance := distance(steps, height)
		meanSpeed := meanSpeed(steps, height, duration)
		caloriesWalk, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err

		}
		walkResult := fmt.Sprintf(`Тип тренировки: %s
Длительность: %0.2f ч.
Дистанция: %0.2f км.
Скорость: %0.2f км/ч
Сожгли калорий: %0.2f
`, trainingType, duration.Hours(), distance, meanSpeed, caloriesWalk)
		return walkResult, nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")

	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		log.Println("ошибка при рассчете: пройдено 0 шагов")
		return 0, fmt.Errorf("")
	}
	if weight <= 0 {
		log.Println("вес не может быть <= 0")
		return 0, fmt.Errorf("")
	}
	if height <= 0 {
		log.Println("рост не может быть <= 0 ")
		return 0, fmt.Errorf("")
	}
	if duration <= 0 {
		log.Println("duration <= 0")
		return 0, fmt.Errorf("")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration / time.Minute
	calories := (weight * meanSpeed * float64(durationInMinutes)) / minInH
	return calories, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("")
	}
	if height <= 0 {
		return 0, fmt.Errorf("")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	timeInMinutes := duration / time.Minute
	calories := (weight * meanSpeed * float64(timeInMinutes)) / minInH
	caloriesWithCoefficient := calories * walkingCaloriesCoefficient
	return caloriesWithCoefficient, nil

}
