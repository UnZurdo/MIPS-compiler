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
		State: "OK",
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
	} else {
		code := r.Form["code"]
		fmt.Println("code:", code)
		justString := strings.Join(code," ")
		now := time.Now();

		response := PageStatus{
			State: "OK",
			Date: now.Format("02-01-2006"),
			Time: now.Format("15:04:05"),
			Type: "Hexadecimal",
			Code: justString,

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
				if _, err := strconv.ParseInt(elem, 2, 64); err != nil {
					fmt.Println(err)
				} else {
					size := len(elem)

					d1, _ := strconv.ParseInt(elem[0:size/8], 2, 64)
					d2, _ := strconv.ParseInt(elem[size/8:2*size/8], 2, 64)
					d3, _ := strconv.ParseInt(elem[2*size/8:3*size/8], 2, 64)
					d4, _ := strconv.ParseInt(elem[3*size/8:4*size/8], 2, 64)
					d5, _ := strconv.ParseInt(elem[4*size/8:5*size/8], 2, 64)
					d6, _ := strconv.ParseInt(elem[5*size/8:6*size/8], 2, 64)
					d7, _ := strconv.ParseInt(elem[6*size/8:7*size/8], 2, 64)
					d8, _ := strconv.ParseInt(elem[7*size/8:8*size/8], 2, 64)

					fmt.Println(d1,d2,d3,d4,d5,d6,d7,d8)

					bin1 := d1
					bin2 := d2
					bin3 := d3
					bin4 := d4
					bin5 := d5
					bin6 := d6
					bin7 := d7
					bin8 := d8

					fmt.Println("decimal",d1,d2,d3,d4,d5,d6,d7,d8)
					ui1 := uint64(bin1)
					ui2 := uint64(bin2)
					ui3 := uint64(bin3)
					ui4 := uint64(bin4)
					ui5 := uint64(bin5)
					ui6 := uint64(bin6)
					ui7 := uint64(bin7)
					ui8 := uint64(bin8)

					fmt.Println("uint",ui1,ui2,ui3,ui4,ui5,ui6,ui7,ui8)

					hexa_code := "0x"+ strconv.FormatUint(ui1, 16)
					hexa_code += strconv.FormatUint(ui2, 16)
					hexa_code += strconv.FormatUint(ui3, 16)
					hexa_code += strconv.FormatUint(ui4, 16)
					hexa_code += strconv.FormatUint(ui5, 16)
					hexa_code += strconv.FormatUint(ui6, 16)
					hexa_code += strconv.FormatUint(ui7, 16)
					hexa_code += strconv.FormatUint(ui8, 16)

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
		code := r.Form["code"]
		fmt.Println("code:", code)
		justString := strings.Join(code," ")
		now := time.Now();

		response := PageStatus{
			State: "OK",
			Date: now.Format("02-01-2006"),
			Time: now.Format("15:04:05"),
			Type: "Hexadecimal",
			Code: justString,

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


func requestDownload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "POST" {
		r.ParseForm()
		// logic part of log in
		code := r.Form["code"]
		var estado string
		now := time.Now();
		if len(code) == 0 {
			estado = "Syntax error"
		} else {
			estado = "OK"
		}
		fmt.Println("code:", code)
		justString := strings.Join(code," ")
		//fmt.Println("justcode:", justString)

		parsed := strings.Split(justString, "\n")
		var new []string
		for index, elem := range parsed {
			fmt.Println(index, elem)
			new = append(new, elem[2:])
		}

		j:= 1
		var hexa_code = ""
		for i:= 0; i < len(parsed) - 1; i++ {
			parsed[i] = strings.TrimSuffix(parsed[i], "\n")
			parsed[i] = strings.TrimSpace(parsed[i])
			parsed[i] = parsed[i][2:]
			if i != len(parsed) {
				parsed[i] = `X"` + parsed[i] +`", `
				hexa_code += parsed[i]
			}
			if j == 8 {
				j= 0
				hexa_code += "\n"
			}
			j+=1
		}

		for i:= 0; i <= 16*8 - len(parsed); i++ {
			if i < 16*8 - len(parsed) {
				hexa_code += `X"00000000", `
			} else{
				hexa_code += `X"00000000"`
			}

			if j == 8 {
				j = 0
				hexa_code += "\n"
			}
			j+=1
		}



		////////////////////////////

		createFile()
		writeFile(hexa_code)

		//Check if file exists and open
		Openfile, err := os.Open("memoriaRAM_I_ejemplo.vhd")
		defer Openfile.Close() //Close after function return
		if err != nil {
			//File not found, send 404
			http.Error(w, "File not found.", 404)
			return
		}

		//File is found, create and send the correct headers

		//Get the Content-Type of the file
		//Create a buffer to store the header of the file in
		FileHeader := make([]byte, 512)
		//Copy the headers into the FileHeader buffer
		Openfile.Read(FileHeader)
		//Get content type of file
		FileContentType := http.DetectContentType(FileHeader)

		//Get the file size
		FileStat, _ := Openfile.Stat()                     //Get info from file
		FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

		//Send the headers
		w.Header().Set("Content-Disposition", "attachment; filename="+"memoriaRAM_I_ejemplo.vhd")
		w.Header().Set("Content-Type", FileContentType)
		w.Header().Set("Content-Length", FileSize)

		//Send the file
		//We read 512 bytes from the file already so we reset the offset back to 0
		Openfile.Seek(0, 0)
		io.Copy(w, Openfile) //'Copy' the file to the client

		//deleteFile()


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
		code := r.Form["code"]
		fmt.Println("code:", code)
		justString := strings.Join(code," ")
		now := time.Now();

		response := PageStatus{
			State: "OK",
			Date: now.Format("02-01-2006"),
			Time: now.Format("15:04:05"),
			Type: "Hexadecimal",
			Code: justString,

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


func createFile() {

	var file, err = os.Create("./memoriaRAM_I_ejemplo.vhd")
	if isError(err) { return }
	defer file.Close()


	fmt.Println("==> done creating file")
}

func writeFile(info string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile("./memoriaRAM_I_ejemplo.vhd", os.O_RDWR, 0644)
	if isError(err) { return }
	defer file.Close()

	// write some text line-by-line to file

	content := Header() + info + Footer()
	_, err = file.WriteString(content)
	if isError(err) { fmt.Println("error escritura"); return}

	// save changes
	err = file.Sync()
	if isError(err) { return }
	fmt.Println(info)
	fmt.Println("==> done writing to file")
}


func deleteFile() {
	// delete file
	var err = os.Remove("./memoriaRAM_I_ejemplo.vhd")
	if isError(err) { return }

	fmt.Println("==> done deleting file")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/request", request)
	http.HandleFunc("/request/download", requestDownload)
	http.HandleFunc("/request/hexa", requestHexa)

	log.Fatal(http.ListenAndServe(getPort(), nil))
}