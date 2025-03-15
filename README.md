# math-argo-book

Working repo of [問題解決のための「アルゴリズム×数学」が基礎からしっかり身につく本](https://gihyo.jp/book/2022/978-4-297-12521-9).

## Links

* <https://atcoder.jp/contests/math-and-algorithm/score>

## Snippets

### Input

```go
// N
var n int
fmt.Scan(&n)
```

```go
// A B
var a, b int
fmt.Scan(&a, &b)
```

```go
// N
// A1 A2 ... AN
//
// N
var n int
fmt.Scan(&n)
scanner := bufio.NewScanner(os.Stdin)
// a1, a2, ..., an
a := make([]int, n)
scanner.Split(bufio.ScanWords)
for i := 0; i < n; i++ {
  scanner.Scan()
  s := scanner.Text()
  v, _ := strconv.Atoi(s)
  a[i] = v
}
```

```go
// N
// A1 B1
// A2 B2
// ...
// AN BN
//
// N
type cord struct {
	x, y float64
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]cord, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s[i].x, _ = strconv.ParseFloat(scanner.Text(), 64)

		scanner.Scan()
		s[i].y, _ = strconv.ParseFloat(scanner.Text(), 64)
	}
}
```
