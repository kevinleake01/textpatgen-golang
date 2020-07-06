/*
####################################
#
# -- TEXTPATGEN GENDATA TEMPLATE --
#
# Example usage:
# tpl-gendata-file-0004.go 100 10 0 65535 100
#
####################################
*/

package main

import (
  "fmt"
  "os"
  "strconv"
  "math/rand"
  "time"
)

var num int64;
var finish int64;
var width int64;
var size int64;
var randnum int64;
var randmin int64;
var randmax int64;
var randseed int64;

func main() {
  finish,_:=strconv.ParseInt(os.Args[1],10,64);
  width,_:=strconv.ParseInt(os.Args[2],10,64);
  randmin,_:=strconv.ParseInt(os.Args[3],10,64);
  randmax,_:=strconv.ParseInt(os.Args[4],10,64);
  randseed,_:=strconv.ParseInt(os.Args[5],10,64);

  rand.Seed(randseed);

  now := time.Now();  /* Get the time */
  epoch := now.Unix();  /* Use Epoch Time */
  epochstr := strconv.FormatInt(epoch, 10); /* Convert this to a string */

  fpo,_:=os.Create("00_" + epochstr + ".txt");
  fmt.Fprintf(fpo, "####################################\n");
  fmt.Fprintf(fpo, "#\n");
  fmt.Fprintf(fpo, "# -- TEXTPATGEN GENERATED RANDOM DATA --\n");
  fmt.Fprintf(fpo, "#\n");
  fmt.Fprintf(fpo, "# -- A generated random data file created from:\n");
  fmt.Fprintf(fpo, "# -- tpl-gendata-file-0004.go\n");
  fmt.Fprintf(fpo, "#\n");
  fmt.Fprintf(fpo, "# -- Created: %s\n", now.Format(time.RFC3339));
  fmt.Fprintf(fpo, "# -- Epoch Time: %d\n", epoch);
  fmt.Fprintf(fpo, "#\n");
  fmt.Fprintf(fpo, "# Total number: %d\n",finish);
  fmt.Fprintf(fpo, "# Number in a line: %d\n",width);
  fmt.Fprintf(fpo, "# Lowest printable number: %d\n",randmin);
  fmt.Fprintf(fpo, "# Highest printable number: %d\n",randmax);
  fmt.Fprintf(fpo, "# Random number seed: %d\n",randseed);
  fmt.Fprintf(fpo, "#\n");
  fmt.Fprintf(fpo, "####################################\n");

  for num:=int64(1); num<=finish; num++   {
    for size:=int64(0); size<width-1; size++ {
      if (num == finish) {
        break;
      }
      randnum=random(randmin, randmax);  /* Get the random number */
      fmt.Fprintf(fpo, "%05d ", randnum);  /* Print this number */
      num++;
    }
    randnum=random(randmin, randmax);  /* Get the random number */
    fmt.Fprintf(fpo, "%05d\n", randnum);  /* Print this number */
  }
  fmt.Fprintf(fpo, "# -- End of file.\n");
  fpo.Close();
}

func random(randmin int64, randmax int64) int64 {
    return rand.Int63n(randmax-randmin) + randmin;
}

