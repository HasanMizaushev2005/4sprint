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
	slice := strings.Split(data, ",")
	if len(slice) != 3 {
		return 0, "", 0, errors.New("Недостаточно данных в стоке")
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, errors.New("Ошибка парсинга (количество шагов равно 0)")
	}

	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, err
	}

	if duration <= 0 {
		return 0, "", 0, errors.New("Неверное время")
	}

	return steps, slice[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	d := float64(steps) * (height * stepLengthCoefficient) / float64(mInKm)
	return d
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	d := distance(steps, height)
	return d / float64(duration.Hours())
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if steps <= 0 {
		return "", err
	}

	var calories float64
	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity,
		duration.Hours(),
		distance(steps, height),
		meanSpeed(steps, height, duration),
		calories,
	)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("Некоректные входящие данные")
	}

	if duration <= 0 {
		return 0, errors.New("Неверное время")
	}

	durationInMinutes := float64(duration.Minutes())

	calories := (weight * meanSpeed(steps, height, duration) * durationInMinutes) / float64(minInH)
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("Некоректные входящие данные")
	}

	if duration <= 0 {
		return 0, errors.New("Неверное время")
	}

	durationInMinutes := float64(duration.Minutes())

	calories := ((weight * meanSpeed(steps, height, duration) * durationInMinutes) / float64(minInH)) * walkingCaloriesCoefficient
	return calories, nil
}
