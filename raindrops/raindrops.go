package raindrops

import "fmt"

var raindrop = map[int]string{
  3: "Pling",
  5: "Plang",
  7: "Plong",
}

// Factors returns all numbers that divided by given number
func Factors(n int) (result []int) {
  for i := 1; i <= n; i++ {
    if n%i == 0 {
      result = append(result, i)
    }
  }

  return
}

// Convert change a number to a string
func Convert(n int) (s string) {
  for _, v := range Factors(n) {
    if val, ok := raindrop[v]; ok {
      s += val
    }
  }

  if s == "" {
    s = fmt.Sprintf("%v", n)
  }

  return
}
