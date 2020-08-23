package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"1233454356\tGET /index.html\t10.10.10.1",
		"1233454356\tGET /index.html\t10.10.10.1",
		"1233454356\tGET /index.html\t10.10.10.1",
		"1233454323\tGET /index.htm\t10.13.10.1",
		"1233454356\tGET /browse.html\t10.14.10.1",
		"1232343356\tGET /index.html\t10.14.10.1",
		"1233454356\tGET /index.html\t10.14.10.1",
		"1233654906\tGET /index.html\t123.14.10.1",
		"1233454316\tGET /index.html\t10.14.10.1",
		"1233887346\tGET /index.html\t10.14.10.1",
		"1233454356\tGET /list.html\t10.14.10.1",
		"1233454356\tPOST /update.php\t10.14.11.1",
		"1233454356\tGET /index.html\t10.14.11.1",
		"1233454356\tGET /index.html\t10.14.11.1",
		"1233452356\tGET /index.html\t10.14.11.1",
		"1233454356\tGET /index.html\t10.14.11.1",
		"1233754356\tGET /index.html\t10.14.10.1",
		"1233454356\tGET /index.html\t123.14.10.1"}
	var ip = []string{}
	var m = make(map[string]int)

	for i := range input {
		j := strings.Split(input[i], "\t")
		ip = append(ip, j[len(j)-1])
	}
	m[ip[0]] = 0

	for _, item := range ip {
		for key, value := range m {
			if item == key {
				value++
				fmt.Println(key, "is  ====>", value)
				m[item] = value
			} else {
				if m[item] == 0 {
					m[item] = 1
					fmt.Println("<===", item)
					break
				}
			}
		}

	}
	// Int slice to store values of map.
	values := make([]int, 0, len(m))
	mNew := map[int]string{}

	for key, v := range m {
		values = append(values, v)
		mNew[v] = key
	}

	sort.Ints(values)

	for _, value := range values {
		fmt.Println(mNew[value], "==>", value)
	}

}
