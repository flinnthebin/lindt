package main

import (
  "fmt"
  "io"
  "bufio"
  "os"
  "strings"
)

func main() {
  var prev bool

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    line = strings.TrimRight(line, " \t")

    if line == "" {
      if prev {
        continue
      }
      prev = true
    } else {
      prev = false
    }
    fmt.Println(line)
  }

  if err := scanner.Err(); err != nil {
    if err != io.EOF {
      fmt.Fprintln(os.Stderr, "Error reading standard input", err)
    }
  }
}
