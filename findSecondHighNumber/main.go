package main

import (
	"fmt"
)

func findSecondMax(arr []int) int {
	var max, secondMax int

	for _, num := range arr {
		if num > max {
			secondMax = max
			max = num   
		} else if num > secondMax {
			secondMax = num
		}
	}

	return secondMax
}

func main() {
	salaries := []int{4, 3, 1, 2, 5}

	secondMax := findSecondMax(salaries)

	fmt.Println("The second maximum salary is:", secondMax)
}

/*
db.employees.aggregate([
    {
        "$sort": { "salary": -1 }
    },
    {
        "$skip": 1
    },
    {
        "$limit": 1
    },
    {
        "$project": {
            "_id": 0,
            "second_max_salary": "$salary"
        }
    }
])
*/

// SELECT MAX(salary) AS second_max_salary
// FROM salaries
// WHERE salary < (SELECT MAX(salary) FROM salaries);
