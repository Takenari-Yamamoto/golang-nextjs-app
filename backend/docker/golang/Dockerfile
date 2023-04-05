# ベースイメージを指定
FROM golang:1.20-alpine

# 作業ディレクトリを指定
WORKDIR /go/src/app

# ローカルファイルをコンテナ内にコピー
COPY backend .

# パッケージをダウンロードし、アプリケーションをビルド
RUN go mod tidy

# アプリケーションを実行
CMD ["go", "run", "cmd/main.go"]
