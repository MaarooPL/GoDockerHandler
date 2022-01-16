package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)
// wrzucanie pliku
func Upload(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload files\n")

	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err) //dont do this
	}
	defer file.Close()
	//obadac czy MKdir is best solution
	//err = os.MkdirAll("./files/", os.ModePerm)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	// copy example
	f, err := os.OpenFile(fmt.Sprintf("./files/%s", handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err) //please dont
	}
	defer f.Close()
	io.Copy(f, file)

}
// badanie listy plikow w folderze
func IOReadDir(root string) (string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return "error", err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	result := strings.Join(files, "\n")
	return result, nil
}

type Person struct {
	Lista   string
}
//wyswietlanie listy plikow
func renderTemplate(w http.ResponseWriter, r *http.Request) {
	wstrzyk, _ := IOReadDir("./files/")
	person := Person{Lista: wstrzyk}
	parsedTemplate, _ := template.ParseFiles("template/list.html")
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error occurred while executing the templateor writing its output : ", err)
		return
	}
}
// progam glowny


func main() {
	fmt.Println("dzialam")
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/list", renderTemplate)
	//fmt.Println(IOReadDir("/Users/marekjaszul/go/src/rekrutacjaWWA/files"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server doesn't started.")
	}

}
