package main

// Загружаем переменные в Environment из файла .env
import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func init() {
	loadEnvFile()
}

func loadEnvFile() {
	// путь к исполняемому файлу
	path, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	// проверка на символическую ссылку
	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		log.Println(err)
	}
	//pathEnv := filepath.Join(path, "..", "..", "..", "server", ".env")
	pathEnv := filepath.Join(path, "..", ".env")
	log.Println("Путь к .env:", pathEnv)

	if err := godotenv.Load(pathEnv); err != nil {
		log.Print("Не найден .env файл")
	}
}
