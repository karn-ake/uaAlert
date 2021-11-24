package services_test

import (
	"testing"
	"uaAlert/repository"
	"uaAlert/services"

	"github.com/stretchr/testify/assert"
)

func TestRevFile(t *testing.T) {
	t.Run("test CheckStatus", func(t *testing.T) {
		repo := repository.NewMock()
		repo.On("GetClient").Return(repository.Client{
			LogFile:    "D:\\N2N\\UA_NOMURA\\BS_BLP\\logs\\log.txt",
			ClientName: "BLP",
		})

		serv := services.New(repo)
		_, err := serv.CheckStatus("BLP", "D:\\N2N\\UA_NOMURA\\BS_BLP\\logs\\log.txt")
		var expected error
		expected = nil

		assert.Equal(t, expected, err)
	})
}
