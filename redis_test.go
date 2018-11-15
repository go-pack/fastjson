package fastjson_test

import (
	"github.com/go-pack/fastjson"
	"github.com/go-redis/redis"
	"io/ioutil"
	"testing"
)
func TestPingPong(t *testing.T)  {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	t.Logf(pong, err)
}
func TestListLen(t *testing.T)  {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	t.Logf("len %s",client.LLen("complete_task"))
}
func TestParseArray(t *testing.T)  {
	jsonByte, _ := ioutil.ReadFile("/Users/chen/GoPath/src/xiaoshiejie.com/redmonitor/config/task.json")
	jsonArr := fastjson.ParseArray(jsonByte[:])
	t.Logf("size %d",jsonArr.Size())

}
func TestJSONObjectGetInt(t *testing.T)  {
	jsonByte, _ := ioutil.ReadFile("/Users/chen/GoPath/src/xiaoshiejie.com/redmonitor/config/task.json")
	jsonArr := fastjson.ParseArray(jsonByte[:])
	jsonObject, e := jsonArr.GetJSONObject(1)
	if e != nil {
		//t.Error("err %s",e.Error())
	}
	t.Log( jsonObject.GetInteger("taskType"))
	t.Log( jsonObject.GetString("queueName"))

}
func TestJSONObjectGetJSONArray(t *testing.T)  {
	jsonByte, _ := ioutil.ReadFile("/Users/chen/GoPath/src/xiaoshiejie.com/redmonitor/config/task.json")
	jsonArr := fastjson.ParseArray(jsonByte[:])
	jsonObject, e := jsonArr.GetJSONObject(1)
	if e != nil {
		//t.Error("err %s",e.Error())
	}
	t.Log( jsonObject.GetString("queueName"))

	jsonArray := jsonObject.GetJSONArray("taskData")
	t.Log(jsonArray.GetString(1))

}