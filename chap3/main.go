package main

import (
  "fmt"
)

func main() {
  // 変数定義
  var n int
  var x,y,z int
  var (
    name string
  )
  // 代入
  n = 5
  x,y = 1,2
  // 定義と代入
  i := 2
  
  var (
    s = "string"
  )
  // 論理値型
  var bo bool
  bo = true
  
  // 数値型
  // 符号付き
  // int8, int16, int32, int64,
  // 符号なし(unsigned)
  // uint8(byte), uint16, uint32, uint64
  // 実装依存
  // int,uint
  // 整数の型変換
  b := byte(n)
  i64 := int64(n)
  u32 := uint32(n)
  
  // 浮動小数点型
  // float32, float64
  f64 := 1.0
  f32 := float32(f64)
  
  // rune型 Unicodeコードポイントを表す
  r := '松'
  fmt.Printf("%v", r) // => "26494"
}