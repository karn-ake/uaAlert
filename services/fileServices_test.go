package services_test

import (
	"testing"
	"time"
	"uaAlert/repository"
	"uaAlert/services"

	"github.com/stretchr/testify/assert"
)

func TestRevFile(t *testing.T) {

	t.Run("test RevFile", func(t *testing.T) {
		repo := repository.NewMock()
		serv := services.New(repo)

		_, err := serv.RevFile("D:\\N2N\\UA_NOMURA\\BS_BLP\\logs\\log.txt")

		assert.ErrorIs(t, err, nil)
	})

	t.Run("test GetLocalLogTime", func(t *testing.T) {
		repo := repository.NewMock()
		serv := services.New(repo)

		_, err := serv.GetLocalLogTime("BLP", "D:\\N2N\\UA_NOMURA\\BS_BLP\\logs\\log.txt")
		assert.ErrorIs(t, err, nil)
	})

	t.Run("test GetAllTimes", func(t *testing.T) {
		repo := repository.NewMock()
		serv := services.New(repo)

		_, err := serv.GetAllTimes("BLP", "D:\\N2N\\UA_NOMURA\\BS_BLP\\logs\\log.txt")
		assert.ErrorIs(t, err, nil)
	})

	t.Run("test CheckValidate", func(t *testing.T) {
		repo := repository.NewMock()
		serv := services.New(repo)
		var dt time.Duration = 3 * time.Minute

		expected := serv.CheckValidate(dt)

		assert.Equal(t, expected, false)
	})

	t.Run("test CheckStatus", func(t *testing.T) {
		repo := repository.NewMock()
		serv := services.New(repo)

		_, err := serv.CheckStatus("BLP", "D:\\N2N\\UA_NOMURA\\BS_BLP\\logs\\log.txt")
		
		assert.ErrorIs(t, err, nil)
	})
}
