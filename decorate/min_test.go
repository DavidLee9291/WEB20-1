package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestDecoHandler(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	buf := &bytes.Buffer{} //아래 한줄씩 버퍼에서 읽도록 만들어서 로거에 남길 수 있도록 하려면?
	log.SetOutput(buf)     //표준정보를 버퍼에 넣는다 표준로거의 출력대상(output destination)을 버퍼에 들어가는 식의 출력세팅을 해준다

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	r := bufio.NewReader(buf)    //bufio는 버퍼에서 한줄씩 읽도록 만들어 준다
	line, _, err := r.ReadLine() //로거를 ReadLine() 한줄씩 읽어오도록(버퍼에 있는 내용을)
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Started") //첫줄에 문자열"[LOGGER1] Started")가 포함되어 있어야 한다, line이 빈문자열이면 err(실패)

	//line, _, err = r.ReadLine() // 로거를 ReadLine() 한줄씩 읽어오도록
	//assert.NoError(err)
	//assert.Contains(string(line), "[LOGGER1] Completed")
}
