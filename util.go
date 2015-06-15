package wxpay

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
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

// Sign the string with app key.
// Please refer to http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=4_3
func Sign(preSignStr, key string) string {
	preSignWithKey := preSignStr + "&key=" + key

	return fmt.Sprintf("%X", md5.Sum([]byte(preSignWithKey)))
}
