package sample

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"net/http"
	"net/url"
	"strings"
)

func init() {
	_ = activity.Register(&Activity{},New)
}

var activityMd = activity.ToMetadata(&Settings{},&Input{}, &Output{})

func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}
	
	c := &http.Client{}

	act := &Activity{settings: s, client: c}

	return act, nil
}


type Activity struct {
	settings	*Settings
	client		*http.Client
}


func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}


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

 		method := "GET"

 		if input.ProcessorType != "" {

			if strings.LastIndex(urlString,"/") == len(urlString)-1 {
				urlString = urlString+input.ProcessorType
			}else{
				urlString = urlString+"/"+input.ProcessorType
			}


			if input.getById != 0 {
				urlString = urlString+"/"+string(input.getById)
			}else{
				if input.CollectionQueryId != 0 {
					urlString = urlString+"?collection_id"+string(input.CollectionQueryId)
				}

				if input.ActivityQueryId != 0 {
					urlString = urlString+"?activity_id"+string(input.ActivityQueryId)
				}

				if input.LocationQueryId != 0 {
					urlString = urlString+"?location_id"+string(input.LocationQueryId)
				}
			}

			req, _ := http.NewRequest(method, url.String(), nil)
			resp, err := a.client.Do(req)

			output := &Output{ResponseCode: resp.StatusCode}
			err = ctx.SetOutputObject(output)
			if err != nil {
				return true, err
			}

			return true, nil
		}
 	}
 	return true, activity.NewError("API Gateway URL is not provided","",nil)
}
