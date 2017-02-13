package main

import (
	"net/http"
	"io"
	"log"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main()  {

	http.HandleFunc("/", IndexPage)

	http.ListenAndServe(":8080", nil)
}

func IndexPage(w http.ResponseWriter, r *http.Request)  {

	var str string

	// Check to see if the form posted using the POST method
	if r.Method == http.MethodPost {

		// Catch the uploaded file, header and errors
		f, h, err := r.FormFile("q")
		if err != nil{
			log.Println(err)
			return
		}

		defer f.Close()

		//for the information
		fmt.Print("\nFile:", f, "\nHeader:", h, "\nError:", err)

		// Convert the file into a byte slice
		bs, err := ioutil.ReadAll(f)
		if err != nil{
			log.Println(err)
			return
		}

		// Optional: create a new empty file to write the uploaded
		// files bytes to

		os.Mkdir("./user/", 0644)

		f1, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			log.Println(err)
			return
		}

		// Write the byte to the new file
		f1.Write(bs)

		// Convert the slice of bytes into a string
		str = string(bs)
	}




	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="post" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
		</form>
		</br>
	` + str)
}
