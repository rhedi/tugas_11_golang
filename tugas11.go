package main

import "fmt"
import "strconv"

func main()  {
  var siswalama string
  fmt.Println("Berapa siswa yang Sekolah disini?")
  fmt.Scanln(&siswalama)

  var input, err = strconv.Atoi(siswalama)

  if err==nil{
    fmt.Println(input, "Jumlah siswa yang bersekolah disini")
    fmt.Println(input + 100, "Siswa yang akan masuk")
    fmt.Println(input - 5, "Siswa Perempuan")
    fmt.Println(input * 5, "Siswa Baru yang akan daftar Tahun depan")
    fmt.Println(input /2, "Siswa perempuan yang akan masuk tahun depan")
  }
}
