package fastjson

import "errors"

type JSONArray struct {
	data []interface{}
}

func (array JSONArray) Size() int {
	return len(array.data)
}
func (array JSONArray) GetJSONObject(index int) (*JSONObject,error) {
	if index > array.Size() {
		return nil,errors.New("index out of")
	}
	return &JSONObject{data:array.data[index]}, nil
}
func (array JSONArray) GetString(index int) (string,error) {
	if index > array.Size() {
		return "",errors.New("index out of")
	}
	return array.data[index].(string), nil
}
