package assets

import (
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"net/http"
	"net/url"
	"sort"
	"time"
)

var client = http.Client{
	Timeout: time.Second * 5,
}

const (
	AssetsURL = "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/"
)

func GetValidators(coin coin.Coin) ([]AssetValidator, error) {
	var results []AssetValidator
	err := blockatlas.Request(&client, AssetsURL+coin.Handle, "/validators/list.json", url.Values{}, &results)
	return results, err
}

func NormalizeValidators(validators []blockatlas.Validator, assets []AssetValidator) []blockatlas.StakeValidator {
	results := make([]blockatlas.StakeValidator, 0)

	for _, v := range validators {
		for _, v2 := range assets {
			if v.ID == v2.ID {
				results = append(results, NormalizeValidator(v, v2))
			}
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].Tokens.Bounded > results[j].Tokens.Bounded
	})

	return results
}

func NormalizeValidator(plainValidator blockatlas.Validator, validator AssetValidator) blockatlas.StakeValidator {
	return blockatlas.StakeValidator{
		ID:     validator.ID,
		Status: plainValidator.Status,
		Info: blockatlas.StakeValidatorInfo{
			Name:        validator.Name,
			Description: validator.Description,
			Image:       GetImage(plainValidator.Coin, plainValidator.ID),
			Website:     validator.Website,
		},
		Reward: plainValidator.Reward,
		Tokens: plainValidator.Tokens,
	}
}

func GetImage(c coin.Coin, ID string) string {
	return AssetsURL + c.Handle + "/validators/assets/" + ID + "/logo.png"
}
