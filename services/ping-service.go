package services

import (
	"Makdyff/dialog-test/models"
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"
)

type PingService struct {
	DispatcherEvents
	UrlString chan interface{}
}

func NewPingService() *PingService {
	return &PingService{
		UrlString: make(chan interface{}),
	}
}

func (s *PingService) Run() {
	go s.pingByUrl()
	logs.Info("PingService run.")
}

func (s *PingService) pingByUrl() {
	for item := range s.UrlString {
		urlFull := item.(string)

		m := models.PingModel{
			Url:       urlFull,
			CreatedOn: time.Now(),
		}

		response, err := http.Get(fmt.Sprint("https://www.", urlFull))

		if err != nil {
			logs.Error("The HTTP request failed with error: ", err)
			m.Err = err
		} else {
			logs.Notice(urlFull, "->", response.Status, "waist time :", time.Now().Sub(m.CreatedOn))
			m.Status = response.Status
			m.StatusCode = response.StatusCode
		}
		m.PingTime = time.Now().Sub(m.CreatedOn)

		s.SendToHandlers(m)
	}
}
