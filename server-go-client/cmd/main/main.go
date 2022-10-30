package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	// сработает на  SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) Не сработает.
	signal.Notify(c, os.Interrupt)

	log.Println("Тест го")
	wg := new(sync.WaitGroup)
	count := 30
	//wg.Add(count)
	ticker := time.NewTicker(4 * time.Second)
	for {
		select {
		case <-ticker.C:
			for a := 0; a < count; a++ {
				wg.Add(1)
				go MakeRequest(wg, a)
			}
			wg.Wait()
		case <-c:
			ticker.Stop()
			return
		}
	}

	//Блокировать, пока мы не получим наш сигнал.
	//<-c

}
func MakeRequest(wg *sync.WaitGroup, idx int) {
	defer wg.Done()
	message := map[string]interface{}{
		"cmd":     "load",
		"tabname": "products",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	// resp, err := http.Get("https://go.x.arkadii.ru/api/products")
	resp, err := http.Post("https://go.x.arkadii.ru/api/products", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	//var result map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&result)
	// g := result["result"] //.([]map[string]any)
	// h := g.(string)
	//var h []byte = []byte{}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("\nresult:", idx, "\n", string(buf[:50]))

}
