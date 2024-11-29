# ベースイメージを指定
FROM golang:1.23.2

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールファイルをコピー
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# Goアプリケーションをビルド
RUN go build -o main .

# コンテナ起動コマンド
CMD ["./main"]
