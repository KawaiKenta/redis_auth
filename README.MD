# redis-auth

gin フレームワークを用いた、ベーシックな認証 API

## 技術構成

- go
- gorm
- mysql
- redis

## API

|                        | メソッド | URI             | 権限           |
| :--------------------- | :------- | :-------------- | :------------- |
| ユーザー登録           | POST     | /signup         | なし           |
| メール認証             | GET      | /verifyemail    | ユーザー登録済 |
| ログイン               | POST     | /login          | メール認証済   |
| ログアウト             | GET      | /logout         | ログイン済     |
| パスワードリセット要求 | POST     | /forgetpassword | ログイン済     |
| パスワードリセット     | PATCH    | /resetpassword  | ログイン済     |
| ログイン情報取得       | GET      | /info           | ログイン済     |

## 開発ロードマップ

- [ ] docker 環境の作成
- [ ] html の全削除
- [ ] メール認証時にリダイレクトするように変更
- [ ] パスワードリセット時にリダイレクトするように変更
- [ ] ユーザーアップデート機能
- [ ] リファクタリング
- [ ] ハードコーディングの除去
- [ ] データベースから model を切り出す
- [ ] validation 層の実装

## 立ち上げ

- config.ini のメールサーバーを設定すること
- データベース、redis の立ち上げ `docker compose up`
- web サーバーの起動 `go run .`
