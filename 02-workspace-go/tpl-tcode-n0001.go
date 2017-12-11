/*
####################################
#
# --- TEXTPATGEN TEMPLATE ---
#
# Users can change the output by editing
# this file directly.
#
####################################
*/

package main

import (
  "fmt"
  "time"
)

var size int64;
var message string = "X-"; /* Message text */
var separator string = " ";  /* Separator text */

func main() {
  start:=int64(0);  /* Start number */
  finish:=int64(255); /* Finish number */
  width:=int64(16);  /* Numbers in a line */

  now := time.Now();  /* Get the time */

  fmt.Printf("####################################\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- TEXTPATGEN GENERATED FILE --\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- Created: %s\n", now.Format(time.RFC3339));
  fmt.Printf("#\n");
  fmt.Printf("# Start number: %d ( 0x%X Hex, 0%o Octal )\n",start,start,start);
  fmt.Printf("# Finish number: %d ( 0x%X Hex, 0%o Octal )\n",finish,finish,finish);
  fmt.Printf("# Numbers in a line: %d\n",width);
  fmt.Printf("#\n");
  fmt.Printf("####################################\n");
  numgen(start, finish, size, width);
  fmt.Printf("# -- End of file.\n");
}

func numgen(start int64, finish int64, size int64, width int64) {
  for num:=start; num<=finish; num++ {
    for size:=int64(0); size<(width - 1); size++ {
      if (num == finish) {
         break;
      }
      fmt.Printf("%s", message);
      fmt.Printf("%04X", num);  /* Number base format */
      fmt.Printf("%s", separator);
      num++;
    }
    fmt.Printf("%s", message);
    fmt.Printf("%04X\n", num);  /* Number base format */
  }
}

