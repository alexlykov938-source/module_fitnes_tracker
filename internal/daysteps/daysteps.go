package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")

	if len(data) != 2 {
		return fmt.Errorf("invalid day steps data")
	}

	ds.Steps, err = strconv.Atoi(data[0])
	if err != nil || ds.Steps <= 0 {
		return fmt.Errorf("invalid steps value: %v", err)
	}

	ds.Duration, err = time.ParseDuration(data[1])
	if err != nil || ds.Duration <= 0 {
		return fmt.Errorf("invalid duration value: %v", err)
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)

	return info, nil
}
