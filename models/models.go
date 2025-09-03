package models

type Stock struct {
	StockID      int64  `json:"stockid"`
	StockName    string `json:"stockname"`
	StockPrice   int64  `json:"stockprice"`
	StockCompany string `json:"stockcompany"`
}
