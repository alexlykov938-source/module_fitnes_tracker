package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")

	if len(data) != 3 {
		return fmt.Errorf("invalid training data")
	}

	t.Steps, err = strconv.Atoi(data[0])

	if err != nil || t.Steps <= 0 {
		return fmt.Errorf("invalid steps value: %v", err)
	}

	t.TrainingType = data[1]

	t.Duration, err = time.ParseDuration(data[2])

	if err != nil || t.Duration <= 0 {
		return fmt.Errorf("invalid duration value: %v", err)
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	if strings.ToUpper(t.TrainingType) == "ХОДЬБА" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else if strings.ToUpper(t.TrainingType) == "БЕГ" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	if err != nil {
		return "", fmt.Errorf("error calculating spent calories: %v", err)
	}

	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)
	return info, nil
}
