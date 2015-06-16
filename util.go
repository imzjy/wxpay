package wxpay

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
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

// Sign the parameter in form of map[string]string with app key.
// Empty string and "sign" key is excluded before sign.
// Please refer to http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=4_3
func Sign(param map[string]string, key string) string {
	newMap := make(map[string]string)
	// fmt.Printf("%#v\n", param)
	for k, v := range param {
		if k == "sign" {
			continue
		}
		if v == "" {
			continue
		}
		newMap[k] = v
	}
	// fmt.Printf("%#v\n\n", newMap)

	preSignStr := SortAndConcat(newMap)
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
	return fmt.Sprintf("%d", time.Now().Unix()+ChinaTimeZoneOffset)
}
