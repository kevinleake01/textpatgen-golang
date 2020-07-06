/*
####################################
#
# -- TEXTPATGEN TEMPLATE --
#
# Example usage:
# tpl-chksum-file-0001.go 255 16 "+"
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
  epochstr := strconv.FormatInt(epoch, 10); /* Convert this to a string */

  fpo,_:=os.Create("00_" + epochstr + ".txt");
  fmt.Fprintf(fpo ,"####################################\n");
  fmt.Fprintf(fpo ,"#\n");
  fmt.Fprintf(fpo ,"# -- TEXTPATGEN GENERATED FILE --\n");
  fmt.Fprintf(fpo ,"#\n");
  fmt.Fprintf(fpo ,"# -- A generated text file created from:\n");
  fmt.Fprintf(fpo ,"# -- tpl-chksum-tcode-0001.go\n");
  fmt.Fprintf(fpo ,"#\n");
  fmt.Fprintf(fpo ,"# -- This file can be very useful in\n");
  fmt.Fprintf(fpo ,"# -- transmission tests.\n");
  fmt.Fprintf(fpo ,"#\n");
  fmt.Fprintf(fpo ,"# -- Created: %s\n", now.Format(time.RFC3339));
  fmt.Fprintf(fpo ,"# -- Epoch Time: %d\n", epoch);
  fmt.Fprintf(fpo ,"#\n");
  fmt.Fprintf(fpo ,"# Total number: %d\n",finish);
  fmt.Fprintf(fpo ,"# Number in a line: %d\n",width);
  fmt.Fprintf(fpo ,"# Text: \"%s\"\n",message);
  fmt.Fprintf(fpo ,"#\n");
  fmt.Fprintf(fpo ,"####################################\n");

  for num=int64(1); num<=finish; num++   {
    for size=int64(0); size<width-1; size++ {
      if (num == finish) {
        break;
      }
      fmt.Fprintf(fpo ,"%s", message);
      num++;
    }
    fmt.Fprintf(fpo ,"%s\n", message);
  }
  fmt.Fprintf(fpo, "# -- End of file.\n");
  fpo.Close();
}

