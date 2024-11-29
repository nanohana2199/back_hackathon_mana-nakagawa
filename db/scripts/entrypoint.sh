#!/bin/bash
# Cloud SQL Auth Proxyをバックグラウンドで実行
/cloud_sql_proxy -dir=/cloudsql &

# Goアプリケーションを起動
exec /app/main