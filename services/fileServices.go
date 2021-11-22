package services

import (
	"bufio"
	"os"
	"regexp"
	"time"
	"uaAlert/repository"
)

type fileService struct {
	repo repository.Repository
}

type FileConfig struct {
	AldnFile string `json:"AldnFile"`
	BlpFile  string `json:"BlpFile"`
	ClvFile  string `json:"ClvFile"`
	InsFile  string `json:"InsFile"`
}

type AllTime struct {
	LogTime    time.Time
	SystemTime time.Time
	DiffTime   time.Duration
}

func New(repo repository.Repository) Services {
	return fileService{repo}
}

// var repo repository.Repository

func (s fileService) RevFile(fn string) (*[]string, error) {

	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var names []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		names = append(names, s)
	}

	for i, j := 0, len(names)-1; i < j; i, j = i+1, j-1 {
		names[i], names[j] = names[j], names[i]
	}

	return &names, nil
}

func (s fileService) GetLocalLogTime(lf string) (*string, error) {
	rFile, err := s.RevFile(lf)
	if err != nil {
		return nil, err
	}

	var logs []string
	for i, line := range *rFile {
		if i > 100 {
			break
		}
		e := `(\d{8}-\d{2}:\d{2}:\d{2})`
		r := regexp.MustCompile(e)
		log := r.FindString(line)
		logs = append(logs, log)
	}

	log := logs[0]
	return &log, nil
}

func (s fileService) GetAllTimes(lf string) (*AllTime, error) {
	const layout = "20060102-15:04:05"
	var a AllTime

	llt, err := s.GetLocalLogTime(lf)
	if err != nil {
		return nil, err
	}

	lt, err := time.Parse(layout, *llt)
	if err != nil {
		return nil, err
	}
	a.LogTime = lt.Add(time.Hour * 7)
	a.SystemTime = time.Now()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a, nil
}