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


## 課題4: fmt.Printf()で様々なフォーマット指定子を使用する

### 学び
| formatter | value |
|-----------|-------|
|     %v      |   value in default format   |
|     %T      |   type   |
|     %p      |   pointer    |
|      %t     |    boolean   |
|       %b    |    base2   |
|       %d    |    base10   |
|       %o    |    base8  |
|       %g    |    float32 / 64  |
|       %s    |    string  |


## 課題5: if文とelse文を使った条件分岐を実装する


## 課題6: for文を使った基本的なループを実装する


## 課題7: switch文を使った複数条件の分岐を実装する


## 課題8: 引数と戻り値を持つ関数を定義する

### 学び
- 関数内に名前付き関数は定義できない
  = 全ての関数はトップレベルで定義されるべき
  = コンパイルが早くなる


## 課題9: 複数の戻り値を返す関数を作成する


## 課題10: エラーを返す関数を作成し、適切にハンドリングする
