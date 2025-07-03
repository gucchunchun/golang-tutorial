# golang-tutorial-todolist



## 課題1: Go言語をインストールし、Hello Worldを出力する

### 学び
- modファイルの作り方
- packageのインポート方法
- fmtの使用方法

### 手順

#### モジュールの追加
1. go.modファイルの作成
```
go mod golang-tutorial-todolist
```

#### mainパッケージの追加
1. main.goファイルの作成
```
touch main.go
```
2. package名の指定
3. ***func main()**の追加
4. ***fmt**パッケージのインポート
5. "Hello World"の出力確認
```
go run main.go
```


## 課題2: 変数の宣言と初期化（var, :=）を使い分ける

### 学び
- var
  - 初期値の宣言を行わない場合に適切
  - 初期値を宣言しない場合はzero-valueが設定される
    bool: false
    int / float: 0
    string: ""
  - ```var()```とすることで複数変数を宣言できる
- :=
  - 初期値の宣言を行う場合に有効。型指定しなくても良い。



## 課題3: 基本データ型（int, string, bool, float64）を使用する

[Data Types in Golang](https://www.geeksforgeeks.org/go-language/data-types-in-go/)

### 学び
- 別タイプの数値は計算することができない
