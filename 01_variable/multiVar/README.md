## 一つの変数宣言キーワードで複数の変数使用

```go
var x,y int = 30, 50
```

x, yは定数で宣言されて初期値は`30`,`50`になります。

```go
var age, name = 10, "Maria"
```

タイプなしで`10`, `"Maria"`の初期値が変数に与えられて、値のタイプが変数に与えられます。

```go
a,b,c,d := 1, 3.4, "Hello, world!", false
```

`変数1,変数2,変数3 := 値1, 値2, 値3`の形で複数の変数に初期値を与えられます。

```go
var (
    x_, y_      int = 30, 50
    age_, name_     = 10, "Maria"
)
```

varキーワードで括弧を利用して、括弧の中はvarなしでvarの文法と同じ形式で変数を定義することができます。