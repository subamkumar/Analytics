package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	API_BASE_URL string `md:"API_Base_URI,required"`
}

type Input struct {
	ProcessType 	string	`md:"processType,required"`
	PathParamId 	int		`md:"getId"`
	CollectionId	int		`md:"filterByCollectionId"`
	LocationId		int		`md:"filterByLocationId"`
	ActivityId 		int		`md:"filterByActivityId"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	processTypeVal, _ := coerce.ToString(values["processType"])
	r.ProcessType = processTypeVal

	pathParamIdVal, _ := coerce.ToInt(values["getId"])
	r.PathParamId = pathParamIdVal

	collectionIdVal, _ := coerce.ToInt(values["filterByCollectionId"])
	r.CollectionId = collectionIdVal

	locationIdVal, _ := coerce.ToInt(values["filterByLocationId"])
	r.LocationId = locationIdVal

	activityIdVal, _ := coerce.ToInt(values["filterByActivityId"])
	r.ActivityId = activityIdVal

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"processType": r.ProcessType,
		"getId": r.PathParamId,
		"filterByCollectionId": r.CollectionId,
		"filterByLocationId": r.LocationId,
		"filterByActivityId": r.ActivityId,
	}
}

type Output struct {
	ResponseCode int 			`md:"responseCode"`
	ResponseData interface{} 	`md:"responseData"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	responseCodeVal, _ := coerce.ToInt(values["responseCode"])
	o.ResponseCode = responseCodeVal

	responseDataVal, _ := coerce.ToInt(values["responseData"])
	o.ResponseData = responseDataVal

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"responseCode": o.ResponseCode,
		"responseData": o.ResponseData,
	}
}