package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	urls := []string{
		"http://gihyo.jp/dev/feature/01/go_4beginners/0001",
		"http://gihyo.jp/dev/feature/01/go_4beginners/0002",
		"http://gihyo.jp/dev/feature/01/go_4beginners/0003",
	}

	/// 逐次実行
	for i, url := range urls {
		res, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		fmt.Println(url, res.Status)
		fmt.Println(i)
	}

	fmt.Printf("time: %f\n\n", (time.Now().Sub(start)).Seconds())

	start = time.Now()

	/// 並列実行
	waitGroup := new(sync.WaitGroup)

	for i, url := range urls {
		// 取得処理をゴルーチンで実行する
		waitGroup.Add(1)

		// go function と書くだけでゴルーチン使って関数を実行できる。便利〜
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)

			waitGroup.Done()
		}(url)

		// ゴルーチンを飛ばしてこれが実行される
		fmt.Println(i)
	}

	// main()が終わらないように待ち合わせる
	waitGroup.Wait()

	fmt.Printf("time: %f\n\n", (time.Now().Sub(start)).Seconds())

	/// channelの利用
	statusChan := make(chan string)

	for i, url := range urls {
		j := i
		// go function と書くだけでゴルーチン使って関数を実行できる。便利〜
		go func(url string) {
			fmt.Println(j)

			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()

			statusChan <- url + " " + res.Status
		}(url)
	}

	// main()が終わらないように待ち合わせる
	waitGroup.Wait()

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}
