# public_Golang_repository

Goの練習

## コマンド

### 一時ファイル(バイナリファイル)作成実行

- ctrl + c で実行停止・一時ファイルの削除
- ファイルを編集したら再起動する必要あり
- 開発中に使用
- 停止すると一時ファイルは消える

```code
go run .

// app.go を直接実行
go run ./app.go

```

### パッケージとして実行ファイル（バイナリ）を作成

- パッケージとして実行ファイル（バイナリ）を作成
- oオプションで出力先を指定
- バイナリデータは残る
- 配布・デプロイ向け

```code
// go build -o [保存先]
go build -o public_Golang_repository .
```

### DBについて

- mysqlの使用を想定
- Dockerコンテナをローカルで作成して使用
- go runする前に起動しておく

```code
docker-compose up -d
```
