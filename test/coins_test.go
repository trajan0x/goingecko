package coins_test

import (
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestCoinsList(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	coinData, _ := cgClient.GetCoinsList()
	if coinData == nil {
		t.Errorf("Error")
	}

}

func TestCoinsMarket(t *testing.T) {

	cgClient := goingecko.NewClient(nil)
	ids := []string{
		"bitcoin",
		"ethereum",
	}
	priceChange := make([]string, 0)

	priceChange = append(priceChange, "24h")
	coinData, _ := cgClient.GetCoinsMarket("usd", ids, "", "", "100", "1", true, priceChange)
	if coinData == nil {
		t.Errorf("Error")
	}

}

func TestCoinsIds(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	coinData, err := cgClient.GetCoinsId("bitcoin", true, true, true, true, true, true)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if coinData == nil {
		t.Errorf("Error")
	}

}

func TestCoinsIdsTickers(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	coinData, err := cgClient.GetCoinsIdTickers("bitcoin", "", "", "", "", "")
	if coinData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}

}

func TestCoinsIdsHistory(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	coinData, err := cgClient.GetCoinsIdHistory("bitcoin", "30-12-17", true)
	if coinData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}

}

func TestCoinsIdsMarketChart(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	coinData, err := cgClient.GetCoinsIdMarketChart("bitcoin", "usd", "10")
	if coinData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestCoinsIdsMarketChartRange(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	coinData, err := cgClient.GetCoinsIdMarketChartRange("bitcoin", "usd", "1392577232", "1422577232")
	if coinData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestCoinsIdsOhlc(t *testing.T) {

	cgClient := goingecko.NewClient(nil)

	coinData, err := cgClient.GetCoinsOhlc("bitcoin", "usd", "7")
	if coinData == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}