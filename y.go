//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//line parser.y:11
type yySymType struct {
	yys int
	n   int
}

const NUM = 57346
const ADD = 57347
const LW = 57348
const SW = 57349
const BEQ = 57350
const NOP = 57351
const ADDFP = 57352
const COMA = 57353
const AC = 57354
const CC = 57355
const EOF = 57356

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUM",
	"ADD",
	"LW",
	"SW",
	"BEQ",
	"NOP",
	"ADDFP",
	"COMA",
	"AC",
	"CC",
	"EOF",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:33

func add_to_bin(op, r0, r1, r2 int) (error, string) {

	bin := ""
	if op == 0 {
		bin = "000001" // op
	} else {
		bin = "100000" // op
	}
	ur1 := uint64(r0)
	ur2 := uint64(r1)
	ur3 := uint64(r2)

	sr1 := strconv.FormatUint(ur1, 2)
	if len(sr1) > 5 {
		return errors.New("syntax error"), ""
	} else if len(sr1) < 5 {
		// rellenar con 0' s
		cero := strings.Repeat("0", 5-len(sr1))
		sr1 = cero + sr1
	}

	sr2 := strconv.FormatUint(ur2, 2)
	if len(sr2) > 5 {
		return errors.New("syntax error"), ""
	} else if len(sr2) < 5 {
		// rellenar con 0' s
		cero := strings.Repeat("0", 5-len(sr2))
		sr2 = cero + sr2
	}

	sr3 := strconv.FormatUint(ur3, 2)
	if len(sr3) > 5 {
		return errors.New("syntax error"), ""
	} else if len(sr3) < 5 {
		// rellenar con 0' s
		cero := strings.Repeat("0", 5-len(sr3))
		sr3 = cero + sr3
	}

	shamt := "00000000000"
	bin = bin + sr1 + sr2 + sr3 + shamt

	fmt.Println(bin)

	return nil, bin
}

func nop_to_bin() (error, string) {
	bin := strings.Repeat("0", 32)
	fmt.Println(bin)
	return nil, bin

}

func LW_SW_to_bin(op, r0, r1, r2 int) (error, string) {

	bin := ""
	if op == 0 {
		bin = "000010" // op
	} else {
		bin = "000011" // op
	}
	ur1 := uint64(r0)
	ur2 := uint64(r1)
	ur3 := uint64(r2)

	sr1 := strconv.FormatUint(ur1, 2)
	if len(sr1) > 5 {
		return errors.New("syntax error"), ""
	} else if len(sr1) < 5 {
		// rellenar con 0' s
		cero := strings.Repeat("0", 5-len(sr1))
		sr1 = cero + sr1
	}

	sr2 := strconv.FormatUint(ur2, 2)
	if len(sr2) > 5 {
		return errors.New("syntax error"), ""
	} else if len(sr2) < 5 {
		// rellenar con 0' s
		cero := strings.Repeat("0", 5-len(sr2))
		sr2 = cero + sr2
	}

	sr3 := strconv.FormatUint(ur3, 2)
	if len(sr3) > 5 {
		return errors.New("syntax error"), ""
	} else if len(sr3) < 5 {
		// rellenar con 0' s
		cero := strings.Repeat("0", 5-len(sr3))
		sr3 = cero + sr3
	}

	shamt := "00000000000"
	bin = bin + sr3 + sr1 + shamt + sr3

	fmt.Println(bin)
	return nil, bin
}

func beq_to_bin(r0, r1, r2 int) (error, string) {

	bin := "000100" // op

	ur1 := uint64(r0)
	ur2 := uint64(r1)
	ur3 := uint64(r2)

	sr1 := strconv.FormatUint(ur1, 2)
	if len(sr1) > 5 {
		return errors.New("syntax error"), ""
	} else if len(sr1) < 5 {
		// rellenar con 0' s
		cero := strings.Repeat("0", 5-len(sr1))
		sr1 = cero + sr1
	}

	sr2 := strconv.FormatUint(ur2, 2)
	if len(sr2) > 5 {
		return errors.New("syntax error"), ""
	} else if len(sr2) < 5 {
		// rellenar con 0' s
		cero := strings.Repeat("0", 5-len(sr2))
		sr2 = cero + sr2
	}

	sr3 := strconv.FormatUint(ur3, 2)
	if len(sr3) > 16 {
		// Ultimos 16 bits menos significativos
		sr3 = sr3[len(sr3)-16:]
	} else if len(sr3) < 16 {
		// rellenar con 0' s
		cero := strings.Repeat("0", 16-len(sr3))
		sr3 = cero + sr3
	}

	bin = bin + sr1 + sr2 + sr3

	fmt.Println(bin)
	return nil, bin
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 11
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 38

var yyAct = [...]int{

	5, 8, 7, 9, 10, 6, 11, 38, 31, 37,
	3, 30, 29, 28, 27, 21, 20, 19, 18, 17,
	36, 35, 34, 33, 32, 26, 25, 24, 23, 22,
	16, 15, 14, 13, 12, 4, 2, 1,
}
var yyPact = [...]int{

	-1000, -5, -1000, -1000, -8, 30, 29, 28, 27, 26,
	-1000, -1000, 8, 7, 6, 5, 4, 25, 24, 23,
	22, 21, 3, 2, 0, -1, -3, 20, 19, 18,
	17, 16, -1000, -1000, -4, -6, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 37, 36, 35,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3,
}
var yyR2 = [...]int{

	0, 0, 2, 1, 2, 6, 6, 7, 7, 6,
	1,
}
var yyChk = [...]int{

	-1000, -1, -2, 15, -3, 5, 10, 7, 6, 8,
	9, 14, 4, 4, 4, 4, 4, 11, 11, 11,
	11, 11, 4, 4, 4, 4, 4, 11, 11, 12,
	12, 11, 4, 4, 4, 4, 4, 13, 13,
}
var yyDef = [...]int{

	1, -2, 2, 3, 0, 0, 0, 0, 0, 0,
	10, 4, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 5, 6, 0, 0, 9, 7, 8,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 5:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:25
		{
			add_to_bin(0, yyDollar[2].n, yyDollar[4].n, yyDollar[6].n)
		}
	case 6:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:26
		{
			add_to_bin(1, yyDollar[2].n, yyDollar[4].n, yyDollar[6].n)
		}
	case 7:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:27
		{
			LW_SW_to_bin(0, yyDollar[2].n, yyDollar[4].n, yyDollar[6].n)
		}
	case 8:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:28
		{
			LW_SW_to_bin(1, yyDollar[2].n, yyDollar[4].n, yyDollar[6].n)
		}
	case 9:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:29
		{
			beq_to_bin(yyDollar[2].n, yyDollar[4].n, yyDollar[6].n)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:30
		{
			nop_to_bin()
		}
	}
	goto yystack /* stack new state and value */
}