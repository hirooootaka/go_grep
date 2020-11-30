package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

/*
 ディレクトリ再帰ファイルGREPプログラム Golang版

 引数
 [0]:再帰ディレクトリパス
 [1]:対象ファイル正規表現

 ex. go run checker.go /usr/local/Cellar .*\\.md
*/

/*
 検索対象ワード
*/
var TARGET_WORDS = [...]string{
	"This",
	"Check",
	"Just",
}

/*
 検索ワード有無Map
*/
var wordExistMap = map[string]bool{}

/*
 RegexパターンMap
*/
var regPatternMap = map[string]*regexp.Regexp{}

/*
 main
*/
func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("args is [0]:path [1]:regex")
		os.Exit(0)
	}

	t1 := time.Now()
	fmt.Println(t1)

	fmt.Println("** init **")
	for _, word := range TARGET_WORDS {
		if word != "" {
			wordExistMap[word] = false
			regPatternMap[word] = regexp.MustCompile(word)
		}
	}

	fmt.Println("** start **")
	fe := regexp.MustCompile(flag.Arg(1))
	paths := dirwalk(fe, flag.Arg(0))

	fmt.Println("** regexp **")
	for n, v := range paths {
		//fmt.Printf("%d/%d\n", n+1, len(paths))
		r := findReg(v)
		if len(r) > 0 {
			fmt.Printf("\n%d/%d[%s]\n", n+1, len(paths), v)
			for _, l := range r {
				fmt.Printf("%s\n", l)
			}
		}
	}

	fmt.Println("** result **")
	for _, word := range TARGET_WORDS {
		if word != "" {
			fmt.Printf("%s: %t\n", word, wordExistMap[word])
		}
	}

	t2 := time.Now()
	fmt.Println(t2)
	fmt.Println(t2.Sub(t1))
}

/*
 対象ワード検索処理
*/
func findReg(fileName string) []string {
	var find = []string{}

	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	num := 0
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		num++
		line := scanner.Text()
		//fmt.Println(line)
		for _, word := range TARGET_WORDS {
			if word != "" {
				if regPatternMap[word].MatchString(line) {
					wordExistMap[word] = true
					s := fmt.Sprintf("[%s](%d): %s", word, num, line)
					find = append(find, s)
				}
			}
		}
	}
	return find
}

/*
 ディレクトリ再帰処理
*/
func dirwalk(fe *regexp.Regexp, dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(fe, filepath.Join(dir, file.Name()))...)
		} else {
			if fe.MatchString(file.Name()) {
				paths = append(paths, filepath.Join(dir, file.Name()))
			}
		}
	}

	return paths
}
