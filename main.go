package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/masatana/go-textdistance"
)

func gofix(filetype string, line string) string {
	// NOTE field よりも自然な区切りにするとよいかも
	tmps := strings.Fields(line)
	var err error
	for i, v := range tmps {
		tmps[i], err = guess(filetype, v)
		if err != nil {
			// TODO: err handling?
			log.Println("gofix", err)
		}
	}
	return strings.Join(tmps, " ")
}

func guess(filetype string, s string) (ret string, err error) {
	powerReg := regexp.MustCompile(`(.*)\*\*([0-9]+)`)

	list, err := loadDict(filetype)
	min := 0.35
	index := -1
	for i, v := range list {
		if s == v {
			return s, nil
		}

		// NOTE: べき乗展開
		// TODO: 個別にON/OFFできるように(特にpythonでは不要)
		result := powerReg.FindStringSubmatch(s)
		if len(result) == 3 {
			a := result[1]
			b, err := strconv.Atoi(result[2])
			if err != nil || b <= 0 {
				// NOTE: no error handling
			} else {
				s = a
				for i := 1; i < b; i++ {
					s += "*" + a
				}
				return s, nil
			}
		}

		tmp := textdistance.LevenshteinDistance(s, v)
		length := len(s) + len(v)
		cost := 0.0
		// 		thLength := 4
		// 		if length <= thLength {
		// 			cost = float64(thLength+1-length) * 0.5
		// 		}
		dist := (float64(tmp) + cost) / float64(length)
		if dist <= min {
			min = dist
			index = i
		}
		//		fmt.Println(s, v, dist)
	}
	if index != -1 {
		//		fmt.Println("min", s, list[index], min)
		s = list[index]
	}
	return s, nil
}

func main() {
	flag.Parse()
	if filetype == "" {
		log.Fatalln("filetype is blank")
	}

	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		ret := gofix(filetype, line)
		fmt.Println(ret)
	}
	if err = scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
