package strutils

import (
	"bytes"
	"regexp"
	"strings"
)

func DeSensitization(s string) string {
	reg := regexp.MustCompile("(-(?i)(p)|((--)?(password)(?-i)(=|([ |\t]))?))([a-zA-Z0-9-_@\\*]{6,})")
	sr := reg.FindAllSubmatch([]byte(s), -1)
	for i, v := range sr {
		// fmt.Println(string(sr[i][0]))
		// sr[i][0] = bytes.Replace(sr[i][0], sr[i][len(v)-1], []byte("*****"), -1)
		// fmt.Println(string(sr[i][0]))
		s = strings.Replace(s, string(sr[i][0]), string(bytes.Replace(sr[i][0], sr[i][len(v)-1], []byte("*****"), -1)), -1)
	}
	return s
}
