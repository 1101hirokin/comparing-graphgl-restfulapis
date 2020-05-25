# Restful Web APIとGraphQLのパフォーマンス検証

## 比較対象

- Restful Web API によるユーザー情報CURD操作のサーバーサイドにおけるベンチマーク
- GraphQL によるユーザー情報CURD操作のサーバーサイドにおけるベンチマーク

<br />

## 使用言語・ソフトウェア・アプリケーション・その他ツール

1. サーバー：localhost

	採用理由：サーバー内部の処理速度測定なのでまぁMacBookでも良いかなとローカル環境で検証。


2. コンテナ仮想化：Docker(v19.03.8)、Docker Compose(3.0)

	採用理由：ローカル開発と本番環境の各ソフトや言語の環境差異吸収のため
	| マシン | OS | プロセッサ | メモリ |
	| :---: | :---: | :---: | :---: |
	| MacBook Pro 2016 | macOS Catalina v10.15.4 | 2.6 Ghz クアッドコア Intel Core i7 | 16GB |

3. プログラミング言語：Go(1.14.3-alpine3.11)

	採用理由：個人的にGolangが好き。最近のコンパイル言語の中で一番シンプルにかける。コンパイル言語なので応答速度は十分速い。より実践的にするため軽量Linuxディストリビューションの「alpine」ベースのイメージを使用

4. データベース：MySQL(8.0.20)

	採用理由：特になし。使い慣れているだけ。最新版を使ってみたかったのでDockerHub上のlatest版を使用。

5. ライブラリ（標準ライブラリは省略）

	選定基準：GitHubなどのGitリポジトリホスティングサービス上にあるオープンソースプロジェクトで2020年5月時点で、直近6ヶ月以内にcommit（最近まで活発に開発）があり、スター数が一番多い（デベロッパーからの支持が大きい）もの
	1. ORM：[jinzhu/gorm](https://github.com/jinzhu/gorm)
	2. APIエンドポイント：[gorilla/mux](https://github.com/gorilla/mux)
	3. GraphQL：[graphql-go/graphql](https://github.com/graphql-go/graphql)

<br />

## 検証するユーザーテーブルの構成

|カラム名| id | created_at | updated_at | name | email | bio | url_avatar |
| :--- | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
|型|unsigined integer|datetime|datetime|text|text|text|text|
|例|1|2020-05-23 23:59:59|2020-05-24 18:42:08|例田 太郎|email@example.com|こんにちは！例田太郎です。|https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png|

## エンドポイント

- POST; http://domain.to.server/api/users

	新規にユーザーを作成（CURDのC, Createを担当）

- PUT; http://domain.to.server/api/users/{id}

	idが{id}のユーザー情報を上書き（CURDのU, Updateを担当）

- GET; http://domain.to.server/api/users/{id}

	idが{id}のユーザー情報を取得（CURDのR, Readを担当）

- DELETE; http://domain.to.server/api/users/{id}

	idが{id}のユーザー情報を削除（CURDのD, Deleteを担当）

<br />

- POST; http://domain.to.server/graphql

	CURD全てを担当

	- 例：ユーザー一覧を取得（レスポンスはid, name, email, bioカラムを取得）
		```
		{
			users{
				id,
				name,
				email,
				bio,
			}
		}
		```
	- 例：指定IDのユーザーを取得（レスポンスはid, name, url_avatarカラムを取得）
		```
		{
			user(id: 1){
				id,
				name,
				url_avatar,
			}
		}
		```
	- 例：新規ユーザーを作成（レスポンスはid, nameカラムを取得）
		```
		mutation {
			createUser (
				name: "hiroki",
				email: "example@example.com",
				bio: "Hello GraphQL World !",
			) {
				id
				name
			}
		}
		```
	- 例：指定IDのユーザー情報を上書き（レスポンスはid, name, url_avatar, bioカラムを取得）
		```
		mutation {
			updateUser (
				id: 7
				name: "hiroki_updated",
			) {
				id
				name
				bio
			}
		}
		```
	- 例：指定IDのユーザーを削除（レスポンスはid, nameカラムを取得）
		```
		mutation {
			deleteUser (
				id: 1
			) {
				id
				name
			}
		}
		```