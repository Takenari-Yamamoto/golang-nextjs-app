# ベースイメージを指定
FROM golang:1.20 as builder

# ワークディレクトリを設定
WORKDIR /app

# 依存関係をコピー
COPY go.mod .
COPY go.sum .

# 依存関係をインストール
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリをビルド
RUN go build -o golang-nextjs-app cmd/golang-nextjs-server/main.go

# ポートをエクスポート
EXPOSE 8080


# コンテナ起動時のコマンドを設定
CMD ["./golang-nextjs-app", "--host=0.0.0.0", "--port=8080"]
