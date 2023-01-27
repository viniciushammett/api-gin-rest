package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"gin-api-rest/controllers"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/vini", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: Valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	}

}