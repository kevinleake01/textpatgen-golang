/*
####################################
#
# -- TEXTPATGEN TEMPLATE --
#
# Example usage:
# tpl-chksum-tcode-0002.go 255 16 "+"
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

var num, finish, width, size int64;
var message string;

func main() {

  finish,_:=strconv.ParseInt(os.Args[1],10,64); /* Finish number */
  width,_:=strconv.ParseInt(os.Args[2],10,64);  /* Numbers in a line */
  message:=string(os.Args[3]);

  now := time.Now();  /* Get the time */
  epoch := now.Unix();  /* Use Epoch Time */

  fmt.Printf("####################################\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- TEXTPATGEN GENERATED FILE --\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- A generated text file created from:\n");
  fmt.Printf("# -- tpl-chksum-tcode-0002.go\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- This file can be very useful in\n");
  fmt.Printf("# -- transmission tests.\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- Created: %s\n", now.Format(time.RFC3339));
  fmt.Printf("# -- Epoch Time: %d\n", epoch);
  fmt.Printf("#\n");
  fmt.Printf("# Total number: %d\n",finish);
  fmt.Printf("# Number in a line: %d\n",width);
  fmt.Printf("# Text: \"%s\"\n",message);
  fmt.Printf("#\n");
  fmt.Printf("####################################\n");

  for num=int64(1); num<=finish; num++   {
    for size=int64(0); size<width-1; size++ {
      if (num == finish) {
        break;
      }
      fmt.Printf("%s", message);
      num++;
    }
    fmt.Printf("%s\n", message);
  }
}

