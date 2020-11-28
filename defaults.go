//Package enviper provides functions for setting viper defaults variables from ENV
package enviper

import (
	"github.com/spf13/viper"
	"log"
	"strconv"
)

//SetDefaultString sets to viper string env variable or defaultValue
func SetDefaultString(name string, defaultValue string) {
	value := getFromEnv(name, defaultValue)
	viper.SetDefault(name, value)
}

//SetDefaultInt sets to viper int env variable or defaultValue
func SetDefaultInt(name string, defaultValue int) {
	value := getFromEnv(name, strconv.Itoa(defaultValue))
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Variable %s is not int", name)
	}
	viper.SetDefault(name, intValue)
}

//SetDefaultBool sets to viper boolean env variable or defaultValue
//From env variable expects integer. 0 - false, another value is true
func SetDefaultBool(name string, defaultValue bool) {
	strValue := "0"
	if defaultValue {
		strValue = "1"
	}
	value := getFromEnv(name, strValue)
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Variable %s is not int", name)
	}
	if intValue == 0 {
		viper.SetDefault(name, false)
	} else {
		viper.SetDefault(name, true)
	}
}
