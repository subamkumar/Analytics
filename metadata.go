package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct{
	URL string `md:"API Gateway URL,required"`
}

type Input struct {
	//ProcessorType 		string 					`md:"processType"`
	getById 			string 					`md:"getId"`
	//CollectionQueryId 	string	 				`md:"filterByCollectionId"`
	//LocationQueryId		string					`md:"filterByLocationId"`
	//ActivityQueryId 	string		 			`md:"filterByActivityId"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	/*processTypeVal, _ := coerce.ToString(values["processType"])
	r.ProcessorType = processTypeVal*/

	pathGetId, _ := coerce.ToString(values["getId"])
	r.getById = pathGetId

	/*collectionFilter, _ := coerce.ToString(values["filterByCollectionId"])
	r.CollectionQueryId = collectionFilter

	locationFilter, _ := coerce.ToString(values["filterByLocationId"])
	r.LocationQueryId = locationFilter

	activityFilter, _ := coerce.ToString(values["filterByActivityId"])
	r.ActivityQueryId = activityFilter*/

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		//"processType": r.ProcessorType,
		"getId": r.getById,
		//"filterByCollectionId": r.CollectionQueryId,
		//"filterByLocationId": r.LocationQueryId,
		//"filterByActivityId": r.ActivityQueryId,
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
