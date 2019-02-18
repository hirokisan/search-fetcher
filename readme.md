# search-fetcher

オープナーとメタキーワードの組み合わせで検索して、結果のページから文章を取得するCLIアプリ
取得した文章はテキストファイルで保存する

## インストール

```
$ go get -u github.com/hirokisan/search-fetcher
```

## 使い方

```
$ search-fetcher --help
$ search-fetcher run --opener=男性 --keyword=不満
```

## 用語
- オープナー
  - ヒトをセグメントに分類する際の切り口
  - ヒトをセグメントに分類する際には特定のオープナー（切り口）を使う
- メタキーワード
  - セグメントから抽出したい情報を取得する際のキーワード
  - セグメントに対してメタキーワードに関連する情報を取得する
