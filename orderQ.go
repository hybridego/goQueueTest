package main

import (
	"fmt"
	"sync"
)

type Order struct {
	Id       int64
	User     string
	Item     string
	Location string
}

type orderQ struct {
	orderList []Order
	mutex     *sync.Mutex
	notEmpty  *sync.Cond
	// You can subscribe to this channel to know whether queue is not empty
	NotEmpty chan struct{}
	orderPk  int64
	orderCnt int64
}

func NewOrderQ() *orderQ {
	q := &orderQ{
		orderList: make([]Order, 0),
		mutex:     &sync.Mutex{},
		NotEmpty:  make(chan struct{}, 1),
		orderPk:   0,
		orderCnt:  0,
	}
	q.notEmpty = sync.NewCond(q.mutex)
	return q
}

// Returns the number of elements in queue
func (q *orderQ) Count() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return int(q.orderCnt)
}

func (q *orderQ) notify() {
	if len(q.orderList) > 0 {
		select {
		case q.NotEmpty <- struct{}{}:
		default:
		}
	}
}

// Adds one element at the back of the queue
func (q *orderQ) Add(newOrder Order) (int64, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	id, ok := repeatCheck(q, newOrder.User)
	if ok {
		return id, false
	}

	newOrder.Id = q.orderPk + 1
	q.orderList = append(q.orderList, newOrder)
	q.orderCnt += 1
	q.orderPk += 1

	q.notify()
	if len(q.orderList) == 1 {
		q.notEmpty.Broadcast()
	}
	return q.orderPk, true
}

func repeatCheck(q *orderQ, userName string) (int64, bool) {
	if q.orderCnt > 0 {
		for _, order := range q.orderList {
			if order.User == userName {
				fmt.Println("It's a repeated order. This order is rejected.")
				return order.Id, true
			}
		}
	}
	return -1, false
}

// Pop removes and returns the element from the front of the queue.
// If the queue is empty, it will block
func (q *orderQ) Pop() (Order, bool) {
	var firstOrder Order
	var ok bool
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.orderCnt > 0 {

		firstOrder = q.orderList[0]
		fmt.Println("firstOrder:", firstOrder)
		q.orderList = q.orderList[1:]
		q.orderCnt -= 1
		ok = true
		return firstOrder, ok
	}
	q.notify()
	ok = false

	return firstOrder, ok
}

// Removes one element from the queue
func (q *orderQ) Cancel(userName string) (int64, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.orderCnt > 0 {
		for n, order := range q.orderList {
			if order.User == userName {
				fmt.Println("This Order has been canceled. ", order)
				q.orderList = append(q.orderList[:n], q.orderList[n+1:]...)
				q.orderCnt -= 1
				return order.Id, true
			}
		}
	}
	return -1, false
}

//Issue an order.
func (q *orderQ) MakeOrder(user string, item string, location string) Order {
	return Order{0, user, item, location}
}

//Print all orders.
func (q *orderQ) PrintAllOrders() {
	fmt.Println("=======  OrderList  =======")
	for n, order := range q.orderList {
		fmt.Println(n, ":", order)
	}
	fmt.Println("===========================")
}
