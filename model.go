package main

type CryptoApi interface {
	Schedule(m *chan int)
}
type AppConfig struct {
	Api []Apiconfig
}
type Apiconfig struct {
	ApiKey string
	Url    string
	Param  map[string]string
}
type CryptoCompareResponse struct {
	Response   string `json:"Response"`
	Message    string `json:"Message"`
	HasWarning bool   `json:"HasWarning"`
	Type       int    `json:"Type"`
	RateLimit  struct {
	} `json:"RateLimit"`
	Data AssetsBundle `json:"Data"`
}

type AssetsBundle struct {
	Aggregated bool     `json:"Aggregated"`
	TimeFrom   int      `json:"TimeFrom"`
	TimeTo     int      `json:"TimeTo"`
	Data       []Assets `json:"Data"`
}

type Assets struct {
	TableName        struct{} `sql:"assets,alias:assets" json:"-"`
	Id               int      `sql:",pk" json:"id"`
	Time             int64    `json:"time"`
	High             float64  `json:"high"`
	Low              float64  `json:"low"`
	Open             float64  `json:"open"`
	Volumefrom       float64  `json:"volumefrom"`
	Volumeto         float64  `json:"volumeto"`
	Close            float64  `json:"close"`
	ConversionType   string   `sql:"conversiontype" json:"conversionType"`
	ConversionSymbol string   `sql:"conversionsymbol" json:"conversionSymbol"`
}
