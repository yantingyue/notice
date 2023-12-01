package api

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dop251/goja"
	"log"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func Sign(jsonData, did, random string) string {
	if did == "" {
		did = getDid()
	}
	if random == "" {
		random = getRandom(16)
	}
	md5List := make([]string, 0)
	a := MD5(did[16:] + did[:16])
	b := MD5(random + a + "8a5f746c1c9c99c0b458e1ed510845e5")
	b = b[16:] + b[:16]
	md5List = append(md5List, a)
	md5List = append(md5List, b)
	f := formatData(jsonData)
	st := signText(f)
	sn := Sha1(st + b)
	return fmt.Sprintf("nice-sign-v1://%s%s:%s/%s", sn[24:], sn[8:24], random, jsonData)
}

func formatData(s string) string {
	js := goja.New()
	_, err := js.RunString(`
 function deal_array(params) {
   let result = [];
   for (let i = 0; i < params.length; i++) {
       value = params[i];
       if (["{}", "[]", ""].indexOf(value.toString()) != -1) {
           value = "";
       }
       if (Array.isArray(value)) {
           value = deal_array(value);
       } else if (typeof value == "object") {
           value = deal_dict(value);
       }
       result.push(i+"="+value);
   }
   return result.join("&");
}

function deal_dict(params) {
   let result = [],
       value;
   for (let key of Object.keys(params).sort()) {
       value = params[key];
       if (["{}", "[]", ""].indexOf(value.toString()) != -1) {
           value = "";
       }
       if (Array.isArray(value)) {
           value = deal_array(value);
       } else if (typeof value == "object") {
           value = deal_dict(value);
       }
       result.push(key+"="+value);
   }
   return result.join("&");
}

function deal(JsonData) {
let obj = JsonData === "" ? {} : typeof JsonData === "string" ? JSON.parse(JsonData) : JsonData
return deal_dict(obj)
}
`)
	if err != nil {
		panic(err)
	}
	var f func(string) string
	if err = js.ExportTo(js.Get("deal"), &f); err != nil {
		panic(err)
	}
	return f(s)
}

func dealDict(s string, result []string) {
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(s), &m)
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		val := m[key]
		v := reflect.ValueOf(val)
		if v.IsNil() || v.IsZero() {
			result = append(result, fmt.Sprintf("%s=", key))
		} else if v.Kind() == reflect.Array {
			dealArray(v.String(), result)
		} else if v.Kind() == reflect.Struct || v.Kind() == reflect.Ptr {
			dealDict(v.String(), result)
		} else {
			result = append(result, fmt.Sprintf("%s=%+v", key, val))
		}
	}
}

func dealArray(s string, result []string) {
	arr := make([]interface{}, 0)
	_ = json.Unmarshal([]byte(s), &arr)
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		v := reflect.ValueOf(val)
		if v.IsNil() || v.IsZero() {
			//result = append(result, fmt.Sprintf("%s="))
		}
		if v.Kind() == reflect.Array {
			dealArray(v.String(), result)
		} else if v.Kind() == reflect.Struct || v.Kind() == reflect.Ptr {
			dealDict(v.String(), result)
		}
	}
}

func MD5(text string) string {
	m := md5.New()
	m.Write([]byte(text))
	return hex.EncodeToString(m.Sum(nil))
}

func Sha1(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func GetRandom() {
	log.Println(getRandom(11))
}

func getRandom(n int) string {
	bytes := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	l := len(letters)
	for i := 0; i < n; i++ {
		bytes[i] = letters[rand.Intn(l)]
	}
	return string(bytes)
}

func GetDid() {
	log.Println(getDid())
}

func getDid() string {
	return MD5(getRandom(16))
}

func SignText() {
	log.Println(signText("asuifkijwoeruwo983483849129038e"))
}

func signText(text string) string {
	output := make([]byte, 0)
	outputIdx := 4
	for len(text)>>1 > outputIdx-4 {
		c1 := text[0]
		c2 := text[1]
		output = append(output, (c1&0xF0)|(c2&0xF))
		text = text[2:]
	}
	return string(output)
}
