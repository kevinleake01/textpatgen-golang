/*
####################################
#
# -- TEXTPATGEN TEMPLATE --
#
# Example usage:
# tpl-hexgen-tcode-0008.go 0 255 16
#
####################################
*/

package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
)

var size int64;
var message string = ""; /* Message text */
var separator string = "";  /* Separator text */

func main() {
  start,_:=strconv.ParseInt(os.Args[1],10,64);  /* Start number */
  finish,_:=strconv.ParseInt(os.Args[2],10,64); /* Finish number */
  width,_:=strconv.ParseInt(os.Args[3],10,64);  /* Numbers in a line */

  now := time.Now();  /* Get the time */
  epoch := now.Unix();  /* Use Epoch Time */

  fmt.Printf("####################################\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- TEXTPATGEN GENERATED FILE --\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- Created from tpl-hexgen-tcode-0008.go at:\n");
  fmt.Printf("# -- %s\n", now.Format(time.RFC3339));
  fmt.Printf("# -- Epoch Time: %d\n", epoch);
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
      fmt.Printf("%02X", num);  /* Number base format */
      fmt.Printf("%s", separator);
      num++;
    }
    fmt.Printf("%s", message);
    fmt.Printf("%02X\n", num);  /* Number base format */
  }
}

