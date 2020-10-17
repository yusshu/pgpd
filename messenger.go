package main

import "fmt"

const RESET = "\033[0m"
const RED = "\033[31m"
const GREEN = "\033[32m"
const YELLOW = "\033[33m"

func infof(format string, params ...interface{}) {
  format = "&a[INFO] &r" + format + "&r"
  format = colorize(format)
  if len(params) == 0 {
    fmt.Println(format)
  } else {
    fmt.Println(fmt.Sprint(format, params))
  }
}

var colorCodes = make(map[int32]string)

/** Replaces codes "&c", "&a" with ASCII escape color codes */
func colorize(text string) string {
  builder := ""
  next := false
  for _, char := range text {
    if char == '&' {
      next = true
    } else if next {
      replacement := colorCodes[char]
      if &replacement == nil {
        builder += string(char)
        continue
      }
      builder += replacement
      next = false
    } else {
      builder += string(char)
      next = false
    }
  }
  return builder
}
