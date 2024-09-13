/*
Copyright © 2024 Kaleb Hawkins <KalebHawkins@outlook.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/KalebHawkins/fitquest/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	id   int
	name string
	goal int
	reps int

	defaultCfgFileName = ".fitquest.json"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fitquest",
	Short: "Fitquest is designed to help track fitness goals in a simple way.",
	Long: `Fitquest is designed to help track fitness goals in a simple way.
To use fitquest you can start by adding exercises. To see how to do that run the following command:

> fitquest track --help

This will provide the help for the track subcommand.
To view your tracked exercises you can run:

> fitquest

This will display a table of your tracked exercises. If you want more
information about a tracked exercise you can use the command:

fitquest --id <id>

Replace <id> with the id of the exercise you want details on.

The datafile that contains all of your saved data is local to your machine at ~/.fitquest.json.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := loadConfig()
		if err != nil {
			return err
		}
		defer f.Close()

		log, err := LoadLog(f)
		if err != nil {
			return err
		}

		id -= 1

		// If no ID is specified and a name is add the exercise for tracking
		if id < 0 && name != "" {
			if reps == 0 {
				log.Add(&types.Exercise{
					Name:    name,
					Goal:    goal,
					Session: make([]*types.Session, 0),
				})
			}

			if reps > 0 {
				log.Add(&types.Exercise{
					Name:    name,
					Goal:    goal,
					Session: []*types.Session{types.NewSession(reps)},
				})
			}

			SaveConfig(f, log)
			fmt.Println(log)
		}

		// Print the log if no id is specified.
		// Print the log if the id is out of range.
		if (id < 0 || id >= len(log.Exercises)) && name == "" {
			fmt.Println(log)
			return nil
		}

		// If the ID is in range we update whatever.
		if id >= 0 && id < len(log.Exercises) {
			if name != "" {
				log.Exercises[id].Name = name
			}
			if goal != 0 {
				log.Exercises[id].Goal = goal
			}
			if reps != 0 {
				log.Exercises[id].Session = append(log.Exercises[id].Session, types.NewSession(reps))
			}

			if name == "" && goal == 0 {
				SaveConfig(f, log)
				fmt.Println(log.Exercises[id])
				return nil
			}

			SaveConfig(f, log)
			fmt.Println(log)
			return nil
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().IntVar(&id, "id", 0, "id of exercise")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of the exercise")
	rootCmd.PersistentFlags().IntVarP(&reps, "reps", "r", 0, "number of reps to log")
	rootCmd.PersistentFlags().IntVarP(&goal, "goal", "g", 0, "the number of reps to reach")
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	viper.AddConfigPath(home)
	viper.SetConfigType("json")
	viper.SetConfigName(".fitquest")
	viper.SetConfigFile(filepath.Join(home, defaultCfgFileName))
}

func loadConfig() (io.ReadWriteCloser, error) {
	f, err := os.OpenFile(viper.ConfigFileUsed(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func LoadLog(r io.Reader) (*types.Log, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return types.NewLog(), nil
	}

	var log types.Log
	if err := json.Unmarshal(data, &log); err != nil {
		return nil, fmt.Errorf("failed to parse log: %q", err)
	}

	return &log, err
}

func SaveConfig(w io.Writer, log *types.Log) error {
	os.Truncate(viper.ConfigFileUsed(), 0)

	data, err := json.MarshalIndent(log, "", "  ")
	if err != nil {
		return err
	}

	w.Write(data)

	return nil
}
