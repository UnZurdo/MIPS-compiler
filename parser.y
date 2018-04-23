%{
package main
import (
"fmt"
"strconv"
"errors"
"strings"
)
%}

%union {
n int
}

%token NUM ADD LW SW BEQ NOP ADDFP COMA AC CC EOF
%%
input:    /* empty */
| input line
;

line:     '\n'
| instruction EOF
;

instruction: ADD NUM COMA NUM COMA NUM  { add_to_bin(0, $2.n,$4.n,$6.n); }
| ADDFP NUM COMA NUM COMA NUM  { add_to_bin(1, $2.n,$4.n,$6.n); }
| SW NUM COMA NUM AC NUM CC  { LW_SW_to_bin(1, $2.n,$4.n,$6.n); }
| LW NUM COMA NUM AC NUM CC  { LW_SW_to_bin(0, $2.n,$4.n,$6.n); }
| BEQ NUM COMA NUM COMA NUM { beq_to_bin($2.n,$4.n,$6.n); }
| NOP { nop_to_bin(); }
;

%%


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
cero := strings.Repeat("0", 5 - len(sr1))
sr1 = cero + sr1
}

sr2 := strconv.FormatUint(ur2, 2)
if len(sr2) > 5 {
return errors.New("syntax error"), ""
} else if len(sr2) < 5 {
// rellenar con 0' s
cero := strings.Repeat("0", 5 - len(sr2))
sr2 = cero + sr2
}

sr3 := strconv.FormatUint(ur3, 2)
if len(sr3) > 5 {
return errors.New("syntax error"), ""
} else if len(sr3) < 5 {
// rellenar con 0' s
cero := strings.Repeat("0", 5 - len(sr3))
sr3 = cero + sr3
}

shamt := "00000000000"
bin = bin + sr3 + sr2 + sr1 + shamt

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
cero := strings.Repeat("0", 5 - len(sr1))
sr1 = cero + sr1
}

sr2 := strconv.FormatUint(ur2, 2)
if len(sr2) > 5 {
return errors.New("syntax error"), ""
} else if len(sr2) < 5 {
// rellenar con 0' s
cero := strings.Repeat("0", 5 - len(sr2))
sr2 = cero + sr2
}

sr3 := strconv.FormatUint(ur3, 2)
if len(sr3) > 5 {
return errors.New("syntax error"), ""
} else if len(sr3) < 5 {
// rellenar con 0' s
cero := strings.Repeat("0", 5 - len(sr3))
sr3 = cero + sr3
}

shamt := "00000000000"
bin = bin + sr3 + sr1 + shamt + sr2

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
cero := strings.Repeat("0", 5 - len(sr1))
sr1 = cero + sr1
}

sr2 := strconv.FormatUint(ur2, 2)
if len(sr2) > 5 {
return errors.New("syntax error"), ""
} else if len(sr2) < 5 {
// rellenar con 0' s
cero := strings.Repeat("0", 5 - len(sr2))
sr2 = cero + sr2
}

sr3 := strconv.FormatUint(ur3, 2)
if len(sr3) > 16 {
// Ultimos 16 bits menos significativos
sr3 = sr3[len(sr3)-16:]
} else if len(sr3) < 16 {
// rellenar con 0' s
cero := strings.Repeat("0", 16 - len(sr3))
sr3 = cero + sr3
}

bin = bin + sr1 + sr2 + sr3

fmt.Println(bin)
return nil, bin
}