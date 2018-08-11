package snek

import (
	"fmt"

	"github.com/amy911/amy911/onfail"

	"github.com/spf13/cobra"
)

func Bind(cmd *cobra.Command, args ...interface{}) {
	var binds []string
	for _, arg := range args {
		switch arg.(type) {
		case string:
			binds = append(binds, arg.(string))
		case onfail.OnFail:
		default:
			panic(fmt.Sprintf("Arguments must be string or github.com/amy911/amy911/onfail.OnFail, not %T", arg))
		}
	}
	for _, bind := range binds {
		BindOne(bind, cmd.Flag(bind), args)
	}
}

func BindOne(configName string, flag *pflag.Flag, onFail ...interface{}) error {
	if err := viper.BindPFlag(configName, flag); err != nil {
		return onfail.Fail(err, flag, onfail.Print, onFail)
	}
	viperPFlagBindings = append(viperPFlagBindings, viperPFlagBinding(configName, flag.Value))
	return nil
}
