package calculations

import (
	"fmt"
	"strconv"
)

func (u *user) getEndLvlAndExp() error {
	var cont = false
	for cont == false {
		err := u.getEndingLvl()
		if err != nil {
			handleInputError(err)
		} else {
			cont = true
		}
	}
	cont = false
	for cont == false {
		err := u.getEndingExp()
		if err != nil {
			handleInputError(err)
		} else {
			cont = true
		}
	}
	fmt.Printf("Your ending level is: %d and %.2f", u.endLvl, u.endExp)
	fmt.Println("%")
	return nil
}

func (u *user) getEndingLvl() error {
	fmt.Println("Input ending level: ")
	input := getInput(u.scanner)
	output, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return err
	}
	if uint8(output) < u.startLvl {
		return fmt.Errorf("Ending lvl cannot be lower than starting lvl")
	}
	u.endLvl = uint8(output)
	return nil
}

func (u *user) getEndingExp() error {
	fmt.Println("Input ending exp %: ")
	input := getInput(u.scanner)
	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return err
	}
	if u.startLvl == u.endLvl {
		if float32(output) < u.startExp {
			fmt.Println("So... You lost exp?")
		}
	}
	u.endExp = float32(output)
	return nil
}
