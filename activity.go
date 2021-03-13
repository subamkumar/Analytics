package sample

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"net/http"
 	"net/url"
	"strconv"
	"strings"
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
		_,err := url.Parse(urlString)
 		if err != nil {
 			return true, err
 		}
		
		//method := "GET"

		if input.ProcessType != "" {

			if strings.LastIndex(urlString,"/") == len(urlString)-1 {
				urlString = urlString+input.ProcessType
			}else{
				urlString = urlString+"/"+input.ProcessType
			}


			if input.PathParamId != 0 {
				urlString = urlString+"/"+strconv.Itoa(input.PathParamId)
			}else{

				queryString := ""

				if input.CollectionId != 0 {
					queryString = queryString+"collection_id="+strconv.Itoa(input.CollectionId)
				}

				if input.ActivityId != 0 {
					if len(queryString) > 0 {
						queryString = queryString+"&"+"activity_id="+strconv.Itoa(input.ActivityId)
					}else{
						queryString = queryString+"activity_id="+strconv.Itoa(input.ActivityId)
					}
				}

				if input.LocationId != 0 {
					if len(queryString) > 0 {
						queryString = queryString+"&"+"location_id="+strconv.Itoa(input.LocationId)
					}else{
						queryString = queryString+"location_id="+strconv.Itoa(input.LocationId)
					}
				}

				if len(queryString) > 0 {
					urlString = urlString+"?"+queryString
				}
			}

			ctx.Logger().Debugf("FORMED_URL: %s", urlString)

			//req, _ := http.NewRequest(method, urlString, nil)
			//resp, err := a.client.Do(req)

			output := &Output{ResponseCode: 2020}
			err = ctx.SetOutputObject(output)
			if err != nil {
				return true, err
			}

			return true, nil
		}
		return true, activity.NewError("Required Process Type is not provided","",nil)
	}

	return true, activity.NewError("API Gateway URL is not provided","",nil)
	//ctx.Logger().Debugf("Input: %s", input.ProcessType)
	//ctx.Logger().Debugf("Input: %d", input.PathParamId)
	//ctx.Logger().Debugf("Input: %d", input.CollectionId)
	//ctx.Logger().Debugf("Input: %d", input.LocationId)
	//ctx.Logger().Debugf("Input: %d", input.ActivityId)

}