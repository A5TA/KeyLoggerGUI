package main

var imageMap = map[string]string{
	"5":           "key_images/5",
	"PGUP_9":      "key_images/PGUP_9",
	"U":           "key_images/U",
	"F1":          "key_images/F1",
	"END":         "key_images/END",
	"4":           "key_images/4",
	"F6":          "key_images/F6",
	"SPACE":       "key_images/SPACE",
	"A":           "key_images/A",
	"F7":          "key_images/F7",
	"F11":         "key_images/F11",
	"C":           "key_images/C",
	".":           "key_images/.",
	"W":           "key_images/W",
	"Z":           "key_images/Z",
	"INS":         "key_images/INS",
	"O":           "key_images/O",
	"L_CTRL":      "key_images/L_CTRL",
	",":           "key_images/,",
	"L_SHIFT":     "key_images/L_SHIFT",
	"3":           "key_images/3",
	"ENTER":       "key_images/ENTER",
	"7":           "key_images/7",
	"F8":          "key_images/F8",
	"RIGHT":       "key_images/RIGHT",
	"Q":           "key_images/Q",
	"6":           "key_images/6",
	"9":           "key_images/9",
	"2":           "key_images/2",
	"F3":          "key_images/F3",
	"'":           "key_images/'",
	"0":           "key_images/0",
	";":           "key_images/;",
	"X":           "key_images/X",
	"UP_8":        "key_images/UP_8",
	"TAB":         "key_images/TAB",
	"F5":          "key_images/F5",
	"Y":           "key_images/Y",
	"H":           "key_images/H",
	"INSERT":      "key_images/INSERT",
	"M":           "key_images/M",
	"E":           "key_images/E",
	"F10":         "key_images/F10",
	"L_ALT":       "key_images/L_ALT",
	"R_ENTER":     "key_images/R_ENTER",
	"/":           "key_images//",
	"UP":          "key_images/UP",
	"PGUP":        "key_images/PGUP",
	"F4":          "key_images/F4",
	"DOWN":        "key_images/DOWN",
	"L":           "key_images/L",
	"ESC":         "key_images/ESC",
	"R_ALT":       "key_images/R_ALT",
	"PGDN":        "key_images/PGDN",
	"D":           "key_images/D",
	"DEL":         "key_images/DEL",
	"G":           "key_images/G",
	"[":           "key_images/[",
	"F9":          "key_images/F9",
	"V":           "key_images/V",
	"I":           "key_images/I",
	"8":           "key_images/8",
	"S":           "key_images/S",
	"]":           "key_images/]",
	"RT_ARROW_6":  "key_images/RT_ARROW_6",
	"`":           "key_images/`",
	"SCROLL_LOCK": "key_images/SCROLL_LOCK",
	"N":           "key_images/N",
	"K":           "key_images/K",
	"LEFT":        "key_images/LEFT",
	"END_1":       "key_images/END_1",
	"1":           "key_images/1",
	"PAUSE":       "key_images/PAUSE",
	"F":           "key_images/F",
	"J":           "key_images/J",
	"R_CTRL":      "key_images/R_CTRL",
	"R":           "key_images/R",
	"NUM_LOCK":    "key_images/NUM_LOCK",
	"F12":         "key_images/F12",
	"PRT_SCR":     "key_images/PRT_SCR",
	"LEFT_4":      "key_images/LEFT_4",
	"-":           "key_images/-",
	"CAPS_LOCK":   "key_images/CAPS_LOCK",
	"BS":          "key_images/BS",
	"P":           "key_images/P",
	"F2":          "key_images/F2",
	"=":           "key_images/=",
	"\\":          "key_images/BACKSLASH",
	"T":           "key_images/T",
	"HOME":        "key_images/HOME",
	"PGDN_3":      "key_images/PGDN_3",
	"R_SHIFT":     "key_images/R_SHIFT",
	"B":           "key_images/B",
	// Special Characters
	"!":  "key_images/!",  // Shift + 1
	"@":  "key_images/@",  // Shift + 2
	"#":  "key_images/#",  // Shift + 3
	"$":  "key_images/$",  // Shift + 4
	"%":  "key_images/%",  // Shift + 5
	"^":  "key_images/^",  // Shift + 6
	"&":  "key_images/&",  // Shift + 7
	"*":  "key_images/*",  // Shift + 8
	"(":  "key_images/(",  // Shift + 9
	")":  "key_images/)",  // Shift + 0
	"_":  "key_images/_",  // Shift + -
	"+":  "key_images/+",  // Shift + =
	"{":  "key_images/{",  // Shift + [
	"}":  "key_images/}",  // Shift + ]
	"|":  "key_images/|",  // Shift + \
	":":  "key_images/:",  // Shift + ;
	"\"": "key_images/\"", // Shift + '
	"<":  "key_images/ <", // Shift + ,
	">":  "key_images/>",  // Shift + .
	"?":  "key_images/?",  // Shift + /
	"~":  "key_images/~",  // Shift + `

}

var charConvertMap = map[string]string{
	"1":  "!",  // Shift + 1 -> !
	"2":  "@",  // Shift + 2 -> @
	"3":  "#",  // Shift + 3 -> #
	"4":  "$",  // Shift + 4 -> $
	"5":  "%",  // Shift + 5 -> %
	"6":  "^",  // Shift + 6 -> ^
	"7":  "&",  // Shift + 7 -> &
	"8":  "*",  // Shift + 8 -> *
	"9":  "(",  // Shift + 9 -> (
	"0":  ")",  // Shift + 0 -> )
	"-":  "_",  // Shift + - -> _
	"=":  "+",  // Shift + = -> +
	"Q":  "Q",  // Uppercase Q (Shifted)
	"W":  "W",  // Uppercase W (Shifted)
	"E":  "E",  // Uppercase E (Shifted)
	"R":  "R",  // Uppercase R (Shifted)
	"T":  "T",  // Uppercase T (Shifted)
	"Y":  "Y",  // Uppercase Y (Shifted)
	"U":  "U",  // Uppercase U (Shifted)
	"I":  "I",  // Uppercase I (Shifted)
	"O":  "O",  // Uppercase O (Shifted)
	"P":  "P",  // Uppercase P (Shifted)
	"[":  "{",  // Shift + [ -> {
	"]":  "}",  // Shift + ] -> }
	"\\": "|",  // Shift + \ -> |
	"A":  "A",  // Uppercase A (Shifted)
	"S":  "S",  // Uppercase S (Shifted)
	"D":  "D",  // Uppercase D (Shifted)
	"F":  "F",  // Uppercase F (Shifted)
	"G":  "G",  // Uppercase G (Shifted)
	"H":  "H",  // Uppercase H (Shifted)
	"J":  "J",  // Uppercase J (Shifted)
	"K":  "K",  // Uppercase K (Shifted)
	"L":  "L",  // Uppercase L (Shifted)
	";":  ":",  // Shift + ; -> :
	"'":  "\"", // Shift + ' -> "
	"Z":  "Z",  // Uppercase Z (Shifted)
	"X":  "X",  // Uppercase X (Shifted)
	"C":  "C",  // Uppercase C (Shifted)
	"V":  "V",  // Uppercase V (Shifted)
	"B":  "B",  // Uppercase B (Shifted)
	"N":  "N",  // Uppercase N (Shifted)
	"M":  "M",  // Uppercase M (Shifted)
	",":  "<",  // Shift + , -> <
	".":  ">",  // Shift + . -> >
	"/":  "?",  // Shift + / -> ?
	"`":  "~",  // Shift + ` -> ~
}
