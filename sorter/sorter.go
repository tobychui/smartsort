package sorter

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func SortString(array []string) []string {
	bufKey := []string{}
	bufMap := map[string]string{}
	for _, thisString := range array {
		re := regexp.MustCompile("[0-9]+")
		matchings := re.FindAllString(thisString, -1)
		cs := thisString
		for _, matchPoint := range matchings {
			mpi, _ := strconv.Atoi(matchPoint)
			replacement := fmt.Sprintf("%018d", mpi)
			cs = strings.ReplaceAll(cs, matchPoint, replacement)
		}

		bufKey = append(bufKey, cs)
		bufMap[cs] = thisString
	}

	sort.Strings(bufKey)

	result := []string{}
	for _, key := range bufKey {
		result = append(result, bufMap[key])
	}
	return result
}
