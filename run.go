// // Refresh Golang
// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"sync"
// 	//"time"
// )

// type Order struct {
// 	OrderId string
// 	StoreId string
// 	Amount  float64
// }
// type MonthlyStoreEarnings struct {
// 	StoreId     string
// 	OrderIds    []string
// 	TotalAmount float64
// }

// type Result struct {
// 	m        sync.RWMutex
// 	storeMap map[string]*MonthlyStoreEarnings
// }

// // Implement a solution that will group the stores together based on StoreID, aggregate their orderId and sum up the total amount
//    var bufferPool = &sync.Pool{
// 	New: func() interface{} {
// 		return new(bytes.Buffer)
// 	},
//     }
// func runner() {

//     buffer := bufferPool.Get().(*bytes.Buffer)

// 	// put it back in the pool
// 	defer bufferPool.Put(buffer)
// 	// use buffer
//     buffer.Write([]byte{'A', 'D', 'E', 'W', 'U', 'M', 'I'})
// 	// reset buffer
// 	buffer.Reset()

// 	orders := []Order{
// 		{
// 			OrderId: "orderA-storeA",
// 			StoreId: "storeA",
// 			Amount:  300.30,
// 		},
// 		{
// 			OrderId: "orderB-storeA",
// 			StoreId: "storeA",
// 			Amount:  300.30,
// 		},
// 		{
// 			OrderId: "orderA-storeB",
// 			StoreId: "storeB",
// 			Amount:  300.30,
// 		},
// 		{
// 			OrderId: "orderB-storeB",
// 			StoreId: "storeB",
// 			Amount:  300.30,
// 		},
// 		{
// 			OrderId: "orderA-storeC",
// 			StoreId: "storeC",
// 			Amount:  300.30,
// 		},
// 		{
// 			OrderId: "orderB-storeC",
// 			StoreId: "storeC",
// 			Amount:  300.30,
// 		},
// 		{
// 			OrderId: "orderA-storeD",
// 			StoreId: "storeD",
// 			Amount:  300.30,
// 		},
// 		{
// 			OrderId: "orderB-storeD",
// 			StoreId: "storeD",
// 			Amount:  300.30,
// 		}, {
// 			OrderId: "orderC-storeD",
// 			StoreId: "storeD",
// 			Amount:  300.30,
// 		},
// 	}
// 	routines := 3
// 	processOrders(routines, orders)

// 	// chan1 := make(chan string)
// 	// chan2 := make(chan string)

// 	// go func() {
// 	// 	for {
// 	// 		time.Sleep(time.Millisecond * 500)
// 	// 		chan1 <- "Msg 1"
// 	// 	}
// 	// }()

// 	// go func() {
// 	// 	for {
// 	// 		time.Sleep(time.Second * 2)
// 	// 		chan2 <- "Msg 2"
// 	// 	}
// 	// }()

// 	// for {
// 	// 	select {
// 	// 	case msg := <-chan1:
// 	// 		{
// 	// 			fmt.Println(msg)
// 	// 		}
// 	// 	case msg := <-chan2:
// 	// 		{
// 	// 			fmt.Println(msg)
// 	// 		}
// 	// 	}

// 	// }

// 	// Test shapes
// 	c := &Circle{radius: 30.0}
// 	r := &Rectangle{length: 4.5, width: 5.0}
// 	s := &Square{length: 6.0}

// 	shapes := []Shape{c, r, s}

// 	for _, s := range shapes {
// 		fmt.Printf("Shape %s has area = %f \n", s.Name(), s.Area())
// 	}
// }

// func processOrders(routines int, orders []Order) {
// 	wg := sync.WaitGroup{}
// 	result := Result{storeMap: map[string]*MonthlyStoreEarnings{}}
// 	orderChan := make(chan Order)

// 	for i := 0; i < routines; i++ {
// 		wg.Add(1)
// 		go worker(&wg, orderChan, &result)
// 	}
// 	go func() {
// 		for _, o := range orders {
// 			orderChan <- o
// 		}
// 		close(orderChan)
// 	}()
// 	wg.Wait()
// 	for _, v := range result.storeMap {
// 		fmt.Println(v)
// 	}

// }

// func worker(wg *sync.WaitGroup, orderChan <-chan Order, result *Result) {
// 	defer wg.Done()
// 	for order := range orderChan {
// 		result.m.Lock()
// 		if earning, ok := result.storeMap[order.StoreId]; ok {
// 			earning.TotalAmount += order.Amount
// 			earning.OrderIds = append(earning.OrderIds, order.OrderId)
// 		} else {
// 			result.storeMap[order.StoreId] = &MonthlyStoreEarnings{StoreId: order.StoreId, OrderIds: []string{order.OrderId}, TotalAmount: order.Amount}
// 		}
// 		result.m.Unlock()
// 	}
// }

// type Shape interface {

// 	// Area of the shape
// 	Area() float32

// 	// Name of the shape
// 	Name() string
// }

// type Rectangle struct {
// 	length float32
// 	width  float32
// }

// func (rectangle *Rectangle) Area() float32 {
// 	return rectangle.length * rectangle.width
// }
// func (rectangle *Rectangle) Name() string {
// 	return "Rectangle"
// }

// type Square struct {
// 	length float32
// }

// func (s *Square) Area() float32 {
// 	return s.length * s.length
// }
// func (s *Square) Name() string {
// 	return "Square"
// }

// type Circle struct {
// 	radius float32
// }

// func (c *Circle) Area() float32 {
// 	return c.radius * c.radius * 3.142
// }
// func (c *Circle) Name() string {
// 	return "Circle"
// }

// type ShapeManager struct{}

// func (shapeManager *ShapeManager) CalculateArea(shape Shape) {
// 	fmt.Printf("Area of %s is %f", shape.Name(), shape.Area())
// }

// // func processOrders(concurrency int, orders []Order) {
// // 	result := make(map[string]MonthlyStoreEarnings)
// // 	ordersChan := make(chan Order)
// // 	done := make(chan struct{})
// // 	var mu sync.Mutex
// // 	go func() {
// // 		for i := 0; i < len(orders); i++ {
// // 			ordersChan <- orders[i]
// // 		}
// // 		done <- struct{}{}
// // 	}()

// // 	for i := 0; i < concurrency; i++ {
// // 		go func(mu *sync.Mutex) {
// // 			worker(mu, ordersChan, result)
// // 		}(&mu)
// // 	}
// // 	<-done
// // 	for _, v := range result {
// // 		fmt.Println(v)
// // 	}
// // }

// // func worker(mu *sync.Mutex, orders chan Order, result map[string]MonthlyStoreEarnings) {
// // 	for order := range orders {
// // 		mu.Lock()
// // 		if earning, ok := result[order.StoreId]; ok {
// // 			earning.TotalAmount += order.Amount
// // 			earning.OrderIds = append(result[order.StoreId].OrderIds, order.OrderId)
// // 			result[order.StoreId] = earning
// // 		} else {
// // 			result[order.StoreId] = MonthlyStoreEarnings{
// // 				StoreId:     order.StoreId,
// // 				OrderIds:    []string{order.OrderId},
// // 				TotalAmount: order.Amount,
// // 			}
// // 		}
// // 		mu.Unlock()
// // 	}
// // }
