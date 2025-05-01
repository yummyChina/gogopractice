package main

import (
	"fmt"
	// "time"

	part1 "practice.com/task2/part1"
	// part2 "practice.com/task2/part2"
	part3 "practice.com/task2/part3"
	part5 "practice.com/task2/part5"
)

func main() {
	a := 10
	part1.AddNum(&a)
	fmt.Println("计算后的值:", a)

	before := []int{1, 2, 3}
	part1.RecevieSlice(&before)
	fmt.Println("计算后的值:", before)

	// part2.TwoRoutine()

	// scheduler := part2.NewScheduler()

	// scheduler.AddTask(part2.Task{
	// 	ID:   1,
	// 	Name: "任务一",
	// 	Job: func() {
	// 		time.Sleep(1 * time.Second)
	// 	},
	// })

	// scheduler.AddTask(part2.Task{
	// 	ID:   2,
	// 	Name: "任务二",
	// 	Job: func() {
	// 		time.Sleep(2 * time.Second)
	// 	},
	// })

	// scheduler.AddTask(part2.Task{
	// 	ID:   3,
	// 	Name: "任务三",
	// 	Job: func() {
	// 		time.Sleep(3 * time.Second)
	// 	},
	// })

	// scheduler.Start()

	// scheduler.PrintStats()

	// rec := part3.Rectangle{
	// 	Width:  10,
	// 	Height: 5,
	// }
	// fmt.Println("矩形的面积:", rec.Area())
	// fmt.Println("矩形的周长:", rec.Perimeter())

	// cir := part3.Circle{
	// 	Radius: 5,
	// }
	// fmt.Println("圆的面积:", cir.Area())
	// fmt.Println("圆的周长:", cir.Perimeter())

	emp := part3.Employee{
		Person: part3.Person{
			Name: "John Doe",
			Age:  30,
		},
		EmployeeId: "E12345",
	}

	emp.PrintInfo()

	// part4.Comu()
	// part4.Comu1()

	// part5.DoCount()
	part5.DoCount2()

}
