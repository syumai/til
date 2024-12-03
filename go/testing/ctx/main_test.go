package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func Test_runServer(t *testing.T) {
	// Contextを用意
	ctx, cancel := context.WithCancel(context.Background())

	// HTTPサーバーを実行
	const testPort = "12345"
	doneCh := make(chan struct{})
	wait := func() { <-doneCh }
	// サーバー終了の待ち受け処理を登録
	t.Cleanup(wait)
	// Contextのキャンセル処理を登録
	t.Cleanup(cancel)
	go func() {
		defer close(doneCh)
		runServer(ctx, testPort)
	}()

	for i := range 10 {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			t.Parallel()

			// GETリクエストを実行
			res, err := http.Get(fmt.Sprintf("http://localhost:%s", testPort))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			// GETリクエストの結果を全て読み取る
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			// 結果を比較
			want := "Hello, World!"
			got := string(body)
			if got != want {
				t.Fatalf("want: %q, got: %q", want, got)
			}
		})
	}
}
