package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	API_Gateway string `md:"API_Gateway,required"`
}

type Input struct {
	Collection string `md:"collection,required"`
	Location   string `md:"location,required"`
	Activity   string `md:"activity,required"`
	Username   string `md:"username"`
	Password   string `md:"password"`
	RequestId  string `md:"requestId"`
	SecretId   string `md:"secretId"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	collectionVal, _ := coerce.ToString(values["collection"])
	r.Collection = collectionVal

	locationVal, _ := coerce.ToString(values["location"])
	r.Location = locationVal

	activityVal, _ := coerce.ToString(values["activity"])
	r.Activity = activityVal

	usernameVal, _ := coerce.ToString(values["username"])
	r.Username = usernameVal

	passwordVal, _ := coerce.ToString(values["password"])
	r.Password = passwordVal

	requestVal, _ := coerce.ToString(values["requestId"])
	r.RequestId = requestVal

	secretVal, _ := coerce.ToString(values["secretId"])
	r.SecretId = secretVal

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"collection": r.Collection,
		"location":   r.Location,
		"activity":   r.Activity,
		"username":   r.Username,
		"password":   r.Password,
		"requestId":  r.RequestId,
		"secretId":   r.SecretId,
	}
}

type Output struct {
}

func (o *Output) FromMap(values map[string]interface{}) error {
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{}
}
