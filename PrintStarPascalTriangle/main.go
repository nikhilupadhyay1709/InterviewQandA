package main

import (
   "fmt"
   "strconv"
)

func main() {

   var i, j, rows, num int
   
   rows = 7
   
   fmt.Scanln(&rows)
   fmt.Println("GOLANG PROGRAM TO PRINT STAR PASCALS TRIANGLE")
   
   for i = 0; i < rows; i++ {
      num = 1
      fmt.Printf("%"+strconv.Itoa((rows-i)*2)+"s", "")
      
      for j = 0; j <= i; j++ {
         fmt.Printf("%4d", num)
         num = num * (i - j) / (j + 1)
      }
      
      fmt.Println()
   }
}