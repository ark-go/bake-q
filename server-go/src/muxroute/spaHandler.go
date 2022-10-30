package muxroute

import (
	"net/http"
	"os"
	"path/filepath"
)

// spaHandler реализует интерфейс http.Handler, поэтому мы можем его использовать
// для ответа на HTTP-запросы. Путь к статической директории и
// путь к индексному файлу в этом статическом каталоге используется для
// обслуживаем SPA в заданном статическом каталоге.
type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// получить абсолютный путь, чтобы предотвратить обход каталога
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// если нам не удалось получить абсолютный путь, отвечаем 400 неверным запросом
		// и стоп
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// добавьте к пути путь к статическому каталогу
	path = filepath.Join(h.staticPath, path)
	// проверить, существует ли файл по указанному пути
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// Файла не существует, подставим индексный файл
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// если мы получили ошибку (не то, что файл не существует) с указанием
		// файл, вернуть внутреннюю ошибку сервера 500 и остановить
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// в противном случае используйте http.FileServer для обслуживания статического каталога
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
