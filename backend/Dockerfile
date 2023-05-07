# Go-swagger アプリ用の基本イメージを指定
FROM golang:1.20

# アプリケーションの作業ディレクトリを設定
WORKDIR /app

# モジュールをダウンロードしてキャッシュ
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main ./cmd/golang-nextjs-server/main.go

# ポート番号を環境変数として設定 (Cloud Run は 8080 を推奨)
ENV PORT 8080

# 実行可能なバイナリを指定
ENTRYPOINT ["/app/main"]
