package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	UgcDir = "ugc"

	ServicePort = "8080"
)

type FileObj struct {
	Name     string    `json:"name"`
	ModTime  time.Time `json:"mod_time"`
	Size     int64     `json:"size"`
	RealName string    `json:"real_name"`
}

// panicHandler handles panic(s)
func panicHandler(w http.ResponseWriter) func() {
	return func() {
		if err := recover(); err != nil {
			log.Fatalf("Panic happened, message:%v", err)
			fmt.Fprintln(w, "Internal server error")
		}
	}
}

// uploadFile handles uploads
func uploadFile(w http.ResponseWriter, r *http.Request) {
	log.Info("Hit file uploading router...")
	defer panicHandler(w)()

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Errorf("Failed parsing file, err:%v", err)
		fmt.Fprintln(w, "Failed to parse file.")
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Errorf("Failed retrieving file, err:%v", err)
		fmt.Fprintln(w, "Failed to retrieve file.")
		return
	}
	defer file.Close()

	log.Infof("Uploaded file:%v", handler.Filename)
	log.Infof("File size:%v", handler.Size)
	log.Infof("MIME header:%v", handler.Header)

	tempFile, err := ioutil.TempFile(UgcDir, fmt.Sprintf("upload-*-%s", handler.Filename))
	if err != nil {
		log.Errorf("Failed to write file:%v", err)
		fmt.Fprintln(w, "Failed to upload file.")
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Errorf("Failed to read file:%v", err)
		fmt.Fprintln(w, "Failed to upload file.")
		return
	}

	if _, err = tempFile.Write(fileBytes); err != nil {
		log.Errorf("Failed to save file:%v", err)
		fmt.Fprintln(w, "Failed to upload file.")
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

// downloadFile handles download
func downloadFile(w http.ResponseWriter, r *http.Request) {
	log.Info("Hit file downloading router...")
	defer panicHandler(w)()

	fileName := r.FormValue("file")
	byteData, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", UgcDir, fileName))
	if err != nil {
		log.Errorf("Failed to read:%v, err:%v", fileName, err)
		fmt.Fprintf(w, "File %v not exists", fileName)
		return
	}

	reader := bytes.NewReader(byteData)
	w.Header().Add("Content-Disposition", "Attachment")
	http.ServeContent(w, r, fileName, time.Now(), reader)
}

// listUploadedFiles lists all files uploaded
func listUploadedFiles(w http.ResponseWriter, r *http.Request) {
	log.Info("Hit file listing router...")
	defer panicHandler(w)()

	files, err := ioutil.ReadDir(UgcDir)
	if err != nil {
		log.Errorf("Failed to list dir:%v", err)
		fmt.Fprintln(w, "Failed to list uploaded files.")
	}

	type Files struct {
		Objs []FileObj `json:"objs"`
	}

	fileList := make([]FileObj, 0)
	for _, f := range files {
		fileList = append(fileList, FileObj{
			Name:     strings.Split(f.Name(), "-")[2],
			ModTime:  f.ModTime(),
			Size:     f.Size(),
			RealName: f.Name(),
		})
	}

	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Errorf("Failed to parse HTML file:%v", err)
		fmt.Fprintln(w, "Failed to list uploaded files.")
	}

	tmpl.Execute(w, Files{Objs: fileList})
}

// setupRoutes router registration
func setupRoutes() {
	log.Println("Initializing router")
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/download", downloadFile)
	http.HandleFunc("/", listUploadedFiles)

	log.Printf("Starting listening port %v", ServicePort)
	http.ListenAndServe(fmt.Sprintf(":%v", ServicePort), nil)
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	setupRoutes()
}
