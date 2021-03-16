package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	API_Gateway string `md:"API_Gateway,required"`
}

type Input struct {
	ProcessURL		string						`md:"process_url"`
	ProcessType 	string						`md:"process_type"`
	Parameters		map[string]interface{}		`md:"parameters"`
	Log				map[string]interface{}		`md:"log"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	processURLVal, _ := coerce.ToString(values["process_url"])
	r.ProcessURL = processURLVal

	processTypeVal, _ := coerce.ToString(values["process_type"])
	r.ProcessType = processTypeVal

	parametersVal, _ := coerce.ToObject(values["parameters"])
	r.Parameters = parametersVal

	logVal, _ := coerce.ToObject(values["log"])
	r.Log = logVal

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"process_url": r.ProcessURL,
		"process_type": r.ProcessType,
		"parameters": r.Parameters,
		"log": r.Log,
	}
}

type Output struct {
}

func (o *Output) FromMap(values map[string]interface{}) error {
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{}
}