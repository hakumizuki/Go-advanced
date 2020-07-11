package main
import(
	"fmt"
	"io/ioutil"
	"net/http"
	// "net/url"
	// "bytes"

	// アルファベット順にしておくといい
	"./mylib"

	// third-party
	// a "github.com/xxxxx"のように、名前を変えて使用できる
	// また、コンパイルのみ必要で、コードでは使用しない場合は_にする
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	// http
resp, _ := http.Get("http://example.com") // 取ってくる
defer resp.Body.Close()
body, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(body)) // bodyはbyteなので、stringでコンバートする
/* http応用
base, err := url.Parse("http://exammmmplll.c om") //存在してるかどうかの判別
reference, _ := url.Parse("/test?a=1&b=2") // ?の後でGETを&でつなげて書いていく
endpoint := base.ResolveReference(reference).String() // base + reference
fmt.Println(base, err)
fmt.Println(endpoint) // baseのurlの最終尾が意味のない文字列になっていても、ベースの部分を取り出すので問題ない

req, _ := http.NewRequest("GET", endpoint, nil) // POSTの場合はnilがbyteになる
/* POSTの場合
req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(ここにパスワード))) // http.NewRequestはdocを読もう

req.Header.Add("If-None-Match", `W/"wndiown"`) // Headerを付け加える???
q := req.URL.Query() // endpointのqueryがmap[string]stringで取得できる
q.Add("c", "3&%") // programmaticにQueryを追加できる
fmt.Println(q)
fmt.Println(q.Encode()) // &は%26, %は%25になる
req.URL.RawQuery = q.Encode() // Encode()してから入れ直す

var client *http.Client = &http.Client{} // アクセスする時は、クライアントを作る必要がある
resp, _ := client.Do(req) // client.Do()でリクエストを送り、レスポンスを受け取る
body, _ := ioutil.ReadAll(resp.Body) // Bodyを読み込む
fmt.Println(string(body))

*/
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

/* ioutil file読み込み、書き込みなど ※ioはI/O input output 入出力のこと */


/* json Unmarshal&Marshal, Encode ☆json=json=json=json=json=json=json=json=json=json=json=json=json=json=json=json=json=json=json=json☆
import(
	"fmt"
	"encoding/json"
)

type Person struct {
	Name      string   `json:"name"`       // encodeするときのjsonの名前をここで指定できる
	Age       int      `json:"age,string"` // encodeするときのデータ型をstringにできる(Unmarshal時には元のint型に戻る)
	Nicknames []string `json:"-"`          // - にすると隠すことができる
	// 補足 
	// `json:"age,omitempty"` は空("", 0, []など)の場合に表示しない
	// type T struct {}について、 T T `json:T,omitempty` にしたい時、Tではemptyの判別ができないため、*T型にする(nilで空)
}

// struct用にMarshalをカスタマイズする(encode結果をカスタマイズしたいときに使う)
func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal()
}

func main() {
	b := []byte(`{"name":"Mike","age":"20","nicknames":["a","b","c"]}`) // `json配列`の形
	// decodeする手順
	var p Person // Personの変数を宣言
	if err := json.Unmarshal(b, &p); err != nil { // Unmarshalがデータをdecodeして、直接&pにぶち込んでくれる 値の頭文字の大小は関係なくやってくれる
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// 逆にencodeする手順
	v, _ := json.Marshal(p) // encode
	fmt.Println(string(v)) // string casted byte
}
*/






// THIRD PARTY LIBRARYYYYYYYYYYYYs !!!!!
// 1. Semaphore
// go の並列処理の数を制限する
/* Usage 1 AcquireとReleaseでロックをかけたり解除したりする(ブロッキング)
import(
	"context"
	"golang.org/x/sync/semaphore"
)
var s *semaphore.Weighted = semaphore.Weighted(制限したい数)

func longProcess(ctx context.Context) {
	if err := s.Acquire(ctx, 1); err != nil { // LOCKED
		// エラー表示など
		return
	}
	defer s.Release(1) // RELEASED
	// Something
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	// time.Sleep()など
}
*/

/* Usage 2 s.TryAcquireで待機中の処理をキャンセルする
import(
	"context"
	"golang.org/x/sync/semaphore"
)
var s *semaphore.Weighted = semaphore.Weighted(制限したい数)

func longProcess(ctx context.Context) {
	isAcquire := s.TryAcquire(1) // Lockが取れるかを試す
	if !isAcquire { // 取れなかったらreturn
		fmt.Println("Could not get LOCK")
		return
	}
	defer s.Release(1) // 取れたらRELEASE
	// Something
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	// time.Sleep()など
}
*/