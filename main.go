package main

type Order struct {
	OrderId string
	StoreId string
	Amount  float64
}
type MonthlyStoreEarnings struct {
	StoreId     string
	OrderIds    []string
	TotalAmount float64
}

func main() {
	orders := []Order{
		{
			OrderId: "orderA-storeA",
			StoreId: "storeA",
			Amount:  300.30,
		},
		{
			OrderId: "orderB-storeA",
			StoreId: "storeA",
			Amount:  300.30,
		},
		{
			OrderId: "orderA-storeB",
			StoreId: "storeB",
			Amount:  300.30,
		},
		{
			OrderId: "orderB-storeB",
			StoreId: "storeB",
			Amount:  300.30,
		},
		{
			OrderId: "orderA-storeC",
			StoreId: "storeC",
			Amount:  300.30,
		},
		{
			OrderId: "orderB-storeC",
			StoreId: "storeC",
			Amount:  300.30,
		},
		{
			OrderId: "orderA-storeD",
			StoreId: "storeD",
			Amount:  300.30,
		},
		{
			OrderId: "orderB-storeD",
			StoreId: "storeD",
			Amount:  300.30,
		}, {
			OrderId: "orderC-storeD",
			StoreId: "storeD",
			Amount:  300.30,
		},
	}
	routines := 3
	processOrders(routines, orders)
}

func processOrders(concurrency int, orders []Order) {

}
