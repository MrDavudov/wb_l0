package main

// import (
// 	"database/sql"
// 	"log"
// )

// func main() {
// 	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
	
// }

// type Sub struct {
// 	id                 string   `json:order_uid`
// 	track_number       string   `json:track_number`
// 	entry              string   `json:entry`
// 	delivery           Delivery `json:delivery`
// 	payment            Payment  `json:payment`
// 	items              []Items  `json:items`
// 	locale             string   `json:locale`
// 	internal_signature string   `json:internal_signature`
// 	customer_id        string   `json:customer_id`
// 	delivery_service   string   `json:delivery_service`
// 	shardkey           string   `json:shardkey`
// 	sm_id              string   `json:sm_id`
// 	data_created       string   `json:date_created`
// 	oof_shard          string   `json:oof_shard`
// }

// type Delivery struct {
// 	name    string `json:"name,omitempty"`
// 	phone   string `json:"phone,omitempty"`
// 	zip     string `json:"zip,omitempty"`
// 	city    string `json:"city,omitempty"`
// 	address string `json:"address,omitempty"`
// 	region  string `json:"region,omitempty"`
// 	email   string `json:"email,omitempty"`
// }

// type Payment struct {
// 	transaction   string `json:transaction`
// 	request_id    string `json:request_id`
// 	currency      string `json:currency`
// 	provider      string `json:provider`
// 	amount        string `json:amount`
// 	payment_dt    string `json:payment_dt`
// 	bank          string `json:bank`
// 	delivery_cost string `json:delivery_cost`
// 	goods_total   string `json:goods_total`
// 	custom_fee    string `json:custom_fee`
// }

// type Items[] struct {
// 	chrt_id      string `json:chrt_id`
// 	track_number string `json:chrt_id`
// 	price        string `json:chrt_id`
// 	rid          string `json:chrt_id`
// 	name         string `json:chrt_id`
// 	sale         string `json:chrt_id`
// 	size         string `json:chrt_id`
// 	total_price  string `json:chrt_id`
// 	nm_id        string `json:chrt_id`
// 	brand        string `json:chrt_id`
// 	status       string `json:chrt_id`
// }