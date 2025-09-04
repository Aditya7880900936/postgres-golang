package models

type Stock struct {
	StockID      int64  `json:"stockid"`
	StockName    string `json:"stockname"`
	StockPrice   float64  `json:"stockprice"`
	StockCompany string `json:"stockcompany"`
}
