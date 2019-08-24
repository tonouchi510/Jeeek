# Jeeek

Jeeekは「エンジニアのエンジニアによるエンジニアのための自己実現アプリ」です。

サービスのドキュメントは[Wiki](https://github.com/tonouchi510/Jeeek/wiki)を参照してください。


# 開発環境構築
## Prerequires
- goenv
  - Macならbrewで簡単に入れられる
- gcpアカウント（Jeeekプロジェクトに追加します）
- gcloud SDK
- 各種シークレットファイルの配置


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

依存パッケージのインストール
```
# 以下のコマンドだけでよい
$ go build
```

ローカルサーバー起動
```
# datastore emulatorの起動と環境変数設定
$ make ds-start

# サーバ起動
$ dev_appserver.py --support_datastore_emulator=true ./app.yaml
     or
$ make run
```

### go modules - goの新しいバージョン・依存管理
- GOPATH外でgolangの開発ができるようになる
- ビルドで自動的にダウンロードしてくれる他、go getするだけでバージョン管理される
- [参考文献](https://www.wantedly.com/companies/wantedly/post_articles/132270)


# 開発
- goaによるデザインファーストなAPI開発
  - swaggerフォーマットのAPI定義ファイルが得られる
  - 実行可能なドキュメントが自動で作成・更新できる（swagger-ui）
  - 管理画面も自動生成可能（viron）
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
└──swaggerui       # 開発者用のapiドキュメントサーブ用。
```

### シークレットファイル
- .env.test
- firebaseAccountKey.json
- GCPServiceAccountKey.json
- secret.yaml
  - こいつで環境変数とか設定できる


## APIデザイン with goa
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
$ make regen
```

- controller/*goの中身を実装  
- main.goは一部、手動修正が必要な場合がある

