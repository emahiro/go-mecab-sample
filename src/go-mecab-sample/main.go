package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bluele/mecab-golang"
)

func parseToNode(m *mecab.MeCab, input string) {
	tg, err := m.NewTagger()
	if err != nil {
		fmt.Printf("NewTagger error. err: %v", err)
		os.Exit(-1)
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(input)
	if err != nil {
		fmt.Printf("NewLattice error. err: %v", err)
		os.Exit(-1)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		fmt.Printf("features: %v\n", features)
		if node.Next() != nil {
			break
		}
	}
}

func main() {
	var input string
	fmt.Println("---- Input your text below ----")
	fmt.Scan(&input)
	m, err := mecab.New("-Owakati")
	if err != nil {
		fmt.Printf("Mecab instance error. err: %v", err)
	}
	defer m.Destroy()

	// parse to node
	parseToNode(m, input)

	fmt.Printf("%v", "Complete !!!")
}
