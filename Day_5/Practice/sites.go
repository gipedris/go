/* 
	Реализовать загрузку страниц с пяти разных сайтов с помощью горутин
	Вывести в консоль URL сайта и размер первоначальной страницы в байтах
	
	Подсказки, что нужно использовать:
	
	import (
		"net/http"
		"io/ioutil"
	)

	// Получить URL:
	response, err := http.Get("www.example.com")
	
	defer response.Body.Close()
	
	// Получить тело ответа
	body, err := ioutil.ReadAll(response.Body)

	// Рекомендую создать соответствующую структуру Page с полями URL и Size
	
*/
package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func responseSize(site string, siteChan chan int) {
	response, _ := http.Get(site)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	siteChan <- len(body)
}	

func main() {

	siteChan := make(chan int)

	sites := map[string]int {
		"www.example.com": 0,
		"www.google.com": 0,
		"www.ya.ru": 0,
		"www.dns.ru": 0,
		"www.rambler.ru": 0,
	}	

	

	for url, _ := range sites{
		go responseSize(url, siteChan)
	}

	for url, size := range sites
	siteSize := <- siteChan


}
