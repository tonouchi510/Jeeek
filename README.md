# Jeeek

Jeeekは「エンジニアのエンジニアによるエンジニアのためのSNSアプリ」です。

サービスのドキュメントは[Wiki](https://github.com/tonouchi510/Jeeek/wiki)を参照してください。

### 稼働サーバ（開発環境）
- swaggger-ui：https://jeeek-dev.appspot.com/_dev/console/
- 管理画面：https://jeeek-dev.appspot.com/_admin/dashboard/


# 開発環境構築
## Prerequires
- goenv
  - Macならbrewで簡単に入れられる
- gcpアカウント（Jeeekプロジェクトに追加します）
- gcloud SDK
- 各種シークレットファイルの配置

### go modules - goの新しいバージョン・依存管理
- GOPATH外でgolangの開発ができるようになる
- ビルドで自動的にダウンロードしてくれる他、go getするだけでバージョン管理される
- [参考文献](https://www.wantedly.com/companies/wantedly/post_articles/132270)


## クイックスタート
golang 1.12.6のインストールと設定
```
$ goenv install 1.12.6
$ goenv local 1.12.6
```

go modulesの初期設定（これは実施済みなので不要）
```
$ go mod init github.com/tonouchi510/Jeeek
```

intellij ideaのgo modules設定  
- GOROOTでv1.12.6のSDKに設定
- https://ema-hiro.hatenablog.com/entry/2019/04/12/020646

依存パッケージのインストール
```
# do.modに記録されているので以下のコマンドだけでよい
$ go build
```

ローカルサーバー起動
```
# datastore emulatorの起動と環境変数設定
$ make ds-start

# サーバ起動
$ make local-run
```


# 開発
- goaによるデザインファーストなAPI開発
  - swaggerフォーマットのAPI定義ファイルが得られる
  - 実行可能なドキュメントが自動で作成・更新できる（swagger-ui）
- GAE/GO gen2 で開発．ローカルではエミュレーターを使うことになる

## ディレクトリ構成
```
$ tree -d -L 1
.
├── cmd            # goaで自動生成される。main.goはここに作られる、がrootに移動させて使う。
├── controller     # コントローラーの置き場。goaでscaffoldを生成してロジックを人が書く。
├── design         # goaのDSLでAPIデザインを書く。
├── factory        # テストデータ生成用に使うライブラリ。
├── gen            # goaが自動生成するコードの置き場。
├── model          # cloud datastoreとのインターフェース。
├── script         # CIとかソースコードの修正に使うshell script等。
└── static         # クライアントサイド、swagger-uiとか管理画面用の静的ファイル置き場。
```

### シークレットファイル
- .env.test
- firebase-service-key-dev.json
- firebase-service-key-prod.json
- gae-service-key-dev.json
- gae-service-key-prod.json
- secret.yaml
  - こいつで環境変数とか設定できる

### 環境変数
- ローカル
  - run-localのスクリプトで設定する
- GCP
  - secret.yamlに記載しておく（ただの記録用）
  - secret.yamlの内容をbase64符号化してcircleciの環境変数に設定する

※環境変数に変更があった時は常に上記の変更を行なっておく

## API開発 with goa
### 1. design/*goの作成、編集
- goa入門者は以下のドキュメントを参照してください
  - https://goa.design/
- dslのドキュメントは以下にある
  - https://godoc.org/gopkg.in/goadesign/goa.v3/dsl

### 2. 自動生成コードの再生成
```
$ make goagen
```

### 3. ビジネスロジックを書く
コントローラーの雛形を生成
```
$ make example
```

- controller/*goの中身を実装  
- main.goは一部、手動修正が必要な場合がある


## ローカルテスト
- firestoreのローカルエミュレーターを使う
- push前に必ずunit test通ることを確認すること

```
$ make ds-start
$ make test
```


## CI/CD
詳細は.circleci/config.ymlを読むように。

- featureブランチ
  - ビルド
  - テスト
- developブランチ
  - ビルド
  - テスト
  - dev環境のカナリアリリース
- masterブランチ
  - ビルド
  - テスト
  - prod環境のカナリアリリース
