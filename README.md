# report-template-generator

## Overview

このジェネレータはあるsectionを繰り返す形式のレポートのテンプレート作成を補助します。
アルゴリズムとデータ構造のレポートで使うテンプレート例を`./template`に添付しています。

## Behavior

このツールはコマンドライン引数で与えられたファイルの数だけsectionを生成し.SourceCodeをファイルの内容で埋めます。

sectionを全て結合し、`./template/full.md`の.Sectionsに埋め込みます。

## How to use

ワンバイナリで動かすためにテンプレートファイルをgoのソースコードに埋め込みます。そのためにstatikというツールを用います。
`go install github.com/rakyll/statik`などでstatikコマンドを使えるようにしてください。

あとはmakeを叩けばいいです。

```bash
make build
```

`report-gen`という実行ファイルが生成されるのでPATHが通っている場所に置いてください。
