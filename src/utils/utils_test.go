package utils

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestRandomString(t *testing.T) {
  randomString := RandomString(32)  
  assert.Equal(t, len(randomString), 32)
}

func TestRandomInt(t *testing.T) {
  randomInt := randInt(20, 50)

  assert.Condition(t, func () bool { return 20 < randomInt && randomInt < 50 })
}
