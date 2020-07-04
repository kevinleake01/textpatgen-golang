/*
####################################
#
# -- TEXTPATGEN GENDATA TEMPLATE --
#
# Example usage:
# tpl-gendata-tcode-0005.go 256 16 0 255 100
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

  fmt.Printf("####################################\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- TEXTPATGEN GENERATED RANDOM DATA --\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- A generated random data file created from:\n");
  fmt.Printf("# -- tpl-gendata-tcode-0005.go\n");
  fmt.Printf("#\n");
  fmt.Printf("# -- Created: %s\n", now.Format(time.RFC3339));
  fmt.Printf("# -- Epoch Time: %d\n", epoch);
  fmt.Printf("#\n");
  fmt.Printf("# Total number: %d\n",finish);
  fmt.Printf("# Number in a line: %d\n",width);
  fmt.Printf("# Lowest printable number: %d\n",randmin);
  fmt.Printf("# Highest printable number: %d\n",randmax);
  fmt.Printf("# Random number seed: %d\n",randseed);
  fmt.Printf("#\n");
  fmt.Printf("####################################\n");

  for num:=int64(1); num<=finish; num++   {
    for size:=int64(0); size<width-1; size++ {
      if (num == finish) {
        break;
      }
      randnum=random(randmin, randmax);  /* Get the random number */
      fmt.Printf("%02x ", randnum);  /* Print this number */
      num++;
    }
    randnum=random(randmin, randmax);  /* Get the random number */
    fmt.Printf("%02x\n", randnum);  /* Print this number */
  }
}

func random(randmin int64, randmax int64) int64 {
    return rand.Int63n(randmax-randmin) + randmin;
}

