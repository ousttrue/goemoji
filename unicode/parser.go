package unicode

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	// "strings"
)

type Emoji struct {
	Group        string
	Subgroup     string
	Unicode      rune
	MajorVersion int
	MinorVersion int
	Name         string
}

func Parse(src []byte) []Emoji {
	buf := bytes.NewBuffer(src)
	scanner := bufio.NewScanner(buf)

	reComment, err := regexp.Compile(`^# (\w+): (\w+)`)
	if err != nil {
		panic("regex")
	}

	reFullyQualified, err := regexp.Compile(`^([0-9A-F]+)\s+; fully-qualified\s+# (\S+) E(\d+)\.(\d+) (.*)`)
	if err != nil {
		panic("regex")
	}

	list := []Emoji{}
	var group = ""
	var subgroup = ""
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		if line[0] == '#' {

			matches := reComment.FindStringSubmatch(line)
			if len(matches) == 3 {
				if matches[1] == "group" {
					group = matches[2]
				} else if matches[1] == "subgroup" {
					subgroup = matches[2]
				} else {
					// fmt.Printf("%s => %s\n", matches[1], matches[2])
				}
			}

		} else {
			matches := reFullyQualified.FindStringSubmatch(line)
			if len(matches) > 0 {
				codepoint, err := strconv.ParseInt(matches[1], 16, 32)
				if err != nil {
					panic("ParseInt")
				}
				major, err := strconv.ParseInt(matches[3], 10, 32)
				if err != nil {
					panic("ParseInt")
				}
				minor, err := strconv.ParseInt(matches[4], 10, 32)
				if err != nil {
					panic("ParseInt")
				}
				list = append(list, Emoji{
					Group:        group,
					Subgroup:     subgroup,
					Unicode:      rune(codepoint),
					MajorVersion: int(major),
					MinorVersion: int(minor),
					Name:         matches[5],
				})
			}
		}
	}
	return list
}
