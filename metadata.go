package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	URL string `md:"URL,required"`
}

type Input struct {
	QueryInput string `md:"queryInput,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["queryInput"])
	r.QueryInput = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"queryInput": r.QueryInput,
	}
}

type Output struct {
	ResponseCode int `md:"responseCode"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["responseCode"])
	o.ResponseCode = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"responseCode": o.ResponseCode,
	}
}
