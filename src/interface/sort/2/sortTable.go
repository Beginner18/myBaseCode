package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

/*
//sort 包 reverse接口，利用结构体匿名类型嵌套，
//继承嵌套类型的方法
package sort

type reverse struct{ Interface } // that is, sort.Interface

func (r reverse) Less(i, j int) bool { return r.Interface.Less(j, i) }

func Reverse(data Interface) Interface { return reverse{data} }
*/
/*
//Length: time.Duration类型
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}
*/
//Length: 浮点数格式
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length float64
}

//结构体赋值
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

/*
//1 将特定格式字符串解析为time.Duration类型
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
*/
//1 将特定格式字符串解析为time.Duration类型
//2 time包可对time.Duration求其小时/分钟/秒浮点数类型
func length(s string) float64 {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d.Minutes()
}

//格式化输出
//%v相应值的默认格式
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 2, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

/*
//输出格式较差
//需要完全自己做对齐处理
func printTracks(tracks []*Track) {
	const format = "%10v\t%10v\t%10v\t%10v\t%10v\t\n"
	//tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(os.Stdout, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(os.Stdout, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(os.Stdout, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	//tw.Flush() // calculate column widths and print table
}
*/

//按artist排序
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//sort.Sort(byArtist(tracks))
//sort.Sort(sort.Reverse(byArtist(tracks)))
//按年份排序
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//多重排序
//类型嵌套匿名函数
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

//方法调用结构体中嵌套的匿名函数
func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

/*
//实现多重排序: Title Year Length
sort.Sort(customSort{tracks, func(x, y *Track) bool {
    if x.Title != y.Title {
        return x.Title < y.Title
    }
    if x.Year != y.Year {
        return x.Year < y.Year
    }
    if x.Length != y.Length {
        return x.Length < y.Length
    }
    return false
}})
*/
/*
//实现排序：正序及逆序
values := []int{3, 1, 4, 1}
fmt.Println(sort.IntsAreSorted(values)) // "false"
sort.Ints(values)
fmt.Println(values)                     // "[1 1 3 4]"
fmt.Println(sort.IntsAreSorted(values)) // "true"
sort.Sort(sort.Reverse(sort.IntSlice(values)))
fmt.Println(values)                     // "[4 3 1 1]"
fmt.Println(sort.IntsAreSorted(values)) // "false"
*/
func main() {
	printTracks(tracks)
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
	//格式化输出补0
	a := 1
	//int转字符串格式，补0
	b := fmt.Sprintf("%.3d", a)
	fmt.Printf("%T %[1]s\n", b)
}
