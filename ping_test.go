package brook

import (
	"fmt"
	"testing"
)

func Test_testmain(t *testing.T) {
	duration, err := PingDuration("www.baidu.com", 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(duration, "ms")
	}

}
