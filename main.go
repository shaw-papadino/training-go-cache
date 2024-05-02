package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"training-go-cache/cache"
)

func main() {
	durationSecondsStr := os.Getenv("CACHE_DURATION_SECONDS")
	durationSeconds, err := strconv.Atoi(durationSecondsStr)
	if err != nil {
		fmt.Println("Error parsing CACHE_DURATION_SEVONDS:", err)
		panic(err)
	}

	c := cache.NewCache()

	// 最初の取得試行
	if value, found := c.Get("example"); found {
		fmt.Println(value)
	} else {
		fmt.Println("Not found, setting value...")
		c.Set("example", "my value", time.Duration(durationSeconds) * time.Second)
	}

	time.Sleep(5 * time.Second) // 5秒後に再取得

	if value, found := c.Get("example"); found {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found after 5 seconds")
	}

	time.Sleep(10 * time.Second) // さらに10秒待機

	if value, found := c.Get("example"); found {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found after 15 seconds, expired")
	}
}