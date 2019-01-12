package helpbookseller

import "testing"

type TestStockListData struct {
	listArt        []string
	listCat        []string
	expectedResult string
}

func TestStockList(t *testing.T) {
	var dataList = []TestStockListData{
		{
			listArt:        []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"},
			listCat:        []string{"A", "B", "C", "D"},
			expectedResult: "(A : 0) - (B : 1290) - (C : 515) - (D : 600)",
		},
		{
			listArt:        []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"},
			listCat:        []string{"A", "B", "C", "D"},
			expectedResult: "(A : 0) - (B : 1290) - (C : 515) - (D : 600)",
		},
		{
			listArt:        []string{},
			listCat:        []string{"A", "B", "C", "D"},
			expectedResult: "",
		},
		{
			listArt:        []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"},
			listCat:        []string{},
			expectedResult: "",
		},
	}

	for _, data := range dataList {
		if res := StockList(data.listArt, data.listCat); res != data.expectedResult {
			t.Errorf("StockList is failed. Expected %s, got %s", data.expectedResult, res)
		}
	}
}
