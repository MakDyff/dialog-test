package services

import (
	"bufio"
	"os"
	"time"

	"github.com/astaxie/beego/logs"
)

type FileService struct {
	DispatcherEvents
	fullPath string
}

func NewFileService(fullPath string) *FileService {
	return &FileService{
		fullPath: fullPath,
	}
}

func (s *FileService) Run() {
	go s.run()
}

func (s *FileService) run() {
	for {
		time.Sleep(10 * time.Second)
		s.readSites()
	}
}

func (s *FileService) readSites() {
	file, err := os.Open(s.fullPath)
	if err != nil {
		logs.Error(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		site := scanner.Text()
		logs.Debug(site)
		s.SendToHandlers(site)
	}

	if err := scanner.Err(); err != nil {
		logs.Error(err)
	}
}
