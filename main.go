package main

import "fmt"

func main() {
	oq := NewOrderQ()

	printCnt := func() { fmt.Println("Cur order count : ", oq.Count()) }
	printCnt()

	oq.Add(oq.MakeOrder("l.j_1", "비타100", "Y3-1"))
	printCnt()

	oq.Add(oq.MakeOrder("l.j_1", "비타100", "Y3-100"))
	printCnt()

	oq.Add(oq.MakeOrder("l.j_2", "비타200", "Y3-2"))
	printCnt()

	firstOrder, _ := oq.Pop()
	fmt.Println("\nPoped order :", firstOrder)
	printCnt()
	fmt.Println()

	oq.Add(oq.MakeOrder("l.j_3", "비타300", "Y3-3"))
	printCnt()
	oq.Add(oq.MakeOrder("l.j_4", "비타400", "Y3-4"))
	printCnt()
	oq.Add(oq.MakeOrder("l.j_5", "비타500", "Y3-5"))
	printCnt()
	oq.Add(oq.MakeOrder("l.j_6", "비타600", "Y3-6"))
	printCnt()
	oq.Add(oq.MakeOrder("l.j_7", "비타700", "Y3-7"))
	printCnt()

	firstOrder, _ = oq.Pop()
	fmt.Println("\nPoped order :", firstOrder)
	printCnt()
	fmt.Println()

	fmt.Println("\nCancel order :")
	oq.PrintAllOrders()
	_, _ = oq.Cancel("l.j_5")
	oq.PrintAllOrders()
	printCnt()

	fmt.Println("\nCancel order :")
	oq.PrintAllOrders()
	_, _ = oq.Cancel("l.j_3")
	oq.PrintAllOrders()
	printCnt()

	fmt.Println("\nCancel order :")
	oq.PrintAllOrders()
	_, _ = oq.Cancel("l.j_6")
	oq.PrintAllOrders()
	printCnt()

	fmt.Println("\nCancel order :")
	oq.PrintAllOrders()
	_, _ = oq.Cancel("l.j_7")
	oq.PrintAllOrders()
	printCnt()

	fmt.Println("\nCancel order :")
	oq.PrintAllOrders()
	_, _ = oq.Cancel("l.j_4")
	oq.PrintAllOrders()
	printCnt()

}
