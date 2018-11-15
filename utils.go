package fastjson

import json2 "encoding/json"

func ParseObject(jsonByte []byte) JSONObject{
	var dat map[string]interface{}
	json2.Unmarshal(jsonByte, &dat)
	return JSONObject{data:dat}
}
func ParseArray(jsonByte []byte) JSONArray {
	var dat []interface{}
	json2.Unmarshal(jsonByte, &dat)
	return JSONArray{data:dat}
}