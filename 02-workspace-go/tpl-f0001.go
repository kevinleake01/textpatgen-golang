/*
####################################
#
# --- TEXTPATGEN TEMPLATE ---
#
# Users can change the output by editing
# this file directly.
#
# The text is written to a named file.
#
####################################
*/

package main

import (
  "fmt"
  "os"
)

var size int64;
var fpo *os.File;
var message string = "X-"; /* Message text */
var separator string = " ";  /* Separator text */

func main() {
  start:=int64(0);  /* Start number */
  finish:=int64(255); /* Finish number */
  width:=int64(16);  /* Numbers in a line */

  fpo,_:=os.Create("file0001.txt");
  fmt.Fprintf(fpo, "####################################\n");
  fmt.Fprintf(fpo, "#\n");
  fmt.Fprintf(fpo, "# -- TEXTPATGEN GENERATED FILE --\n");
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
      fmt.Fprintf(fpo, "%04X", num);  /* Number base format */
      fmt.Fprintf(fpo, "%s", separator);
      num++;
    }
    fmt.Fprintf(fpo, "%s", message);
    fmt.Fprintf(fpo, "%04X\n", num);  /* Number base format */
  }
}
