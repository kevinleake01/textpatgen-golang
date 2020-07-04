/*
####################################
#
# --- MAKE-GO-TPL-0002.C ---
#
# This converts a standard text file
# into a Textpatgen Go Template file.
#
# The program has been written in C because
# of implementation difficulties in the Go
# progamming language.
#
####################################
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

int main(int argc, char *argv[]);

int c; /* Character */ 
FILE *fp;  /* File pointer */

int main(int argc, char *argv[])
{
  /* Timestamp info */
  time_t timer;
  char buffer1[32], buffer2[32];
  struct tm* tm_info;

  time(&timer);  /* get the time */
  tm_info=localtime(&timer);

  strftime(buffer1, 32, "%Y-%m-%dT%H:%M:%S%z", tm_info); /* ISO 8601 Timestamp to print */
  strftime(buffer2, 32, "%s", tm_info); /* Unix Epoch Timestamp to print */

  fp=fopen(argv[1],"r"); /* Open file */
  
  printf("/*\n");
  printf("####################################\n");
  printf("#\n");
  printf("# --- MAKE-G0-TPL-0002.GO ---\n");
  printf("#\n");
  printf("# This file named %s has\n",argv[1]);
  printf("# been converted to a Textpatgen Go\n");
  printf("# Template file.\n");
  printf("#\n");
  printf("# -- ISO 8601 Time: %s\n", buffer1);
  printf("# -- Epoch Time: %s\n", buffer2);
  printf("#\n");
  printf("####################################\n");
  printf("*/\n\n");
  printf("package main\n\n");
  printf("import (\n");
  printf("  \"fmt\"\n");
  printf(")\n\n");
  printf("func main() {\n");
  printf("  fmt.Printf(\""); /* Open first fmt.Printf statement */

  while (c != EOF) /* Loop until the end of the file */
  {  
    c=fgetc(fp); /* Get a character from the file */
    if (c == '\\') /* Backslash */
    { 
      printf("\\\\");
    }
    else if (c == '%') /* Percent sign */
    {
      printf("%%%%");
    }
    else if (c == '\'') /* Single quote */
    {
      printf("\\u0027");
    }
    else if (c == '\"') /* Double quotes */
    {
      printf("\\u0022");
    }
    else if (c == '\?') /* Question mark */
    {
      printf("\\u003F");
    }
    else if (c == '\t') /* Tabs */
    {
      printf("\\t");
    }
    else if (c == '\f') /* Form feeds */
    {
      printf("\\f");
    }
    else if (c == '\a') /* Bell */
    {
      printf("\\a");
    }
    else if (c == '\b') /* Backspace */
    {
      printf("\\b");
    }
    else if (c == '\v') /* Vertical tabs */
    {
      printf("\\v");
    }
    else if (c >= 127) /* Characters 127 to 255 */
    {
      printf("\\x%02X", c);
    }
    else if (c == '\n') /* Newline */
    {
      printf("\\n\");\n"); /* Close printf statement */
      printf("  fmt.Printf(\""); /* Open next fmt.Printf statement */

    }
    else if (c == EOF) /* End of file */
    {
      printf("\\n\");\n"); /* Close final printf statement and exit */
      break;
    }
    else putchar(c);
  }
  printf("}\n\n");
  fclose(fp);
  return 0;
}
