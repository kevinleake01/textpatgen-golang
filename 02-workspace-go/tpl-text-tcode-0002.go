/*
####################################
#
# -- TEXTPATGEN TEMPLATE --
#
# Example usage:
# tpl-text-tcode-0002.go 0 25 "The_Quick_Brown_Fox_Jumps_over_the_Lazy_Dog"
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

var num, start, finish int64;
var message string;

func main() {

  start,_:=strconv.ParseInt(os.Args[1],10,64); /* Finish number */
  finish,_:=strconv.ParseInt(os.Args[2],10,64);  /* Numbers in a line */
  message:=string(os.Args[3]);

  now := time.Now();  /* Get the time */
  epoch := now.Unix();  /* Use Epoch Time */

  fmt.Printf("####################################\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- TEXTPATGEN GENERATED FILE --\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- A generated text file created from:\n");
  fmt.Printf("# -- tpl-text-tcode-0002.go\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- Created: %s\n", now.Format(time.RFC3339));
  fmt.Printf("# -- Epoch Time: %d\n", epoch);
  fmt.Printf("#\n");
  fmt.Printf("# Start number: %d\n",start);
  fmt.Printf("# Finish number: %d\n",finish);
  fmt.Printf("# Message text: \"%s\"\n",message);
  fmt.Printf("#\n");
  fmt.Printf("####################################\n");

  for num=start; num<=finish; num++ {
    fmt.Printf("X%04X:%s:D%04d\n", num, message, num);
  }
  fmt.Printf("# -- End of file.\n");
}

