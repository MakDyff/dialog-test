package services

import (
	"Makdyff/dialog-test/models"

	"github.com/astaxie/beego/logs"
)

type StatisticService struct {
	Ping chan interface{}

	AllInfo map[string][]models.PingModel
	Pings   map[string]models.PingModel

	requestStat1 []models.PingModel
	requestStat2 []models.MaxMinPingModel
}

func NewStatisticService() *StatisticService {
	return &StatisticService{
		Ping:         make(chan interface{}),
		AllInfo:      make(map[string][]models.PingModel),
		Pings:        make(map[string]models.PingModel),
		requestStat1: make([]models.PingModel, 0),
		requestStat2: make([]models.MaxMinPingModel, 0),
	}
}

func (s *StatisticService) Run() {
	go s.write()
}

func (s *StatisticService) GetInfo(siteName string) models.PingModel {
	m := s.Pings[siteName]
	s.requestStat1 = append(s.requestStat1, m)

	return m
}

func (s *StatisticService) GetMaxMinPing() models.MaxMinPingModel {
	m := models.MaxMinPingModel{}

	max := models.PingModel{}
	min := models.PingModel{}
	for _, v := range s.Pings {
		if v.Err != nil {
			continue
		}

		if m.Max == nil {
			max = v
			m.Max = &max
		}

		if m.Min == nil {
			min = v
			m.Min = &min
		}

		if v.PingTime > max.PingTime {
			max = v
		}

		if v.PingTime < min.PingTime {
			min = v
		}
	}

	s.requestStat2 = append(s.requestStat2, m)

	return m
}

// RequestStat Получение топ 5 запросов
func (s *StatisticService) RequestStat() models.RequestStatisticModel {
	top := 5
	m := models.RequestStatisticModel{
		RequestStat1: s.requestStat1,
		RequestStat2: s.requestStat2,
	}

	if len(s.requestStat1) > top {
		m.RequestStat1 = s.requestStat1[len(s.requestStat1)-top : len(s.requestStat1)]
	}
	if len(s.requestStat2) > top {
		m.RequestStat2 = s.requestStat2[len(s.requestStat2)-top : len(s.requestStat2)]
	}

	return m
}

func (s *StatisticService) write() {
	for item := range s.Ping {
		m := item.(models.PingModel)
		logs.Notice(m)

		s.AllInfo[m.Url] = append(s.AllInfo[m.Url], m)
		s.Pings[m.Url] = m
	}
}
