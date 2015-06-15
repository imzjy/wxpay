package wxpay

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
	"time"
	"strconv"
)

// SortAndConcat sort the map by key in ASCII order,
// and concat it in form of "k1=v1&k2=2"
func SortAndConcat(param map[string]string) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}

	var sortedParam []string
	sort.Strings(keys)
	for _, k := range keys {
		// fmt.Println(k, "=", param[k])
		sortedParam = append(sortedParam, k+"="+param[k])
	}

	return strings.Join(sortedParam, "&")
}

// Sign the string in form of "k1=v1&k2=v2" with app key.
// Please refer to http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=4_3
func Sign(preSignStr, key string) string {
	preSignWithKey := preSignStr + "&key=" + key

	return fmt.Sprintf("%X", md5.Sum([]byte(preSignWithKey)))
}

// NewNonceString return random string in 32 characters
func NewNonceString() string {
	nonce := strconv.FormatInt(time.Now().UnixNano(), 36)
	return fmt.Sprintf("%x", md5.Sum([]byte(nonce)))
}

const ChinaTimeZoneOffset = 8 * 60 * 60 //Beijing(UTC+8:00)

// NewTimestampString return 
func NewTimestampString() string {
	return fmt.Sprintf("%d", time.Now().Unix() + ChinaTimeZoneOffset)
}


// ToXmlString convert the map[string]string to xml string
func ToXmlString(param map[string]string) string {
	xml := "<xml>"
	for k, v := range param {
		xml = xml + fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	xml = xml + "</xml>"

	return xml
}