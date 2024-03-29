package json

import (
	"encoding/json"
	"log"
)

func ToString(i interface{}) string {
	if i == nil {
		return ""
	}
	b, err := json.Marshal(i)
	if err != nil {
		log.Println("json.ToString:", err)
	}
	return string(b)
}
func ToBytes(i interface{}) (bs []byte) {
	if i == nil {
		return
	}
	b, err := json.Marshal(i)
	if err != nil {
		log.Println("json.ToBytes:", err)
	}
	return b
}

func ToObject(s string, i interface{}) {
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		log.Println("json.ToObject:", err)
	}
}

func ToObjectByBytes(b []byte, i interface{}) {
	err := json.Unmarshal(b, &i)
	if err != nil {
		log.Println("json.ToObjectByBytes:", string(b), err)
	}
}
