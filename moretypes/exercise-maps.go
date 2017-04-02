// https://tour.golang.org/moretypes/23 (2017-04-02)
// Exercise: Maps

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		ret[words[i]] = (ret[words[i]]) + 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
