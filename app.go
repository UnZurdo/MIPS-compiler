package main

import (
	"net/http"
	"log"
	"time"
	"html/template"
	"fmt"
	"strings"
	"os"
	"io"

	"os/exec"
	"bytes"
	"strconv"
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
	Hexadecimal string
	Type string
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
		//Time: now.Format("15:04:05"),
		Type: "Binary",
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
		var estado string
		//writer := bufio.NewReader(os.Stdin)
		var reader io.Reader
		reader = strings.NewReader(justString)
		//var writer io.Writer
		writer := new(bytes.Buffer)


		cmd := exec.Command("./y")
		cmd.Stdout = writer
		//cmd.Stderr = os.Stderr
		cmd.Stdin = reader
		err := cmd.Run()
		if err != nil {
			//fmt.Printf("Failed to start Ruby. %s\n", err.Error())
			//os.Exit(1)
		}

		parsed_code := writer.String()
		if parsed_code == "" {
			estado = "Syntax error"
		} else {
			estado = "OK"
		}
		response := PageStatus{
			State: estado,
			Date: now.Format("02-01-2006"),
			Time: now.Format("15:04:05"),
			Code: justString,
			Hexadecimal: parsed_code,
			Type: "Binary",
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


func requestHexa(w http.ResponseWriter, r *http.Request) {
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
		var estado string
		//writer := bufio.NewReader(os.Stdin)
		var reader io.Reader
		reader = strings.NewReader(justString)
		//var writer io.Writer
		writer := new(bytes.Buffer)


		cmd := exec.Command("./y")
		cmd.Stdout = writer
		//cmd.Stderr = os.Stderr
		cmd.Stdin = reader
		err := cmd.Run()
		if err != nil {
			//fmt.Printf("Failed to start Ruby. %s\n", err.Error())
			//os.Exit(1)
		}

		parsed_code := writer.String()
		if parsed_code == "" {
			estado = "Syntax error"
		} else {
			estado = "OK"
		}

		splited := strings.Split(parsed_code, "\n")
		var hexa_code_array[] string
		for index, elem := range splited {
			fmt.Println(index, elem)
				var bin int64
				if i, err := strconv.ParseInt(elem, 2, 64); err != nil {
					fmt.Println(err)
				} else {
					bin = i
					fmt.Println("decimal",i)
					ui := uint64(bin)
					fmt.Println("uint",ui)

					hexa_code := strconv.FormatUint(ui, 16)
					hexa_code_array = append(hexa_code_array, hexa_code)
					fmt.Println("hexa",hexa_code)
				}




		}
		hexa_code := strings.Join(hexa_code_array,"\n")


		response := PageStatus{
			State: estado,
			Date: now.Format("02-01-2006"),
			Time: now.Format("15:04:05"),
			Code: justString,
			Hexadecimal: hexa_code,
			Type: "Hexadecimal",
		}
		t, err := template.ParseFiles("hexadecimal.html")
		if err != nil {
			log.Print("template parsing error: ", err)
		}
		err = t.Execute(w, response)
		if err != nil {
			log.Print("template executing error: ", err)
		}
	} else {
		now := time.Now();

		response := PageStatus{
			State: "ok",
			Date: now.Format("02-01-2006"),
			Time: now.Format("15:04:05"),
			Type: "Hexadecimal",
		}
		t, err := template.ParseFiles("hexadecimal.html")
		if err != nil {
			log.Print("template parsing error: ", err)
		}
		err = t.Execute(w, response)
		if err != nil {
			log.Print("template executing error: ", err)
		}
	}
}


func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/request", request)
	http.HandleFunc("/request/hexa", requestHexa)

	log.Fatal(http.ListenAndServe(getPort(), nil))
}