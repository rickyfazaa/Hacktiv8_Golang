package views

type OrderGetAll struct {
	OrderId      int          `json:"orderId"`
	OrderedAt    string       `json:"orderedAt"`
	CustomerName string       `json:"customerName"`
	Items        []ItemGetAll `json:"items"`
}
