package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config contains path to configuration file and allows to read it.
type Config struct {
	File  string
	viper *viper.Viper
	flags []*pflag.FlagSet
}

// Read configuration.
func (c *Config) Read() error {
	c.init()
	if c.File != "" {
		// Use config file from the flag.
		c.viper.SetConfigFile(c.File)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			return err
		}

		// Search config in home directory with name ".c" (without extension).
		c.viper.AddConfigPath(home)
		c.viper.SetConfigName(".pmm-agent")
	}

	c.viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := c.viper.ReadInConfig()
	if err != nil {
		return err
	}

	// Set all flags.
	for i := range c.flags {
		c.flags[i].VisitAll(func(f *pflag.Flag) {
			f.Value.Set(c.viper.GetString(f.Name))
		})
	}

	return nil
}

// Bind flags to configuration.
func (c *Config) Bind(cmd *cobra.Command) {
	c.init()
	flags := cmd.Flags()
	c.flags = append(c.flags, flags)
	c.viper.BindPFlags(flags)
}

func (c *Config) init() {
	if c.viper == nil {
		c.viper = viper.New()
	}
}

// Flags assigns flags to cmd.
func (c *Config) Flags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&c.File, "config", "", "config File (default is $HOME/.pmm-agent.yaml)")
}
