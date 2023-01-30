package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"sort" // 排序
)

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var names []string
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	type Point struct {
		X, Y int
	}

	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w) // 打印结构体

	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	// var movies = []Movie{
	// 	{Title: "Casablanca", Year: 1942, Color: false,
	// 		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	// 	{Title: "Cool Hand Luke", Year: 1967, Color: true,
	// 		Actors: []string{"Paul Newman"}},
	// 	{Title: "Bullitt", Year: 1968, Color: true,
	// 		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// }
	// data, err := json.Marshal(movies)
	// if err != nil {
	// 	log.Fatalf("JSON marshaling failed: %s", err)
	// }
	// fmt.Printf("%s\n", data)

	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
