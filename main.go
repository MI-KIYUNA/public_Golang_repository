package main

import (
	"fmt"
	"math"
	"math/rand" // 乱数生成用
	"net/http"
	"time" // 時刻用
)

func main() {
	// コンソールに出力
	fmt.Println("Hello, 世界")
	fmt.Println("The time is", time.Now())
	fmt.Println("Random number:", rand.Intn(10)) // 0-9 の乱数
	fmt.Printf("25の平方根は %g です。\n", math.Sqrt(25)) // 25の平方根
	fmt.Println(math.Pi)


	// localhost:8080 で HTTP サーバーを起動
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok2")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "hello from Go")
	})

	// 起動（Ctrl+C で停止）
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
