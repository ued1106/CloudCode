# CloudCode

Claude Code (クラウド) のお試し用リポジトリです。

## hello

起動すると「こんにちは」とメッセージボックスで表示する Windows 用の exe です。

- `hello/hello.exe` … Windows 64bit 用の実行ファイル
- `hello/main.go` … ソースコード (Go)

ビルド方法:

```bash
cd hello
GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui -s -w" -o hello.exe .
```
