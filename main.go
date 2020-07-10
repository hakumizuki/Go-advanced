package main
import(
	"fmt"

	// アルファベット順にしておくといい
	"./mylib"

	// third-party
	// a "github.com/xxxxx"のように、名前を変えて使用できる
	// また、コンパイルのみ必要で、コードでは使用しない場合は_にする
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
}

// ディレクトリに移動したのち、$gofmt -w ファイル名 でインデントなどを修正してくれる

// $go get でthird party packageを $GOPATH/src/github.comに取り込む(vgoもある)

// godoc


// 便利な標準パッケージ集
/* time
import(
	"fmt"
	"time" // ☆
)
t := time.Now()
fmt.Println(t.Format(time.RFC3339)) // RFC3339はDBから時間を持ってくるときに使われる
fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()) // など
*/

/* regexp 正規表現
import(
	"fmt"
	"regexp"
)
match, _ := regexp.MatchString("a([a-z]+)e", "apple") // Boolを返す
fmt.Println(match)

r := regexp.MustCompile("a([a-z]+)e") // 繰り返し使うならこうしたほうがいい
ms := r.MatchString("apple")
fmt.Println(ms)

r2 := regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9]+)$")
fs := r2.FindString("/view/test1") // 検索
fss := r2.FindStringSubMatch("/view/tesT3") // 部分分割検索
*/

/* sort
import(
	"fmt"
	"sort"
)

i := []int{2, 7, 1, 5, 3}
s := []string{"a", "g", "f"}
p := []struct{
	Name string
	Age int
}{
	{"Nancy", 20}
	{"Bob", 20}
	{"Jacob", 20}
	{"Keen", 20}
} // ここでしか使わないstructはこう書ける
fmt.Println(i, s, p)
sort.Ints(i)
sort.Strings(s)
sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })

fmt.Println(i, s, p)
*/

/* iota 連番(0, 1, 2, 3 ...)
const(
	c1 = iota
	c2 // 同じ内容なら省略できる
	c3
)

example; バイト
const(
	_ = iota // 0をスルー
	KB = 1 << (10 * iota) // 10回シフト
	MB // 同じ内容なら省略できる
	GB
)

fmt.Println(c1, c2, c3)
fmt.Println(KB, MB, GB)
*/

/* context 長い処理をキャンセルさせる
import(
	"fmt"
	"time"
	"context"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("running...")
	time.Sleep(2 * time.Second)
	fmt.Println("Finish!")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background() // contextをつくる
	ctx, cancel := context.WithTimeout(ctx, 3 * time.Second) // ctxにTimeoutをつけて再代入する
	defer cancel() // contextにcancelが走って、Err():context canceledでDone()になる

	go longProcess(ctx, ch)

	CTXLOOP:
	for {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <- ch:
			fmt.Println(success!!!)
			break CTXLOOP
		}
	}
	fmt.Println("######################")
}

// ctx := context.TODO()は、コンテキストを後から実装する際に使える
*/

/* ioutil file読み込み、書き込みなど ※ioはI/O input output 入出力のこと*/