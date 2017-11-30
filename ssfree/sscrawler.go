package ssfree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//SSFree is SSClient Free Acounts
type SSFree struct {
	Accounts []SSAccount
	best     int
}

//SSAccount is one account
type SSAccount struct {
	Health   int
	IP       string
	Port     string
	Password string
	Method   string
	Verified string
	Geo      string
}

//GetSSFree read from https
func (s *SSFree) GetSSFree() {
	now := (time.Now().UnixNano()) / 1000000
	url := "https://ss.weirch.com/ss.php?_=" + strconv.FormatInt(now, 10)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("错误信息：", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	ssAounts := make(map[string][][]interface{})
	e2 := json.Unmarshal(body, &ssAounts)
	if e2 != nil {
		fmt.Println("错误信息2：", e2)
		return
	}
	bestDuration := 1000
	s.best = -1
	for _, v := range ssAounts["data"] {
		idx := v[0].(float64)
		var accout SSAccount
		accout.Health = int(idx)
		accout.IP = v[1].(string)
		accout.Port = v[2].(string)
		accout.Password = v[3].(string)
		accout.Method = v[4].(string)
		accout.Verified = v[5].(string)
		accout.Geo = v[6].(string)
		if accout.Health == 100 && accout.Method == "aes-256-cfb" {
			duration, err := testTime(accout.IP)
			fmt.Println("ping ", accout.IP, "  ", duration, " ms")
			if err == nil {
				if bestDuration > duration {
					bestDuration = duration
					s.best = len(s.Accounts)
				}
				s.Accounts = append(s.Accounts, accout)
			}
		}
	}

}

//testTime test the dalay of the network ms
func testTime(address string) (int, error) {

	duration, err := PingDuration(address, 1)

	return duration, err
}

func pingTest(host string) {

}

func jsonto() {
	//jsonStr := "{\"data\":[[100,\"128.199.234.60\",\"19910\",\"69613630\",\"rc4-md5\",\"10:57:05\",\"SG\"],[100,\"45.77.34.244\",\"55555\",\"04625590\",\"chacha20\",\"10:57:05\",\"SG\"]]}"
	m1 := make(map[string][][]string)
	m1["data"] = [][]string{{"Red", "Blue"}, {"Green", "Yellow"}}

	jsn, _ := json.Marshal(m1)
	fmt.Println(string(jsn))
}
