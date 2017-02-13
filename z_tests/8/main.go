package main

import (
	"log"
	"net/http"
	"io"
	"os"
	"path/filepath"
	"fmt"
	"io/ioutil"
)

func main()  {


	http.HandleFunc("/", FileUpload)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func FileUpload(w http.ResponseWriter, r *http.Request)  {

	if r.Method == http.MethodPost{

		f, h, err := r.FormFile("f")
		if err != nil {
			log.Println(err)
			return
		}

		defer f.Close()


		path := "./user/"

		finfo, err := os.Stat(path)
		if os.IsNotExist(err){
			fmt.Print(finfo)
			os.Mkdir(path, 0644)
		}

		f1, err := os.Create(filepath.Join(path, h.Filename))
		if err != nil {
			log.Println(err)
			return
		}

		bs, err := ioutil.ReadAll(f)
		f1.Write(bs)

	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
		<form method="post" enctype="multipart/form-data">
		<input type="file" name="f">
		<input type="submit">
		</form>
	`)

}
