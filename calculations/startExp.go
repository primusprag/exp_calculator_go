package calculations

import (
	"fmt"
	"strconv"


)

func (u *user) getStartLvlAndExp() error {
	var cont = false
	for cont == false {
		err := u.getStartingLvl()
		if err != nil {
			handleInputError(err)
		} else {
			cont = true
		}
	}
	cont = false
	for cont == false {
		err := u.getStartingExp()
		if err != nil {
			handleInputError(err)
		} else {
			cont = true
		}
	}
	fmt.Printf("Your starting level is: %d and %.2f", u.startLvl, u.startExp)
	fmt.Println("%")
	return nil
}

func (u *user) getStartingLvl() error {
	fmt.Println("Input starting level: ")
	input := getInput(u.scanner)
	output, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return err
	}
	u.startLvl = uint8(output)
	return nil
}

func (u *user) getStartingExp() error {
	fmt.Println("Input starting exp %: ")
	input := getInput(u.scanner)
	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return err
	}
	u.startExp = float32(output)
	return nil
}
