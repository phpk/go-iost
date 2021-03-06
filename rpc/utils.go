package rpc

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/iost-official/go-iost/common"
	"github.com/iost-official/go-iost/core/tx"
	"math"
	"regexp"
	"strconv"
)

func checkAmount(amount string, token string) error {
	matched, err := regexp.MatchString("^([0-9]+[.])?[0-9]+$", amount)
	if err != nil || !matched {
		return fmt.Errorf("invalid amount: %v", amount)
	}
	f1, err := common.NewFixed(amount, -1)
	if err != nil {
		return fmt.Errorf("invalid amount: %v, %v", err, amount)
	}
	f2, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("invalid amount: %v, %v", err, amount)
	}
	if math.Abs(f1.ToFloat()-f2) > 1e-7 {
		return fmt.Errorf("invalid amount: %v, %v", err, amount)
	}
	if token == "iost" && f1.Decimal > 8 {
		return fmt.Errorf("invalid decimal: %v", amount)
	}
	return nil
}

func checkBadAction(action *tx.Action) error {
	if action.Contract == "token.iost" && action.ActionName == "transfer" {
		data := action.Data
		js, err := simplejson.NewJson([]byte(data))
		if err != nil {
			return fmt.Errorf("invalid json array: %v, %v", err, data)
		}
		arr, err := js.Array()
		if err != nil {
			return fmt.Errorf("invalid json array: %v, %v", err, data)
		}
		if len(arr) != 5 {
			return fmt.Errorf("wrong args num: %v", data)
		}
		token, err := js.GetIndex(0).String()
		if err != nil {
			return fmt.Errorf("invalid token: %v, %v", err, data)
		}
		amount, err := js.GetIndex(3).String()
		if err != nil {
			return fmt.Errorf("invalid amount: %v, %v", err, data)
		}
		err = checkAmount(amount, token)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func checkBadTx(tx *tx.Tx) error {
	for _, a := range tx.Actions {
		err := checkBadAction(a)
		if err != nil {
			return err
		}
	}
	return nil
}
