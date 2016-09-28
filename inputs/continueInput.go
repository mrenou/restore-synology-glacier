package inputs

import (
	"rsg/outputs"
	"rsg/utils"
	"rsg/consts"
)

func QueryContinue() {
	outputs.Print(outputs.Info, "Press to continue...")
	for _, _ = range consts.LINE_BREAK {
		_, err := StdinReader.ReadByte()
		utils.ExitIfError(err)
	}

}
