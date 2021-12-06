package validate

import (
	"errors"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func IdValidate(idAccount string) (int64, error) {
	id, err := strconv.ParseInt(idAccount, 0, 64)
	if err != nil {
		log.WithError(err).Errorf("Error with parcing id")
		return 0, errors.New("error with parcing id")
	}
	if id < 1 {
		log.Errorf("Incorrect value id user")
		return 0, errors.New("incorrect value id user")
	}

	return id, nil
}

func AmountValidate(amountS string) (float64, error) {
	validAmount := strings.Split(amountS, ".")
	if len(validAmount) > 1 {
		if len(validAmount[1]) > 2 {
			log.Errorf("The amount have more then 2 decimal places")
			return 0, errors.New("the amount have more then 2 decimal places")
		}
	}

	amount, err := strconv.ParseFloat(amountS, 64)
	if err != nil {
		log.WithError(err).Errorf("Error with parcing amount")
		return 0, errors.New("error with parcing amount")
	}

	if amount < 0.01 {
		log.Errorf("The amount is negative")
		return 0, errors.New("the amount is negative")
	}

	return amount, nil
}
