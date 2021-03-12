package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct{
	URL string `md:"API Gateway URL,required"`
}

type Input struct {
	ProcessorType 		string 					`md:"processType,required"`
	getById 			int 					`md:"getId"`
	CollectionQueryId 	int	 					`md:"filterByCollectionId"`
	LocationQueryId		int						`md:"filterByLocationId"`
	ActivityQueryId 	int		 				`md:"filterByActivityId"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	processTypeVal, _ := coerce.ToString(values["processType"])
	r.ProcessorType = processTypeVal

	pathGetId, _ := coerce.ToInt(values["getId"])
	r.getById = pathGetId

	collectionFilter, _ := coerce.ToInt(values["filterByCollectionId"])
	r.CollectionQueryId = collectionFilter

	locationFilter, _ := coerce.ToInt(values["filterByLocationId"])
	r.LocationQueryId = locationFilter

	activityFilter, _ := coerce.ToInt(values["filterByActivityId"])
	r.ActivityQueryId = activityFilter

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"processType": r.ProcessorType,
		"getId": r.getById,
		"filterByCollectionId": r.CollectionQueryId,
		"filterByLocationId": r.LocationQueryId,
		"filterByActivityId": r.ActivityQueryId,
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
