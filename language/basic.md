## Golang

- camelCase стиль

- int - платформозависимый тип, 32/64
- int8, int16, int32, int64
- uint - платформозависимый тип, 32/64  
- uint8, unit16, uint32, unit64
- byte = uint8
- float32, float64 - 0 по-умолчанию
- bool - false по-умолчанию
- complex64, complex128 (-1.1 + 7.12i)
- string UTF-8 - пустая строка по-умолчанию

'''golang
var num1 int = 1
num := 30

var weight, height int = 10, 20
weight, age := 12, 22
'''

**String**
одинарные кавычки для байт (uint8), rune (uint32) для UTF-8 символов
'''golang
var rawBinary byte = ’\x27’
helloWorld := "Привет Мир"
helloWorld[0] = 72

byteLen := len(helloWorld) // 19 байт
symbols := utf8.RuneCountInString(helloWorld) // 10 рун
hello := helloWorld[:12] // Привет, 0-11 байты
'''

**конвертация в слайс байт и обратно**
'''golang
byteString = []byte(helloWorld)
helloWorld = string(byteString)
'''

**const**
'''golang
const pi = 3.141
const (
  hello = "Привет"
  e = 2.718
)
'''

**Указатели**
указатель это тип данных, отличается от Си
'''golang
a := 2
b := &a
*b = 3    // a = 3
c := &a   // новый указатель на переменную a
d := new(int) // указатель на тип данных
*d = 12
'''

**Массивы**
размер массива является частью его типа
инициализация значениями по-умолчанию
'''golang
var a1 [3]int   // [0,0,0]

const size = 2
var a2 [2 * size]bool    // [false,false,false,false]

a3 := [...]int{1, 2, 3}  // определение размера при объявлении
'''

**Слайсы**
capacity - сколько еще можно добавить без его расширения (копирования в новую область)
'''golang
var buf0 []int // len=0, cap=0
buf1 := []int{} // len=0, cap=0
buf2 := []int{42} // len=1, cap=1
buf3 := make([]int, 0) // len=0, cap=0
buf4 := make([]int, 5) // len=5, cap=5
buf5 := make([]int, 5, 10) // len=5, cap=10
'''

добавление элементов
'''golang
var buf []int // len=0, cap=0
buf = append(buf, 9, 10) // len=2, cap=2
buf = append(buf, 12) // len=3, cap=4
'''

добавление друго слайса
'''golang
otherBuf := make([]int, 3) // [0,0,0]
buf = append(buf, otherBuf...) // len=6, cap=8
'''

получение среза, указывающего на ту же память
'''golang
sl1 := buf[1:4] // [2, 3, 4]
sl2 := buf[:2] // [1, 2]
'''

если необходимо скопировать
'''golang
newBuf = make([]int, len(buf), len(buf))
copy(newBuf, buf)
'''


