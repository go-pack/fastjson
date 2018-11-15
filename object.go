package fastjson

type JSONObject struct {
	data interface{}
}

func (obj JSONObject) GetInteger(key string) int32 {
	i := obj.data.(map[string]interface{})
	return int32(i[key].(float64))
}

func (obj JSONObject) GetString(key string) string {
	i := obj.data.(map[string]interface{})
	return i[key].(string)
}
func (obj JSONObject) GetJSONArray(key string) *JSONArray {
	i := obj.data.(map[string]interface{})
	i2 := i[key].([]interface{})
	return &JSONArray{data:i2[:]}
}
func (obj JSONObject) EntrySet() (map[string]JSONObject,error) {
	hashmap := make(map[string]JSONObject)
	for key,v := range obj.data.(map[string]interface{}){
		hashmap[key] = JSONObject{data:v}
	}

	return hashmap,nil
}