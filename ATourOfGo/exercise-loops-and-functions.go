package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z := 1.0
    for d := 1.0; //繰り返しの前に初期化
    d*d > 1e-10; //条件式 false になったら終了
    z-=d { //イテレーション毎の最後に実行
        d = (z*z - x) / (2*z)
        fmt.PrintLn(d)
    }
    return z
}

func main(){
    fmt.Println(Sqrt(2))
}
