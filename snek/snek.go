package snek

import (
	"fmt"
	"os"

	"github.com/amy911/amy911/onfail"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	InitCopyright func(*cobra.Command) error
	InitEula func(*cobra.Command) error
	InitLegal func(*cobra.Command) error
	InitLicense func(*cobra.Command) error
	InitRoot func(*cobra.Command) error
	InitVersion func(*cobra.Command) error
)

func Main(onFail ...onfail.OnFail) error {
	copyrightCmd := &cobra.Command{
		Use: "copyright",
		Short: "Print the copyright",
		Long: `Print the copyright`,
		Run: func(cmd *cobra.Command, args []string) {
			copyright := NewCopyright(CopyrightFirstYear, CopyrightHolder)
			opath := pflag.GetString("out")
			out := os.Stdout
			if opath != "-" {
				var err error
				if out, err = os.Create(opath); err != nil {
					onfail.Fail(err, opath, onfail.Fatal, onFail)
				}
				defer out.Close()
			}
			switch {
			case pflag.GetBool("json"):
				fmt.Fprintln(copyright.Json())
			case pflag.GetBool("xml"):
				fmt.Fprintln(copyright.Xml())
			default:
				fmt.Fprintln(copyright.Robots)
			}
		},
	}
	if InitCopyright != nil {
		if err := InitCopyright(copyrightCmd); err != nil {
			return onfail.Fail(err, copyrightCmd, onfail.Print, onFail)
		}
	}
	eulaCmd := &cobra.Command{
		Use: "eula",
		Short: "Print the End User License Agreement (EULA)",
		Long: `Print the End User License Agreement (EULA)`,
		Run: func(cmd *cobra.Command, args []string) {
			eula := NewLegal(NewCopyright(CopyrightFirstYear, CopyrightHolder), License, Eula)
			opath := pflag.GetString("out")
			out := os.Stdout
			if opath != "-" {
				var err error
				if out, err = os.Create(opath); err != nil {
					onfail.Fail(err, opath, onfail.Fatal, onFail)
				}
				defer out.Close()
			}
			switch {
			case pflag.GetBool("json"):
				fmt.Fprintln(eula.Json())
			case pflag.GetBool("xml"):
				fmt.Fprintln(eula.Xml())
			default:
				fmt.Fprintln(eula.Robots)
			}
		},
	}
	if InitEula != nil {
		if err := InitEula(eulaCmd); err != nil {
			return onfail.Fail(err, eulaCmd, onfail.Print, onFail)
		}
	}
	legalCmd := &cobra.Command{
		Use: "legal",
		Short: "Print the End User License Agreement (EULA)",
		Long: `Print the End User License Agreement (EULA)`,
		Run: func(cmd *cobra.Command, args []string) {
			legal := NewLegal(NewCopyright(CopyrightFirstYear, CopyrightHolder), License, Eula)
			opath := pflag.GetString("out")
			out := os.Stdout
			if opath != "-" {
				var err error
				if out, err = os.Create(opath); err != nil {
					onfail.Fail(err, opath, onfail.Fatal, onFail)
				}
				defer out.Close()
			}
			switch {
			case pflag.GetBool("json"):
				fmt.Fprintln(legal.Json())
			case pflag.GetBool("xml"):
				fmt.Fprintln(legal.Xml())
			default:
				fmt.Fprintln(legal.Robots)
			}
		},
	}
	if InitLegal != nil {
		if err := InitLegal(legalCmd); err != nil {
			return onfail.Fail(err, legalCmd, onfail.Print, onFail)
		}
	}
	licenseCmd := &cobra.Command{
		Use: "license",
		Short: "Print the End User License Agreement (EULA)",
		Long: `Print the End User License Agreement (EULA)`,
		Run: func(cmd *cobra.Command, args []string) {
			license := NewLegal(NewCopyright(CopyrightFirstYear, CopyrightHolder), License, Eula)
			opath := pflag.GetString("out")
			out := os.Stdout
			if opath != "-" {
				var err error
				if out, err = os.Create(opath); err != nil {
					onfail.Fail(err, opath, onfail.Fatal, onFail)
				}
				defer out.Close()
			}
			switch {
			case pflag.GetBool("json"):
				fmt.Fprintln(license.Json())
			case pflag.GetBool("xml"):
				fmt.Fprintln(license.Xml())
			default:
				fmt.Fprintln(license.Robots)
			}
		},
	}
	if InitLicense != nil {
		if err := InitLicense(licenseCmd); err != nil {
			return onfail.Fail(err, licenseCmd, onfail.Print, onFail)
		}
	}
	versionCmd := &cobra.Command{
		Use: "version",
		Short: "Print the version",
		Long: `Print the version`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Version)
		},
	}
	if InitVersion != nil {
		if err := InitVersion(versionCmd); err != nil {
			return onfail.Fail(err, versionCmd, onfail.Print, onFail)
		}
	}
	rootCmd := &cobra.Command{Use: os.Args[0]}
	rootCmd.PersistentFlags().StringP("out", "o", "-", "Output to this file (or \"-\" for stdout)")
	rootCmd.PersistentFlags().Bool("json", false, "Output in json")
	rootCmd.PersistentFlags().Bool("xml", false, "Output in xml")
	if InitRoot != nil {
		if err := InitRoot(rootCmd); err != nil {
			return onfail.Fail(err, rootCmd, onfail.Print, onFail)
		}
	}
	rootCmd.AddCommand(copyrightCmd)
	rootCmd.AddCommand(eulaCmd)
	rootCmd.AddCommand(legalCmd)
	rootCmd.AddCommand(licenseCmd)
	rootCmd.AddCommand(versionCmd)
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
