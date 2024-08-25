# pesh-snsc-api

各SNSとの連携管理アプリケーションの[pesh-snsc](https://github.com/penysho/pesh-snsc)からAPIを切り出して管理するプロジェクト

## 使用技術一覧

<p style="display: inline">
 <img src="https://img.shields.io/badge/-Go-555.svg?logo=go&style=for-the-badge">
 <img src="https://img.shields.io/badge/-Gin-555.svg?logo=gin&style=for-the-badge">
 <img src="https://img.shields.io/badge/-Docker-1488C6.svg?logo=docker&style=for-the-badge">
 <img src="https://img.shields.io/badge/-PostgreSQL-336791.svg?logo=postgresql&style=for-the-badge">
</p>

## アーキテクチャ

* クリーンアーキテクチャを採用したプロジェクト構成とする
* モノリポによって、各マイクロサービスをプロジェクト配下で一括管理する
  * 現状、投稿情報を管理する[post-app](post-app)のみ

### 言語

* Golang(1.22.2)

### フレームワーク/ライブラリ

一覧とバージョンは各マイクロサービス配下の[go.mod](post-app/go.mod)を参照する

## 実行

本アプリケーションの実行方法を記載する

### 前提

* pesh-snscのDBが起動していること

### 手順

* 各マイクロサービス配下に`.env.local.compose`を作成する、[環境変数一覧](#環境変数一覧)を参考にすること
* プロジェクトルートで`docker compose up`
* `docker compose exec {マイクロサービス名} /bin/bash`でコンテナの中に入った後、`go run {マイクロサービス名}/cmd/container/api/main.go`を実行

### 環境変数一覧

* DBはpesh-snscで利用するPostgreSQLを想定

| 変数名 | 説明 | デフォルト値 |
| - | - | - |
| DB_HOST | DBホスト | |
| DB_PORT | DBポート | |
| DB_USER | DBユーザー | |
| DB_PASSWORD | DBユーザーのパスワード | |
| DB_NAME | DB名 | |
| DB_SSL_MODE | DBへの暗号化接続設定| false |
| DB_TIMEOUT | DB接続開始時のタイムアウト時間(秒) | 10 |
| DB_MAX_OPEN_CONNECTIONS | DB最大コネクション数 | 10 |
| DB_MIN_OPEN_CONNECTIONS | DB最小コネクション数 | 5 |
| DB_CONNECTION_MAX_LIFETIME | DBコネクションを利用可能な最長時間(秒) | 300 |
| DB_CONNECTION_MAX_IDLE_TIME | DBコネクションがアイドル状態で保持される時間(秒)| 300 |
| APP_ADDRESS | マイクロサービスを起動するアドレス | 0.0.0.0 |
| APP_PORT | マイクロサービスで使用するポート | 8081 |
| APP_ENVIRONMENT | アプリケーションの実行環境(production, staging, development, localから設定) | local |
