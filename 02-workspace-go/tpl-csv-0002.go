/*
####################################
#
# -- TEXTPATGEN TEMPLATE --
#
# Example usage:
# tpl-csv-0002.go "0,255,16"
#
####################################
*/

package main

import (
  "fmt"
  "os"
)

var num, start,finish, width, size int64;
var extinfo string;

func main() {
  extinfo:=string(os.Args[1]);  /* Get the extra information */
  fmt.Sscanf(extinfo, "%d,%d,%d", &start, &finish, &width );  /* Process this information */
  numgen(start, finish, size, width);
}

func numgen(start int64, finish int64, size int64, width int64) {
  for num:=start; num<=finish; num++ {
    for size:=int64(0); size<(width - 1); size++ {
      if (num == finish) {
         break;
      }
      fmt.Printf("\"D%04d\",", num);
      num++;
    }
    fmt.Printf("\"D%04d\"\n", num);
  }
}

