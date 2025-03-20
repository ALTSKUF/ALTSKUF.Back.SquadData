package utils

import "math/rand"

func RandomString(l int) string {
  bytes := make([]byte, l)
  for i := range l {
    bytes[i] = byte(randInt(65, 90))
  }
  return string(bytes)
}

func randInt(min, max int) int {
  return min + rand.Intn(max - min)
}
