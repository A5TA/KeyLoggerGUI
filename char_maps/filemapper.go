package charmaps

var ImageMap = map[string]string{
	"5":           "5",
	"PGUP_9":      "PGUP_9",
	"U":           "U",
	"F1":          "F1",
	"END":         "END",
	"4":           "4",
	"F6":          "F6",
	"SPACE":       "SPACE",
	"A":           "A",
	"F7":          "F7",
	"F11":         "F11",
	"C":           "C",
	".":           "DOT", // Special character renamed to DOT
	"W":           "W",
	"Z":           "Z",
	"INS":         "INS",
	"O":           "O",
	"L_CTRL":      "L_CTRL",
	",":           "COMMA", // Special character renamed to COMMA
	"L_SHIFT":     "L_SHIFT",
	"3":           "3",
	"ENTER":       "ENTER",
	"7":           "7",
	"F8":          "F8",
	"RIGHT":       "RIGHT",
	"Q":           "Q",
	"6":           "6",
	"9":           "9",
	"2":           "2",
	"F3":          "F3",
	"'":           "QUOTE", // Special character renamed to QUOTE
	"0":           "0",
	";":           "SEMICOLON", // Special character renamed to SEMICOLON
	"X":           "X",
	"UP_8":        "UP_8",
	"TAB":         "TAB",
	"F5":          "F5",
	"Y":           "Y",
	"H":           "H",
	"INSERT":      "INSERT",
	"M":           "M",
	"E":           "E",
	"F10":         "F10",
	"L_ALT":       "L_ALT",
	"R_ENTER":     "R_ENTER",
	"/":           "SLASH", // Special character renamed to SLASH
	"UP":          "UP",
	"PGUP":        "PGUP",
	"F4":          "F4",
	"DOWN":        "DOWN",
	"L":           "L",
	"ESC":         "ESC",
	"R_ALT":       "R_ALT",
	"PGDN":        "PGDN",
	"D":           "D",
	"DEL":         "DEL",
	"G":           "G",
	"[":           "LEFTBRACKET", // Special character renamed to LEFTBRACKET
	"F9":          "F9",
	"V":           "V",
	"I":           "I",
	"8":           "8",
	"S":           "S",
	"]":           "RIGHTBRACKET", // Special character renamed to RIGHTBRACKET
	"RT_ARROW_6":  "RT_ARROW_6",
	"`":           "GRAVE", // Special character renamed to GRAVE
	"SCROLL_LOCK": "SCROLL_LOCK",
	"N":           "N",
	"K":           "K",
	"LEFT":        "LEFT",
	"END_1":       "END_1",
	"1":           "1",
	"PAUSE":       "PAUSE",
	"F":           "F",
	"J":           "J",
	"R_CTRL":      "R_CTRL",
	"R":           "R",
	"NUM_LOCK":    "NUM_LOCK",
	"F12":         "F12",
	"PRT_SCR":     "PRT_SCR",
	"LEFT_4":      "LEFT_4",
	"-":           "MINUS", // Special character renamed to MINUS
	"CAPS_LOCK":   "CAPS_LOCK",
	"BS":          "BACKSPACE", // Special character renamed to BACKSPACE
	"P":           "P",
	"F2":          "F2",
	"=":           "EQUALS",    // Special character renamed to EQUALS
	"\\":          "BACKSLASH", // Special character renamed to BACKSLASH
	"T":           "T",
	"HOME":        "HOME",
	"PGDN_3":      "PGDN_3",
	"R_SHIFT":     "R_SHIFT",
	"B":           "B",
	// Special Characters
	"!":  "EXCLAMATION", // Shift + 1 renamed to EXCLAMATION
	"@":  "AT",          // Shift + 2 renamed to AT
	"#":  "HASH",        // Shift + 3 renamed to HASH
	"$":  "DOLLAR",      // Shift + 4 renamed to DOLLAR
	"%":  "PERCENT",     // Shift + 5 renamed to PERCENT
	"^":  "CARROT",      // Shift + 6 renamed to CARROT
	"&":  "AMPERSAND",   // Shift + 7 renamed to AMPERSAND
	"*":  "ASTERISK",    // Shift + 8 renamed to ASTERISK
	"(":  "LPAREN",      // Shift + 9 renamed to LPAREN
	")":  "RPAREN",      // Shift + 0 renamed to RPAREN
	"_":  "UNDERSCORE",  // Shift + - renamed to UNDERSCORE
	"+":  "PLUS",        // Shift + = renamed to PLUS
	"{":  "LBRACE",      // Shift + [ renamed to LBRACE
	"}":  "RBRACE",      // Shift + ] renamed to RBRACE
	"|":  "VERTICALBAR", // Shift + \ renamed to VERTICALBAR
	":":  "COLON",       // Shift + ; renamed to COLON
	"\"": "DOUBLEQUOTE", // Shift + ' renamed to DOUBLEQUOTE
	"<":  "LT",          // Shift + , renamed to LT
	">":  "GT",          // Shift + . renamed to GT
	"?":  "QUESTION",    // Shift + / renamed to QUESTION
	"~":  "TILDE",       // Shift + ` renamed to TILDE
}

var CharConvertMap = map[string]string{
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
