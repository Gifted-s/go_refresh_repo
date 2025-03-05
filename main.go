// Refresh Golang
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

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

type Result struct {
	mutex    sync.Mutex
	storeMap map[string]*MonthlyStoreEarnings
}

func NewMonthlyStoreEarnings(storeId string, orderId string, totalAmount float64) *MonthlyStoreEarnings {
	return &MonthlyStoreEarnings{
		StoreId:     storeId,
		OrderIds:    []string{orderId},
		TotalAmount: totalAmount,
	}
}

func worker2(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker stopped:", ctx.Err())
            return
        default:
            // Do work
            time.Sleep(time.Second)
            fmt.Println("Working...")
        }
    }
}



func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    go worker2(ctx)
    time.Sleep(6 * time.Second)

	
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

func processOrders(routines int, orders []Order) {
	orderChan := make(chan Order)
	wg := sync.WaitGroup{}
	result := Result{storeMap: map[string]*MonthlyStoreEarnings{}}

	for i := 0; i < routines; i++ {
		wg.Add(1)
		go worker(&wg, orderChan, &result)
	}

	for _, o := range orders {
		orderChan <- o
	}
	close(orderChan)

	wg.Wait()
	for _, v := range result.storeMap {
		fmt.Println(v)
	}
}

func worker(wg *sync.WaitGroup, orderChan <-chan Order, result *Result) {
	defer wg.Done()
	for order := range orderChan {
		result.mutex.Lock()
		if earning, exists := result.storeMap[order.StoreId]; exists {
			earning.TotalAmount += order.Amount
			earning.OrderIds = append(earning.OrderIds, order.OrderId)
		} else {
			result.storeMap[order.StoreId] = NewMonthlyStoreEarnings(order.StoreId, order.OrderId, order.Amount)
		}
		result.mutex.Unlock()
	}
}
