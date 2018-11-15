package fastjson

import json2 "encoding/json"

func ParseObject(){
	
}
func ParseArray(jsonByte []byte) JSONArray {
	var dat []interface{}
	json2.Unmarshal(jsonByte, &dat)
	return JSONArray{data:dat}
}