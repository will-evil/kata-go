package helpbookseller

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	catMap := make(map[string]uint)
	for _, art := range listArt {
		artSlice := strings.Split(art, " ")
		num, err := strconv.ParseUint(strings.TrimSpace(artSlice[1]), 10, 64)
		if err != nil {
			continue
		}
		ch := strings.ToUpper(strings.TrimSpace(string([]rune(artSlice[0])[0])))
		catMap[ch] += uint(num)
	}

	var numStr string
	var resSlice []string
	for _, category := range listCat {
		category = strings.ToUpper(strings.TrimSpace(category))
		if val, ok := catMap[category]; ok {
			numStr = fmt.Sprint(val)
		} else {
			numStr = "0"
		}
		resSlice = append(resSlice, fmt.Sprintf("(%s : %s)", category, numStr))
	}

	return strings.Join(resSlice, " - ")
}
