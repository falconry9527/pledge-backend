package main

import (
	"encoding/json"
	"fmt"
	"pledge-backend/schedule/models"
	"pledge-backend/utils"
)

func main() {
	res, err := utils.HttpGet("https://tokens.pancakeswap.finance/pancakeswap-top-100.json", map[string]string{})
	if err != nil {
		fmt.Println("UpdateTokenLogo HttpGet err", err)
	} else {
		tokenLogoRemote := models.TokenLogoRemote{}
		err = json.Unmarshal(res, &tokenLogoRemote)
		fmt.Println(tokenLogoRemote)

	}
}
