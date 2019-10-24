package test_time

import (
	"fmt"
	"time"
)

func CycleTask() {
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for _ = range ticker.C {
			fmt.Println(time.Now())
		}
	}()

	time.Sleep(time.Minute)
	fmt.Println("test")
}

func CycleTask1() {
	ticker := time.NewTicker(time.Second * 5)

	<-ticker.C
	fmt.Println("test")
}

func CycleTask2() {
	ticker := time.NewTicker(time.Second * 5)
	fmt.Println(time.Now())
	go func() {
		<-ticker.C
		fmt.Println(time.Now())
	}()

	time.Sleep(time.Second * 10)
	fmt.Println("xixixi", time.Now())
}

func CycleTask3() {
	go func() {
		for {
			fmt.Println(time.Now())
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

func CycleTask5() {
	go func() {
		fmt.Println(time.Now())
		t := time.NewTicker(time.Second * 5)
		for {
			<-t.C
			fmt.Println(time.Now())
			t = time.NewTicker(time.Second * 10)
		}
	}()
}
