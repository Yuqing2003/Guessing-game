//package main

//func main() {
//	好康的 := 欢迎来我家玩()
//	打电动()
//	fmt.Println(好康的)
//}
//func 欢迎来我家玩() string {
//	time.Sleep(5 * time.Second)
//	return "登dua郎
//"
//}

//	func 打电动() {
//		fmt.Println("输了啦，都是你害的")
//
// package main
//
// import (
//
//	"fmt"
//	"time"
//
// )
//
// func main() {
// var (
// ch1  = make(chan struct{})
// stop = make(chan struct{})
// )
// go handleCh1(ch1)
// go func() {
// time.Sleep(10 * time.Second)
// stop <- struct{}{}
// }()
// LOOP:
// for {
// select {
// case _ = <-ch1:
// fmt.Println("get from ch1")
// case _ = <-stop:
// break LOOP
// }
// }
// }
//
// func handleCh1(ch1 chan struct{}) {
// for {
// time.Sleep(3 * time.Second)
// ch1 <- struct{}{}
// }
// }
package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		ch1  = make(chan struct{})
		stop = make(chan struct{})
	)
	go handleCh1(ch1)
	go func() {
		time.Sleep(10 * time.Second)
		stop <- struct{}{}
	}()
LOOP:
	for {
		select {
		case _ = <-ch1:
			fmt.Println("get from ch1")
		case _ = <-stop:
			break LOOP
		}
	}
}

func handleCh1(ch1 chan struct{}) {
	for {
		time.Sleep(3 * time.Second)
		ch1 <- struct{}{}
	}
}
