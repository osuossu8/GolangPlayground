package main

import "golang.org/x/tour/pic"

func MyPic(dx, dy int) [][]uint8 {
    var s [][]uint8

    for i:=0;i<dy;i++{
        a := make([]uint8, dx)
        s = append(s, a)
    }

    return s
}

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy)
    for y := range pic {
        pic[y] = make([]uint8, dx)
        for x := range pic[y]{
            pic[y][x] = uint(x^y+x^y)
        }
    }
}

func main() {
    pic.Show(MyPic)
}
