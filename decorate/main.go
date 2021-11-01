package main

import (
	"goweb/WEB10/decoHandler"
	"goweb/WEB10/myapp"
	"log"
	"net/http"
	"time"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) { // log데코레이터를 만든다, h http.Handler 인자를 추가받는다
	start := time.Now()                                                       //(4) req를 수행하는데 걸리는 시간
	log.Println("[LOGGER1] Started")                                          //(1) req와 왔을때 Hadler 호출하기전에 먼저 started log를 찍는다
	h.ServeHTTP(w, r)                                                         //(2) http.ServerHTTP 호출
	log.Println("[LOGGER1] Completed time", time.Since(start).Milliseconds()) //(3) 끝날때 logger를 찍는다 (5) 특정time으로부터 duration(걸리는) 시간 Mullisecondes()단위로
}

//func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) { //log데코레이터를 만든다
//start := time.Now()
//log.Println("[LOGGER1] Started")
//h.ServeHTTP(w, r)
//log.PrintLn("[LOGGER1] Completed time", time Since(start).Milliseconds())
//}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	//h := myapp.NewHandler()
	h := decoHandler.NewDecoHandler(mux, logger) //먹스와 로거가 동시에 래핑되어있는 것(감싸주는것) 새로운 decoHandler패키지에서 myapp.NewHandler() mus를 감싸준다
	//h := decoHandler.NewDecoHandler(h, logger2)
	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux) // 서버가 먹스에 들어온 것에 대한 전용서버되는 것
}
