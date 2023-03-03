package english

import (
	"strings"
)

var ANSItextcolors = struct {
	black   string
	red     string
	green   string
	yellow  string
	blue    string
	magenta string
	cyan    string
	white   string
	normal  string
}{
	black:   "30",
	red:     "31",
	green:   "32",
	yellow:  "33",
	blue:    "34",
	magenta: "35",
	cyan:    "36",
	white:   "37",
	normal:  "39",
}

var ANSItextstyles = struct {
	bold      string
	underline string
	slowblink string
	reverse   string
}{
	bold:      "1",
	underline: "4",
	slowblink: "5",
	reverse:   "7",
}

var ANSIbackgroundcolors = struct {
	black   string
	red     string
	green   string
	yellow  string
	blue    string
	magenta string
	cyan    string
	white   string
	normal  string
}{
	black:   "40",
	red:     "41",
	green:   "42",
	yellow:  "43",
	blue:    "44",
	magenta: "45",
	cyan:    "46",
	white:   "47",
	normal:  "49",
}

func StyleTerminalOuput(s string, styles_optional ...string) string {

	styleLength := len(styles_optional)

	if styleLength <= 0 {

		return s
	}

	var st strings.Builder

	st.WriteString("\033[")
	for index, style := range styles_optional {

		if index == styleLength-1 {

			st.WriteString(style)
		} else {

			st.WriteString(style + ";")
		}

	}

	st.WriteString("m " + s + " \033[0m")

	return st.String()

}
