package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, req *http.Request) {
	filePath := filepath.Join("./", "index.html")

	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	// резервируем место в оперативной памяти
	req.ParseMultipartForm(10 << 20) // 10 MB maximum

	// получаем файл из формы
	file, fileHeader, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// читаем файл и конвертируем
	buff := make([]byte, fileHeader.Size)
	file.Read(buff)

	data, err := service.ConvertMorseOrText(string(buff))
	if err != nil {
		http.Error(w, "ошибка сервера", http.StatusInternalServerError)
		return
	}

	// записываем в файл
	fileName := time.Now().UTC().Format("2006-01-02")
	fileExt := filepath.Ext(fileHeader.Filename)

	filePath := filepath.Join("./", fileName+fileExt)

	err = os.WriteFile(filePath, []byte(data), 0755)
	if err != nil {
		http.Error(w, "ошибка при записи файла", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
