
## 開発環境
- Go 1.21.2
- MacOSX
- Mysql8系


## 開発環境立ち上げ手順

1. ミドルウェアの立ち上げ
```shell
$ docker-compose up -d
```
2. webサーバの立ち上げ

```shell
$ go run cmd/main.go
```

## データベースマイグレーション手順

1. schema/sql配下に以下ルールの元、ファイルを作成

`X.Y.Z__hoge.sql`

- X: 破壊的変更を加える際
- Y: テーブル追加などの後方互換性が保たれる変更を加える際
- Z: パッチ系
 
2. マイグレーション実行
```shell
$ docker compose run flyway-migrate
```

参考：https://github.com/flyway/flyway-docker?tab=readme-ov-file

3. 初期データ投入

```shell
$ mysql -h 127.0.0.1 -u user -p password < schema/sql/insert_data.sql
```

## ユニットテスト実施方法

```shell
$ go test -v --cover ./...
```

## mockコードの作成方法

internal/service/model/time.goのCustomTimeInterfaceのモックを作成する例
```shell
$ mockgen -source=internal/service/model/time.go -destination=internal/mock/mock_time.go
```

参考：https://github.com/uber-go/mock