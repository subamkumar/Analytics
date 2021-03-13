package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	API_BASE_URL string `md:"API_Base_URI,required"`
}

type Input struct {
	ProcessType 	string `md:"processType,required"`
	PathParamId 	string `md:"getId"`
	CollectionId	string `md:"filterByCollectionId"`
	LocationId		string `md:"filterByLocationId"`
	ActivityId 		string `md:"filterByActivityId"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	processTypeVal, _ := coerce.ToString(values["processType"])
	r.ProcessType = processTypeVal

	pathParamIdVal, _ := coerce.ToString(values["getId"])
	r.PathParamId = pathParamIdVal

	collectionIdVal, _ := coerce.ToString(values["filterByCollectionId"])
	r.CollectionId = collectionIdVal

	locationIdVal, _ := coerce.ToString(values["filterByLocationId"])
	r.LocationId = locationIdVal

	activityIdVal, _ := coerce.ToString(values["filterByActivityId"])
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