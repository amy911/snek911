package snek

import (
	"github.com/amy911/amy911/onfail"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Bind(configName string, flag *pflag.Flag, onFail ...onfail.OnFail) error {
	if err := viper.BindPFlag(configName, flag); err != nil {
		return onfail.Fail(err, flag, onfail.Print, onFail)
	}
	viperPFlagBindings = append(viperPFlagBindings, viperPFlagBinding(configName, flag.Value))
	return nil
}

var Init func(*cobra.Command)

func Main(rootCmd *cobra.Command, onFail ...onfail.OnFail) {
	if rootCmd == nil {
		rootCmd = &cobra.Command{}
	}
	if Init != nil {
		Init(rootCmd)
	}
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return onfail.Fail(err, cmd, onfail.Print, onFail)
			}
		}
		for _, v := range viperPFlagBindings {
			v.flagValue.Set(viper.GetString(v.configName))
		}
		return nil
	}
}

///////////////////////////////////////////////////////////////////////////////

type viperPFlagBinding struct {
	configName string
	flagValue  pflag.Value
}
var viperPFlagBindings []viperPFlagBinding

// Derivative of [snippet](https://github.com/spf13/viper/issues/82#issuecomment-403165788) by [WGH-](https://github.com/WGH-)
// Currently awaiting permission to use in this repo
