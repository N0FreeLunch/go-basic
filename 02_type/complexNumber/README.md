## 複素数の表記

数値の右にiを付けると複素数になります。

`1 + 2i`: 定数 + 虚数 = 複素数

`4.2342 + 2.3527i`: 浮動小数店　+ 虚数 = 複素数

`1.e+3i`: 指数表記法の虚数

`7.27151e-10i`: 指数表記法の虚数

`5.32521e-10 + .12345e+2i`: 指数表記法の実数 + 指数表記法の虚数

## 複素数のタイプ

`complex64`: 実数部32ビット＋虚数部32ビット

`complex128`: 実数部64ビット＋虚数部64ビット

## real関数

複素数タイプの値から実数部を取得。

## imag関数

複素数タイプの値から虚数部を取得。

## complex64からfloat32
```go
var r1 float32 = real(num1)
var i1 float32 = imag(num1)
```

例のコードで`num1`はcomplex64なのでcomplex64タイプの複素数から実数部や虚数部は32bitなので取得した値がfloat32タイプの変数に与えれれます。

## complex128からfloat64
```go
var r2 float64 = real(num5)
var i2 float64 = imag(num5)
```

例のコードで`num5`はcomplex128なのでcomplex128タイプの複素数から実数部や虚数部は64bitなので取得した値がfloat64タイプの変数に与えれれます。

## 用語

複素数（ふくそすう）: 복소수

虚数(きょすう): 허수
