# go_grep

## 速度比較用<br/>ディレクトリ再帰ファイルGREPプログラム<br/>Golang版

 * 検索対象ワードは配列に設定してください。

 checker.go
``` go
/*
 検索対象ワード
*/
var TARGET_WORDS = [...]string{
	"This",
	"Check",
	"Just",
}
```
ex.
```
go run checker.go /usr/local/Cellar .*\\.md
```
 * 引数
   * 0 : 再帰ディレクトリパス
   * 1 : 対象ファイル正規表現

(注) シンボリックリンクは未対応です。
