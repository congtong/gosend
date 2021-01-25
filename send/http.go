package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(url string) string {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n,err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}

type Balance struct {
	Balacne float64 `json:"balance"`
	DelayHashesLastDay int `json:"delay_hashes_last_day"`
	FixedBalance float64 `json:"fixed_balance"`
	FixedValue float64 `json:"fixed_value"`
	HashesLastDay int64 `json:"hashed_last_day"`
	HashedLastHour int64 `json:"hashed_last_hour"`
	Hashrate int64 `json:"hashrate"`
	HashrateHistory map[string]int64 `json:"hashrate_history"`
}
func Status() bool{
	jsonStr := Get("https://api.f2pool.com/eth/0x2aa0d045423e9c5f1adabfc5afd8a9fe44d3a041")
	var result Balance
	res := false
	if err := json.Unmarshal([]byte(jsonStr), &result); err == nil {
		fmt.Println(result.Hashrate)
		if result.Hashrate > 0 {
			res = true
		}
	} else {
		fmt.Println(err)
	}
	return res
}
