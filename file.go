const maxUploadSize = 2 * 1024 * 2014 // 2 MB 
const uploadPath = "./tmp"
 
func main() {
    http.HandleFunc("/upload", uploadFileHandler())
 
    fs := http.FileServer(http.Dir(uploadPath))
    http.Handle("/files/", http.StripPrefix("/files", fs))
 
    log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading files.")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
func uploadFileHandler() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
    if err := r.ParseMultipartForm(maxUploadSize); err != nil {
        renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
        return
    }
	fileType := r.PostFormValue("type")
    file, _, err := r.FormFile("uploadFile")
    if err != nil {
        renderError(w, "INVALID_FILE", http.StatusBadRequest)
        return
    }
    defer file.Close()
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        renderError(w, "INVALID_FILE", http.StatusBadRequest)
        return
    }
	iletype := http.DetectContentType(fileBytes)
    if filetype != "image/jpeg" && filetype != "image/jpg" &&
        filetype != "image/gif" && filetype != "image/png" &&
        filetype != "application/pdf" {
        renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
        return
    }
	fileName := randToken(12)
    fileEndings, err := mime.ExtensionsByType(fileType)
    if err != nil {
        renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
        return
    }
    newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
    fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)
	newFile, err := os.Create(newPath)
    if err != nil {
        renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
        return
    }
    defer newFile.Close()
    if _, err := newFile.Write(fileBytes); err != nil {
        renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
        return
    }
    w.Write([]byte("SUCCESS"))