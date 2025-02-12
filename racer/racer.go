package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(s, f string) (w string) {
	startA := time.Now()
	http.Get(s)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(f)
	bDuration := time.Since(startB)

	fmt.Println(aDuration, bDuration)
	if aDuration > bDuration {
		return f
	}
	return s
}
