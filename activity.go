package sample

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"net/http"
 	"net/url"
)

func init() {
	_ = activity.Register(&Activity{},New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	c := &http.Client{}

	//ctx.Logger().Debugf("Setting: %s", s.API_BASE_URL)

	act := &Activity{settings: s, client: c} //add aSetting to instance

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

	urlString := a.settings.API_BASE_URL

	if urlString!= "" {
		url,err := url.Parse(urlString)
 		if err != nil {
 			return true, err
 		}

		method := "GET"

		req, _ := http.NewRequest(method, url.String(), nil)
		resp, err := a.client.Do(req)

		output := &Output{ResponseCode: resp.StatusCode}
		err = ctx.SetOutputObject(output)
		if err != nil {
			return true, err
		}

		return true, nil
	}

	return true, activity.NewError("API Gateway URL is not provided","",nil)
	//ctx.Logger().Debugf("Input: %s", input.ProcessType)
	//ctx.Logger().Debugf("Input: %d", input.PathParamId)
	//ctx.Logger().Debugf("Input: %d", input.CollectionId)
	//ctx.Logger().Debugf("Input: %d", input.LocationId)
	//ctx.Logger().Debugf("Input: %d", input.ActivityId)

}