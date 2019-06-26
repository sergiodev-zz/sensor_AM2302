package sample

import "github.com/project-flogo/core/data/coerce"

type Output struct {
	OutputMessage string `md:"OutputMessage"`
	Temp          string `md:"Temp"`
	Humi          string `md:"Humi"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["OutputMessage"])
	o.OutputMessage = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"OutputMessage": o.OutputMessage,
		"Temp":          o.Temp,
		"Humi":          o.Humi,
	}
}
