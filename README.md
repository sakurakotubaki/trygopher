# try gopher
Go言語の学習用プロジェクト

**ブランチ毎に学習するコードを書く**

ページのチャプターに合わせて１とかにする。

ブランチ名（例）: basic1

Gitの操作
- コマンドでやりましよう
- [参考記事](https://zenn.dev/zmb/articles/054ba4189244a5)


シンプルなソースコード
```go
package main

import "fmt"

func bazz() {
	fmt.Println("bazz")
}

func main() {
	bazz()
}

```