## byteの範囲

1バイトは8ビットです。2^8範囲のデータを保存できます。-128 ~ 127の範囲です。

ある値のデータの大きさが数-128 ~ 127範囲のビットデータに変換されない場合はbyteタイプの変数に保存できません。

```go
var num1 byte = 'a'
```

文字列'a'は2進数`1100001(2)`(10進数の97)です。-128 ~ 127の範囲内なので値が変数に与えられます。

## 文字値の表記

go言語で文字一つを保存できるデータタイプはrune, char, byteタイプがあります。これらのタイプに文字値を保存するためにはdouble qutationではなくsingle qutationを使用して''値を使用する必要があります。

## 値が変数に与えられない場合

```go
var num2 byte = "ab"
```

"ab"は文字列です。文字ではないのでsingle qutationの使用ができないし、double qutationを使用する必要があります。

文字列はbyteタイプに保存できません。次のようなエラーが発生します。

> cannot use "ab" (untyped string constant) as byte value in variable declaration

```go
var num3 byte = 'ab'
```

一つの文字を保存するためには最小1バイトの空間を要求します。一つの文字が2,3バイト程度の空間を取る場合もあります。

２個の文字は最小限２バイトで表記されます。byteタイプに保存できません。したがって、次のエラーメッセージが出力されます。

> more than one character in rune literal
