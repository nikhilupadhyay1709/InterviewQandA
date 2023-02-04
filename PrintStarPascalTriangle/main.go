
// GOLANG PROGRAM TO PRINT STAR PASCALS TRIANGLE
package main

// fmt package provides the function to print anything
// Package strconv implements conversions to and from
// string representations of basic data types
import (
   "fmt"
   "strconv"
)

// start function main ()
func main() {

   // declare the variables
   var i, j, rows, num int
   
   // initialize the rows variable
   rows = 7
   
   // Scanln() function scans the input, reads and stores
   //the successive space-separated values into successive arguments
   fmt.Scanln(&rows)
   fmt.Println("GOLANG PROGRAM TO PRINT STAR PASCALS TRIANGLE")
   
   // Use of For Loop
   // This loop starts when i = 0
   // executes till i<rows condition is true
   // post statement is i++
   for i = 0; i < rows; i++ {
      num = 1
      fmt.Printf("%"+strconv.Itoa((rows-i)*2)+"s", "")
      
      // run next loop as for (j=0; j<=i; j++)
      for j = 0; j <= i; j++ {
         fmt.Printf("%4d", num)
         num = num * (i - j) / (j + 1)
      }
      
      fmt.Println()
      // print the result
   }
}