package utils

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	//ConnTimeoutMS conn timeout
	ConnTimeoutMS = 1000
	//ServeTimeoutMS request timeout
	ServeTimeoutMS = 3000
)

//HTTPGet get method
func HTTPGet(url string, connTimeoutMs int, serveTimeoutMs int) (str string, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Duration(connTimeoutMs)*time.Millisecond)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(time.Duration(serveTimeoutMs) * time.Millisecond))
				return c, nil
			},
		},
	}

	reqest, _ := http.NewRequest("GET", url, nil)
	response, err := client.Do(reqest)
	if err != nil {
		err = fmt.Errorf("http failed, GET url:%s, reason:%s", url, err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = fmt.Errorf("http status code error, GET url:%s, code:%d", url, response.StatusCode)
		return
	}

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("cannot read http response, GET url:%s, reason:%s", url, err.Error())
		return
	}
	str = string(resBody)
	return
}

//HTTPPost post method
func HTTPPost(url string, data string, connTimeoutMs int, serveTimeoutMs int) (str string, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Duration(connTimeoutMs)*time.Millisecond)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(time.Duration(serveTimeoutMs) * time.Millisecond))
				return c, nil
			},
		},
	}

	body := strings.NewReader(data)
	reqest, _ := http.NewRequest("POST", url, body)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(reqest)
	if err != nil {
		err = fmt.Errorf("http failed, POST url:%s, reason:%s", url, err.Error())
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = fmt.Errorf("http status code error, POST url:%s, code:%d", url, response.StatusCode)
		return
	}

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("cannot read http response, POST url:%s, reason:%s", url, err.Error())
		return
	}
	str = string(respBody)
	return

}
