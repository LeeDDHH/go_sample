# go_sample

- VSCode で作業する際、go の import で指定したパッケージの指定が消える現象が起きるため、以下の設定をする
  - VSCode の拡張機能から[Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)をインストールする
  - `ファイル > ユーザ設定 > 設定` を開き、go の Format Tool を検索する
  - `default` になっている場合、 `gofmt` へ変更する
- モジュール対応モードで動かすことを前提にしている
  - [Go のモジュール管理【バージョン 1.17 改訂版】](https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode)
  - `go env -w GO111MODULE=on`

---

## 構成

|     |        |
| :-: | :----: |
| Go  | 1.18.1 |

---

## コマンド

### ビルド

```bash
go build -o build goファイルパス
```

- `build` に指定した go ファイルのビルドを書き出す
