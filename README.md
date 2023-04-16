# golang-nextjs-app

- Go
- Next.js
- go-swagger
- REST API
- PostgreSQL
- golang-migrate
- SQL Boiler

# Swagger の型定義から Go のコードを生成

```
swagger generate server -A アプリ名 -f  path/swagger.yml -m modelの場所 -t 出力先
```

swagger generate server: Swagger ツールセットの一部である、Swagger コードジェネレーターを起動するコマンドです。

- -A api: 生成される API の名前を指定します。
- -f schema/swagger.yml: Swagger ファイルのパスを指定します。
- -m restapi/models: 生成されたコードで使用されるモデルの場所を指定します。
- -t internal: 生成されたコードの出力先を指定します。ここでは、internal ディレクトリに出力されます。

```
swagger generate server -A golang-nextjs -f schema/swagger.yml -m restapi/models -t backend/internal
```
