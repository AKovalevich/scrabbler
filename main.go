package main

import (
	//"os"
	//"github.com/AKovalevich/scrabbler/cmd/scrabbler"
	"reflect"
	"github.com/AKovalevich/scrabbler/config"
	"github.com/BurntSushi/toml"
	log "github.com/AKovalevich/scrabbler/log/logrus"
	"fmt"
)

func main() {
	//os.Exit(scrabbler.Run(os.Args[1:]))
	tmpMain()
}

func tmpMain() {
	var tmpScreabblerConfiguration config.ScrabblerConfiguration

	testConfig := config.ScrabblerConfiguration{
		ServerHost: "test",
	}

	cfg := "configuration.default.toml"

	if _, err := toml.DecodeFile(cfg, &tmpScreabblerConfiguration); err != nil {
		log.Do.Error(err)
	}

	tmpConfigStructureValues := reflect.ValueOf(&tmpScreabblerConfiguration).Elem()
	tmpConfigStructureTypes := reflect.TypeOf(tmpConfigStructureValues)
	realConfigStructureValues := reflect.ValueOf(&testConfig).Elem()
	realConfigStructureTypes := reflect.TypeOf(testConfig)

	for indexRealConfigStructure := 0; indexRealConfigStructure < realConfigStructureValues.NumField(); indexRealConfigStructure++ {
		//fmt.Println(tmpConfigStructure.Index(index))
		//fmt.Println(realConfigStructure.Index(index))
		//name := structType.Field(index).Name
		//fmt.Println(structType.Field(index).Tag.Get("default"))
		fmt.Println("Name: " + realConfigStructureTypes.Field(indexRealConfigStructure).Name + " | Value: " + realConfigStructureValues.Field(indexRealConfigStructure).String())

		for indexTmpConfigStructure := 0; indexTmpConfigStructure < tmpConfigStructureValues.NumField(); indexTmpConfigStructure++ {
			if tmpConfigStructureTypes.Field(indexTmpConfigStructure).Name == realConfigStructureTypes.Field(indexRealConfigStructure).Name {

			}
		}


		//f := tmpConfigStructure.Field(index)
		//f2:= realConfigStructure.Field(index)
		//
		//
		//
		//c1 := f.Interface()
		//c2 := f2.Interface()
		//
		//if c1 ==
		//
		//fmt.Print(c1)
		//fmt.Print(c2)
	}
}
