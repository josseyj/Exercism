package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

func (self *Robot) Name() string {
	if len(self.name) == 0 {
		self.name = createName()
	}
	return self.name
}

func (self *Robot) Reset() {
	self.name = ""
}

var names map[string]bool

func getNewName() string {
	var name string
	for {
		name = createName()
		if !names[name] {
			break
		}
	}
	return name
}

func createName() string {
	return fmt.Sprintf("%s%03d", randomString(2), rand.Intn(999))
}

const chars string = "ASDGHJKLZXCVBNMQWERTYUIOP"

func randomString(n int) string {
	value := make([]byte, n)
	for i := 0; i < n; i++ {
		value[i] = chars[rand.Intn(len(chars))]
	}
	return string(value)
}
