package api

//
//import (
//	"crypto/md5"
//	"crypto/sha1"
//	"encoding/hex"
//	"encoding/json"
//	"fmt"
//	"math/rand"
//	"net/url"
//	"sort"
//	"strings"
//)
//
//func sha1Hash(input string) string {
//	hash := sha1.Sum([]byte(input))
//	return hex.EncodeToString(hash[:])
//}
//
//func md5Hash(input string) string {
//	hash := md5.Sum([]byte(input))
//	return hex.EncodeToString(hash[:])
//}
//
//func getRandom(length int) string {
//	b := make([]byte, length)
//	for i := range b {
//		b[i] = 'a' + byte(rand.Intn(26))
//	}
//	return string(b)
//}
//
//func getDid() string {
//	return md5Hash(getRandom(16))
//}
//
//func getSignText(input string) string {
//	params := strings.Split(input, "&")
//	sortedParams := make([]string, len(params))
//	for i, param := range params {
//		keyValue := strings.Split(param, "=")
//		sortedParams[i] = url.QueryEscape(keyValue[0]) + "=" + url.QueryEscape(keyValue[1])
//	}
//	sort.Strings(sortedParams)
//	return strings.Join(sortedParams, "&")
//}
//
//func niceSignV3(jsonData string, did string, random string) string {
//	result := url.Values{}
//	result.Set("data", "")
//	result.Set("random", random)
//	result.Set("did", did)
//
//	md5List := []string{}
//	md5List = append(md5List, md5Hash(result.Get("did")+" "+result.Get("did")))
//	md5List = append(md5List, md5Hash(random+md5List[0]+" 8a5f746c1c9c99c0b458e1ed510845e5"))
//	md5List[1] = md5List[1][:16] + md5List[1][24:]
//
//	obj := map[string]string{}
//	if jsonData != "" {
//		if err := json.Unmarshal([]byte(jsonData), &obj); err != nil {
//			//log.Fatal(err)
//		}
//	}
//
//	signText := getSignText(url.ValuesToMap(obj).Encode())
//	//signText := ""
//	return signText
//}
//
//func main() {
//	jsonData := `{"key1":"value1","key2":"value2"}`
//	did := getDid()
//	random := getRandom(8)
//	signedData := niceSignV3(jsonData, did, random)
//	fmt.Println(signedData)
//}
