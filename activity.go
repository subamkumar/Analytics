package sample

import (
	"bytes"
	"encoding/json"
 	"github.com/project-flogo/core/activity"
 	"github.com/project-flogo/core/data/metadata"
 	"io"
 	"net/http"
 	"net/url"
)

func init() {
	_ = activity.Register(&Activity{},New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{},&Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}
	
	c := &http.Client{}

	act := &Activity{settings: s, client: c} //add Setting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
	settings	*Settings
	client		*http.Client
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	
	urlString := a.settings.URL

	if urlString!= "" {
 		url,err := url.Parse(urlString)
 		if err != nil {
 			return true, err
 		}

 		method := "POST"

		bodyData := make(map[string]interface{})
		
		ctx.Logger().Infof("====TESTING DATA: ", bodyData)
		if input.ProcessURL != "" {
			ctx.Logger().Infof("===PROCESSURL: ", input.ProcessURL)	
			bodyData["process_url"] = input.ProcessURL
			ctx.Logger().Infof("BodyData: ", bodyData)
		}
		
		if input.ProcessorType != "" {
			ctx.Logger().Infof("===PROCESSTYPE: ", input.ProcessorType)
			bodyData["processor_type"] = input.ProcessorType
			ctx.Logger().Infof("BodyData: ", bodyData)
		}
		
		if input.Parameters != nil {
			ctx.Logger().Infof("===PARAM: ", input.Parameters)
			bodyData["parameters"] = input.Parameters
			ctx.Logger().Infof("BodyData: ", bodyData)
		}
		
		if input.Log != nil {
			ctx.Logger().Infof("===LOG: ", input.Log)
			bodyData["log"] = input.Log
			ctx.Logger().Infof("BodyData: ", bodyData)
		}
		
		if input.PostProcesses != nil {
			ctx.Logger().Infof("===POST_PROCESSES: ", input.PostProcesses)
			bodyData["post_processes"] = input.PostProcesses
			ctx.Logger().Infof("BodyData: ", bodyData)
		}
		
		b,err := json.Marshal(bodyData)
		
		if err != nil {
			return true, err
		}

		var body io.Reader
		body = bytes.NewBuffer(b)

	
		req, _ := http.NewRequest(method, url.String(), body)

		resp, err := a.client.Do(req)

		ctx.Logger().Info("======")
		ctx.Logger().Infof(resp.Status)

		output := &Output{ResponseCode: 200}
		err = ctx.SetOutputObject(output)
		if err != nil {
			return true, err
		}

 		return true, nil
 	}
 	return true, activity.NewError("API Gateway URL is not provided","",nil)
}
