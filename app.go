package main

import (
	"net/http"
	"log"
	"time"
	"html/template"
	"fmt"
	"strings"
	"errors"
	"regexp"
	"strconv"
	"os"
	"bufio"
)

type FormSubmission struct {
	Code string
	Time string
}

type PageStatus struct {
	State string
	Time string
	Date string
	Code string
	Hexadecimal[] string
}


func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func index(w http.ResponseWriter, r *http.Request){
	// fmt.Fprint(w, "Index page")

	now := time.Now();

	response := PageStatus{
		State: "ok",
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, response)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}


func request(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "POST" {
		r.ParseForm()
		// logic part of log in
		code := r.Form["code"]
		fmt.Println("code:", code)
		justString := strings.Join(code," ")
		//fmt.Println("justcode:", justString)

		parsed := strings.Split(justString, "\n")
		for index, elem := range parsed {
			fmt.Println(index, elem)
		}


		now := time.Now();
		reader := bufio.NewReader(os.Stdin)

		err, hexa_code := parse_code(parsed)
		var estado string
		if err != nil {
			estado = err.Error()
			fmt.Println(err)
		} else {
			estado = "OK"
		}

		response := PageStatus{
			State: estado,
			Date: now.Format("02-01-2006"),
			Time: now.Format("15:04:05"),
			Code: justString,
			Hexadecimal: hexa_code,
		}
		t, err := template.ParseFiles("index.html")
		if err != nil {
			log.Print("template parsing error: ", err)
		}
		err = t.Execute(w, response)
		if err != nil {
			log.Print("template executing error: ", err)
		}
	}
}

func parse_code(code[] string) (error, []string) {
	cero := "00000000000000000000000000000000"
	var instr_hexa[] string
	var instr_binary[] string
	for index, line := range code {
		instruction := strings.Split(line, " ")
		if len(instruction) < 3 {
			return errors.New("syntax error"), nil
		}
		for i := 0; i < len(instruction); i++ {
			instruction[i] = strings.ToLower(instruction[i])
			instruction[i] = strings.TrimRight(instruction[i], "\n")
		}

		fmt.Println(index, instruction[0])

		numeric := regexp.MustCompile("[0-9]+")

		if instruction[0] == "nop" {
			instr_hexa = append(instr_hexa, "00000000")
			instr_binary = append(instr_binary, cero)
			fmt.Println(instr_hexa)

		} else if instruction[0] == "add" {
			instr_hexa = append(instr_hexa, "000001")
			instr_binary = append(instr_binary, "000000")
			fmt.Println(instr_hexa)

			r1 := numeric.FindString(instruction[1])
			r2 := numeric.FindString(instruction[2])
			r3 := numeric.FindString(instruction[3])

			i1, _ := strconv.Atoi(r1)
			i2, _ := strconv.Atoi(r2)
			i3, _ := strconv.Atoi(r3)
			//fmt.Println(instruction[0],r1,r2,r3)

			i64_1 := int64(i1)
			i64_2 := int64(i2)
			i64_3 := int64(i3)

			h1 := strconv.FormatInt(i64_1, 16)
			h2 := strconv.FormatInt(i64_2, 16)
			h3 := strconv.FormatInt(i64_3, 16)
			fmt.Println(i1,i2,i3)
			fmt.Println(h1,h2,h3)



		} else if instruction[0] == "lw" {
			instr_hexa = append(instr_hexa, "000010")
			instr_binary = append(instr_binary, "000000")
			fmt.Println(instr_hexa)

		} else if instruction[0] == "sw" {
			instr_hexa = append(instr_hexa, "000011")
			instr_binary = append(instr_binary, "000000")
			fmt.Println(instr_hexa)

		} else if instruction[0] == "beq" {
			instr_hexa = append(instr_hexa, "000100")
			instr_binary = append(instr_binary, "000000")
			fmt.Println(instr_hexa)

		} else if instruction[0] == "addfp" {
			instr_hexa = append(instr_hexa, "100000")
			instr_binary = append(instr_binary, "000000")
			fmt.Println(instr_hexa)

		} else{
			return errors.New("syntax error"), nil
		}
	}
	return nil, instr_hexa
}



func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/request", request)

	log.Fatal(http.ListenAndServe(getPort(), nil))
}