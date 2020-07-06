/*
####################################
#
# -- TEXTPATGEN TEMPLATE --
#
# Example usage:
# tpl-decgen-file-0002.go 0 99 10
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
var separator string = "-";  /* Separator text */

func main() {
  start,_:=strconv.ParseInt(os.Args[1],10,64);  /* Start number */
  finish,_:=strconv.ParseInt(os.Args[2],10,64); /* Finish number */
  width,_:=strconv.ParseInt(os.Args[3],10,64);  /* Numbers in a line */

  now := time.Now();  /* Get the time */
  epoch := now.Unix();  /* Use Epoch Time */
  epochstr := strconv.FormatInt(epoch, 10); /* Convert this to a string */

  fpo,_:=os.Create("00_" + epochstr + ".txt");
  fmt.Fprintf(fpo, "####################################\n");
  fmt.Fprintf(fpo, "#\n");
  fmt.Fprintf(fpo, "# -- TEXTPATGEN GENERATED FILE --\n");
  fmt.Fprintf(fpo, "#\n");
  fmt.Fprintf(fpo, "# -- Created from tpl-decgen-file-0002.go at:\n");
  fmt.Fprintf(fpo, "# -- %s\n", now.Format(time.RFC3339));
  fmt.Fprintf(fpo, "# -- Epoch Time: %d\n", epoch);
  fmt.Fprintf(fpo, "#\n");
  fmt.Fprintf(fpo, "# Start number: %d ( 0x%X Hex, 0%o Octal )\n",start,start,start);
  fmt.Fprintf(fpo, "# Finish number: %d ( 0x%X Hex, 0%o Octal )\n",finish,finish,finish);
  fmt.Fprintf(fpo, "# Numbers in a line: %d\n",width);
  fmt.Fprintf(fpo, "#\n");
  fmt.Fprintf(fpo, "####################################\n");
  numgen(start, finish, size, width, fpo);
  fmt.Fprintf(fpo, "# -- End of file.\n");
  fpo.Close();
}

func numgen(start int64, finish int64, size int64, width int64, fpo *os.File) {
  for num:=start; num<=finish; num++ {
    for size:=int64(0); size<(width - 1); size++ {
      if (num == finish) {
         break;
      }
      fmt.Fprintf(fpo, "%s", message);
      fmt.Fprintf(fpo, "%04d", num);  /* Number base format */
      fmt.Fprintf(fpo, "%s", separator);
      num++;
    }
    fmt.Fprintf(fpo, "%s", message);
    fmt.Fprintf(fpo, "%04d\n", num);  /* Number base format */
  }
}

