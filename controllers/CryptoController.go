package controllers

import (
	"crud-go-echo-gorm/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	coingeckoURL = "https://api.coingecko.com/api/v3"
)

// Constantes para mensagens de erro
const (
	ErrMsgCryptoList    = "Erro ao obter a lista de criptomoedas"
	ErrMsgCryptoInfo    = "Erro ao obter informações da criptomoeda"
	ErrMsgCryptoMarkets = "Erro ao obter informações de mercados da criptomoeda"
)

// CryptoController representa a controller para endpoints relacionados a criptomoedas
type CryptoController struct {
	Client *http.Client
}

// ListaCriptomoedas retorna a lista de criptomoedas
func (cc *CryptoController) CryptoCoinList(c echo.Context) error {
	coins, err := cc.getCoinList()
	if err != nil {
		fmt.Println("err: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrMsgCryptoList})
	}

	return c.JSON(http.StatusOK, coins)
}

// CriptomoedasStatus retorna o status da API
func (cc *CryptoController) CryptoCoinStatus(c echo.Context) error {
	coins, err := cc.getCoinStatus()
	if err != nil {
		fmt.Println("err: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrMsgCryptoList})
	}

	return c.JSON(http.StatusOK, coins)
}

// CriptomoedasSimplePrice retorna o valor simples da moeda
func (cc *CryptoController) CryptoCoinSimplePrice(c echo.Context) error {

	coinIDs := c.QueryParam("ids")
	vsCurrencies := c.QueryParam("vs_currencies")

	simplePrice, err := cc.getSimplePrice(coinIDs, vsCurrencies)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrMsgCryptoMarkets})
	}

	fmt.Println("simplePrice: ", simplePrice)

	return c.JSON(http.StatusOK, simplePrice)
}

// InfoCriptomoeda retorna informações sobre uma criptomoeda específica
func (cc *CryptoController) CryptoCoinInfo(c echo.Context) error {
	coinID := c.Param("id")

	info, err := cc.getCoinInfo(coinID)
	if err != nil {
		fmt.Println("err: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrMsgCryptoInfo})
	}

	return c.JSON(http.StatusOK, info)
}

// CriptomoedasTransaction solicita uma transação
func (cc *CryptoController) CryptoCoinTransaction(c echo.Context) error {

	// coinID := c.Param("id")
	// vsCurrency := c.QueryParam("vs_currency")
	// days := c.QueryParam("days")

	coins, err := cc.getCoinTransaction()
	if err != nil {
		fmt.Println("err: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrMsgCryptoList})
	}

	return c.JSON(http.StatusOK, coins)
}

// MercadosCriptomoeda retorna informações sobre os mercados de uma criptomoeda
func (cc *CryptoController) CryptoCoinMarketChart(c echo.Context) error {
	coinID := c.Param("id")
	vsCurrency := c.QueryParam("vs_currency")
	days := c.QueryParam("days")

	markets, err := cc.getCoinMarketChart(coinID, vsCurrency, days)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrMsgCryptoMarkets})
	}

	return c.JSON(http.StatusOK, markets)
}

// Função auxiliar para obter a lista de criptomoedas
func (cc *CryptoController) getCoinList() ([]models.CoinInfo, error) {
	url := coingeckoURL + "/coins/list"

	body, err := cc.fetchData(url)
	if err != nil {
		return nil, err
	}

	var coins []models.CoinInfo
	err = json.Unmarshal(body, &coins)
	return coins, err
}

// Função auxiliar para obter o status da API
func (cc *CryptoController) getCoinStatus() (string, error) {
	url := coingeckoURL + "/ping"

	statusBytes, err := cc.fetchData(url)
	if err != nil {
		fmt.Println("err: ", err)
		return "", err
	}

	// Convertendo o slice de bytes para string
	status := string(statusBytes)

	return status, nil
}

// Função auxiliar para efetuar a transaction
func (cc *CryptoController) getCoinTransaction() (string, error) {
	url := coingeckoURL + "/transaction"

	statusBytes, err := cc.fetchData(url)
	if err != nil {
		fmt.Println("err: ", err)
		return "", err
	}

	// Convertendo o slice de bytes para string
	status := string(statusBytes)

	return status, nil
}

// Função auxiliar para obter informações sobre uma criptomoeda específica
func (cc *CryptoController) getCoinInfo(coinID string) (models.CoinInfo, error) {
	url := coingeckoURL + "/coins/" + coinID

	body, err := cc.fetchData(url)
	if err != nil {
		return models.CoinInfo{}, err
	}

	var info models.CoinInfo
	err = json.Unmarshal(body, &info)
	return info, err
}

// Criar uma nova estrutura para representar a resposta da API
type SimplePriceResponse struct {
	Prices map[string]map[string]float64 `json:""`
}

// Função auxiliar para obter o preço simples
func (cc *CryptoController) getSimplePrice(coinIDs string, vsCurrencies string) (map[string]map[string]float64, error) {
	url := coingeckoURL + "/simple/price?ids=" + coinIDs + "&vs_currencies=" + vsCurrencies

	body, err := cc.fetchData(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter dados da API: %w", err)
	}

	// Mapear o corpo da resposta para a nova estrutura
	var response SimplePriceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	// Retornar os preços mapeados de acordo com o formato desejado
	formattedResponse := make(map[string]map[string]float64)
	for coinID, prices := range response.Prices {
		formattedResponse[coinID] = make(map[string]float64)
		for currency, price := range prices {
			formattedResponse[coinID][currency] = price
		}
	}

	return formattedResponse, nil
}

// Função auxiliar para obter informações de mercados de uma criptomoeda
func (cc *CryptoController) getCoinMarketChart(coinID string, vsCurrency string, days string) ([]models.MarketData, error) {
	url := coingeckoURL + "/coins/" + coinID + "/market_chart?vs_currency=" + vsCurrency + "&days=" + days

	body, err := cc.fetchData(url)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	var response struct {
		Prices       [][]interface{} `json:"prices"`
		MarketCaps   [][]interface{} `json:"market_caps"`
		TotalVolumes [][]interface{} `json:"total_volumes"`
	}

	// Decodifica a resposta JSON
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return nil, err
	}

	var markets []models.MarketData
	for i := 0; i < len(response.Prices); i++ {
		market := models.MarketData{
			CurrentPrice: models.CurrentPrice{
				USD: response.Prices[i][1].(float64),
			},
			MarketCap: models.MarketCap{
				USD: response.MarketCaps[i][1].(float64),
			},
			TotalVolume: models.TotalVolume{
				USD: response.TotalVolumes[i][1].(float64),
			},
		}
		markets = append(markets, market)
	}

	return markets, nil
}

// Função auxiliar para buscar dados da API
func (cc *CryptoController) fetchData(url string) ([]byte, error) {
	resp, err := cc.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Erro na requisição HTTP: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Resposta HTTP com código de status não OK: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler corpo da resposta HTTP: %v", err)
	}

	return body, nil
}
