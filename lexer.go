package main
import ("fmt";"os")
func main() {
var nLines, nChars int
NN_FUN(NewLexer(os.Stdin))
fmt.Printf("%d %d\n", nLines, nChars)
}