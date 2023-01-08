package main

import "fmt"

type Teacher struct {
	teacher1, teacher2 int
}
type Student struct {
	Teacher
	student1, student2 float32
}

func main() {
	output := Student{Teacher{1, 2}, 3.0, 4.0}
	fmt.Println(output.teacher1, output.teacher2, output.student1, output.student2)
	fmt.Println(output.Teacher)
}
