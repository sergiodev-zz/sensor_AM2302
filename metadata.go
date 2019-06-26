package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
}

type Input struct {
}

func (r *Input) FromMap(values map[string]interface{}) error {
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return nil
}

type Output struct {
	AnOutput string `md:"anOutput"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["anOutput"])
	o.AnOutput = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anOutput": o.AnOutput,
	}
}
