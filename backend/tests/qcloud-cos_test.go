package tests

import (
	"coltstail-server/utils"
	"testing"
)

func TestGetToken(t *testing.T) {
	data, err := utils.Token()
	t.Log(data, err)
}
