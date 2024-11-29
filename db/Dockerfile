# ベースイメージとして最新のGoを使用
FROM golang:1.23 as builder

WORKDIR /app

# モジュールファイルをコピーして依存関係をダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# Goアプリケーションをビルド
RUN go build -o main .

# Cloud SQL Auth Proxyバイナリを取得して権限を変更
ADD https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 /cloud_sql_proxy
RUN chmod +x /cloud_sql_proxy

# entrypoint.sh を builder ステージにコピー
COPY scripts/entrypoint.sh /scripts/entrypoint.sh

# 実行用のイメージとして軽量な Alpine を使用
FROM alpine:latest as final

# 必要なツールをインストール
RUN apk add --no-cache bash

# 必要なバイナリをコピー
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /cloud_sql_proxy .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /scripts/entrypoint.sh /entrypoint.sh

# 権限を変更
RUN chmod +x /entrypoint.sh

# 起動スクリプトをエントリーポイントとして設定
ENTRYPOINT ["/entrypoint.sh"]

# サーバーポートを公開
EXPOSE 8000
