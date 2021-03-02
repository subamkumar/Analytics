package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct{
	URL string `md:"API Gateway URL,required"`
}

type Input struct {
	ProcessURL 		string 					`md:"process_url"`
	ProcessorType 	string 					`md:"processor_type"`
	Parameters 		map[string]interface{} 	`md:"parameters"`
	Log				map[string]interface{}	`md:"log"`
	PostProcesses 	[]interface{} 			`md:"post_processes"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	processURLVal, _ := coerce.ToString(values["process_url"])
	r.ProcessURL = processURLVal

	processTypeVal, _ := coerce.ToString(values["processor_type"])
	r.ProcessorType = processTypeVal

	parametersVal, _ := coerce.ToObject(values["parameters"])
	r.Parameters = parametersVal

	logVal, _ := coerce.ToObject(values["log"])
	r.Log = logVal

	postProcessesVal, _ := coerce.ToArray(values["post_processes"])
	r.PostProcesses = postProcessesVal

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"process_url": r.ProcessURL,
		"processor_type": r.ProcessorType,
		"parameters": r.Parameters,
		"log": r.Log,
		"post_processes": r.PostProcesses,
	}
}

type Output struct {
	ResponseCode int `md:"responseCode"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToInt(values["responseCode"])
	o.ResponseCode = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"responseCode": o.ResponseCode,
	}
}
