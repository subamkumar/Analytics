package sample

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
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
	

	act := &Activity{settings: s} //add Setting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
	settings	*Settings
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

	ctx.Logger().Infof("API Gateway URL", a.settings.URL)

	ctx.Logger().Infof("Input Process URL: ", input.ProcessURL)

	ctx.Logger().Infof("Input Param: ", input.Parameters)

	/*client := &http.Client{}
	
	urlString := input.URL
	url,err := url.Parse(urlString)
	if err != nil {
		return true, err
	}
	req, _ := http.NewRequest("GET",url.String(),nil)
	resp, err := client.Do(req)

	ctx.Logger().Infof(resp.Status)

	output := &Output{ResponseCode: 200}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}*/

	return true, nil
}
