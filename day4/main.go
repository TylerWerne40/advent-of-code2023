package main

import (
	"fmt"
	"os"
  "bufio"
	"slices"
	"strings"
)


func main() {
  var fname string
  fmt.Print("Enter file: ")
  fmt.Scanln(&fname)
  fmt.Printf("Reading file... %s\n", fname)
  file, err := os.Open(fname)
  if err != nil {
    fmt.Printf("Failed to open file: %v", err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  hand := make([]string, 0)
  winners := make([]string, 0)
  sum := 0
  for scanner.Scan() {
    line := scanner.Text()
    nums := strings.Split(line, " ")
    nums = nums[2:]
    found := false
    hand = []string{}
    winners = []string{} // empty arrays
    for _, v := range nums {
      if v == "" {
        continue
      }
      if v == "|" {
        found = true
    	  continue       
      }
      if !found {
        winners = append(winners, v)
      } else {
        hand = append(hand, v)
      }
    }
    val := 0
    for _, v := range hand {
      if slices.Contains(winners, v) {
        if val == 0 {
          val = 1
        } else {
          val = val*2
        }
      }
    }
    fmt.Printf("The value of this hand is... %d\n", val)
    sum += val
  }
  fmt.Printf("The sum of the points is... %d\n", sum)
}
