package order

import "time"

type Order struct {
	Region        string    `json:"region"`
	Country       string    `json:"country"`
	ItemType      string    `json:"item_type"`
	SalesChannel  string    `json:"sales_channel"`
	OrderPriority string    `json:"order_priority"`
	OrderDate     time.Time `json:"order_date"`
	OrderID       uint      `json:"order_id"`
	ShipDate      time.Time `json:"ship_date"`
	UnitsSold     uint      `json:"units_sold"`
	UnitPrice     float64   `json:"unit_price"`
	UnitCost      float64   `json:"unit_cost"`
	TotalRevenue  float64   `json:"total_revenue"`
	TotalCost     float64   `json:"total_cost"`
	TotalProfit   float64   `json:"total_profit"`
}
