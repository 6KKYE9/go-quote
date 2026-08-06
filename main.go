// go-quote：随机名言输出（内置若干条，标准库实现）。
// 用法：go run .            # 随机一条
//      go run . -list       # 列出全部
//      go run . -n 3        # 随机 3 条
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type quote struct {
	Text   string
	Author string
}

var quotes = []quote{
	{"Talk is cheap. Show me the code.", "Linus Torvalds"},
	{"Programs must be written for people to read, and only incidentally for machines to execute.", "Harold Abelson"},
	{"Simplicity is the soul of efficiency.", "Austin Freeman"},
	{"First, solve the problem. Then, write the code.", "John Johnson"},
	{"Any fool can write code that a computer can understand. Good programmers write code that humans can understand.", "Martin Fowler"},
	{"Premature optimization is the root of all evil.", "Donald Knuth"},
	{"Make it work, make it right, make it fast.", "Kent Beck"},
	{"The best error message is the one that never shows up.", "Thomas Fuchs"},
	{"Code is like humor. When you have to explain it, it’s bad.", "Cory House"},
	{"Walking on water and developing software from a specification are easy if both are frozen.", "Edward Berard"},
}

func main() {
	list := flag.Bool("list", false, "列出全部名言")
	n := flag.Int("n", 1, "随机输出的条数")
	seed := flag.Int64("seed", 0, "随机种子（0=用当前时间）")
	flag.Parse()

	rng := rand.New(rand.NewSource(*seed))
	if *seed == 0 {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	if *list {
		for i, q := range quotes {
			fmt.Printf("%2d. %s —— %s\n", i+1, q.Text, q.Author)
		}
		return
	}

	count := *n
	if count < 1 {
		count = 1
	}
	if count > len(quotes) {
		count = len(quotes)
	}
	idxs := rng.Perm(len(quotes))[:count]
	for _, i := range idxs {
		q := quotes[i]
		fmt.Printf("“%s” —— %s\n", q.Text, q.Author)
	}
}
