package services

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
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

type Customer struct {
	Client     string
	Status     bool
	LogTime    time.Time
	SystemTime time.Time
	DiffTime   time.Duration
}

func New(repo repository.Repository) Services {
	return &fileService{repo}
}

// var repo repository.Repository

<<<<<<< HEAD
func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func (s fileService) RevFile(fn string) (*[]string, error) {
=======
func (s *fileService) RevFile(fn string) (*[]string, error) {
>>>>>>> office

	file, err := os.Open(fn)
	if err != nil {
		return nil, ErrOpen
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

<<<<<<< HEAD
func (s fileService) GetLocalLogTime(cn string, lf string) (*string, error) {
=======
func (s *fileService) GetLocalLogTime(cn string, lf string) (*string, error) {
>>>>>>> office
	rFile, err := s.RevFile(lf)
	if err != nil {
		return nil, ErrRevFile
	}

	var logs []string
	for i, line := range *rFile {
		if i > 100 {
			break
		}
		if strings.Contains(line, cn) {
			e := `(\d{8}-\d{2}:\d{2}:\d{2})`
			r := regexp.MustCompile(e)
			log := r.FindString(line)
			logs = append(logs, log)
		}
	}
	l := deleteEmpty(logs)
<<<<<<< HEAD
	log.Println(l)
=======
>>>>>>> office
	log := l[0]
	return &log, nil
}

<<<<<<< HEAD
func (s fileService) GetAllTimes(cn string, lf string) (*AllTime, error) {
=======
func (s *fileService) GetAllTimes(cn string, lf string) (*AllTime, error) {
>>>>>>> office
	const layout = "20060102-15:04:05"
	var a AllTime

	llt, err := s.GetLocalLogTime(cn, lf)
	if err != nil {
		return nil, ErrGetLocalLogTime
	}

	lt, err := time.Parse(layout, *llt)
	if err != nil {
		return nil, ErrParse
	}
	a.LogTime = lt
	a.SystemTime = time.Now().UTC()
	a.DiffTime = a.SystemTime.Sub(a.LogTime)
	return &a, nil
}

<<<<<<< HEAD
func (s fileService) CheckValidate(dt time.Duration) bool {
=======
func (s *fileService) CheckValidate(dt time.Duration) bool {
>>>>>>> office
	const t2 time.Duration = 2 * time.Minute
	return t2 > dt
}

<<<<<<< HEAD
func (s fileService) CheckStatus(cn string, lf string) (*Customer, error) {
	at, err := s.GetAllTimes(cn, lf)
	log.Println(at)
=======
func (s *fileService) CheckStatus(cn string, lf string) (*Customer, error) {
	at, err := s.GetAllTimes(cn, lf)
>>>>>>> office
	if err != nil {
		return nil, ErrGetAllTime
	}
	var c Customer
	c.Client = cn
	c.Status = s.CheckValidate(at.DiffTime)
	c.LogTime = at.LogTime
	c.SystemTime = at.SystemTime
	c.DiffTime = at.DiffTime
	return &c, nil
}
