package blockatlas

import "github.com/trustwallet/blockatlas/coin"

type ValidatorPage []Validator

type DocsResponse struct {
	Docs interface{} `json:"docs"`
}

const ValidatorsPerPage = 100

type StakingReward struct {
	Annual float64 `json:"annual"`
}

type Tokens struct {
	// Number of staked tokens
	Bounded int64 `json:"bounded"`
}

type Validator struct {
	Coin   coin.Coin
	ID     string        `json:"id"`
	Status bool          `json:"status"`
	Reward StakingReward `json:"reward"`
	Tokens Tokens        `json:"tokens"`
}

type StakeValidatorInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Website     string `json:"website"`
}

type StakeValidator struct {
	ID     string             `json:"id"`
	Status bool               `json:"status"`
	Info   StakeValidatorInfo `json:"info"`
	Reward StakingReward      `json:"reward"`
	Tokens Tokens             `json:"tokens"`
}
