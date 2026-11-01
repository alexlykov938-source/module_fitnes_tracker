package actioninfo

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			continue
		}

		println(info)
	}
}
