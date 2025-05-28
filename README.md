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

## git stash

作業を中断するとき

コメントをつけて中断する。

```shell
git stash -m "wip: write md"
```

`git stash` のリスト（`stash@{0}` など）から特定のスタッシュを削除したい場合、下記のコマンドを使います。
### 個別削除コマンド

```shell
git stash drop stash@{<番号>}
```

例として、`stash@{1}` を削除したい場合：

```shell
git stash drop stash@{1}
```

### すべて削除（全消し）
すべてのスタッシュを削除したい場合は、以下のコマンドです：

```shell
git stash clear
```

**注意:** 削除したスタッシュは元に戻せませんので、ご注意ください。
他にも知りたいことがあればご質問ください！

## 1. **squash mergeの実行**
1. **dev2.0ブランチへ切り替え**

```shell
git checkout dev2.0
```

2 **squash mergeの実行**
```shell
git merge --squash topic
```

これで`topic`ブランチの変更がワーキングツリーに反映され、まだコミットはされません。

### 補足
- `--squash` は、マージ元のコミット履歴を1つにまとめてマージします。
- `git commit` を実行するまではブランチの内容は確定しません。

### rebase
1. **リベースを終了する**
   もし作業をやめてリベースを中止したい場合：

```shell
git rebase --abort
```

これでリベース開始前の状態に戻ります。
2. **リベースを完了する**
   コンフリクトなどを解決してリベースを完了したい場合：

```shell
git add <修正したファイル>
git rebase --continue
```

必要なだけ繰り返すとリベースが終わります。
3. **どうしても切り替えたい場合**
   上記どちらかでリベースを終了してから改めて

