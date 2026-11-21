package tests

import (
	"coltstail-server/models"
	"testing"
)

func TestSql(t *testing.T) {
	//config := conf.Conf()
	//engine := routes.New()

	//log.Fatal(engine.Run(config.URL))
	models.TestSql()
}
