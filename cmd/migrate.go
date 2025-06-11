package cmd

import (
	"log"

	"github.com/igorfranzoi/golibfunctions/database/infrastructure"
	"github.com/igorfranzoi/golibfunctions/utils"
	"github.com/spf13/cobra"
)

var typeCommand string

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Execute all of models to create on database tables",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logFile, connDatabase, err := utils.InitEnviroment()

		if err != nil {
			if logFile != nil {
				utils.WriteLog(logFile, utils.InfoLevel, "abort with abnormal termination", err)
				return
			} else {
				log.Fatalf("Error create log...")
			}
		}

		if connDatabase == nil {
			if logFile != nil {
				utils.WriteLog(logFile, utils.InfoLevel, "connection database can´t be nil", err)
				return
			} else {
				log.Fatalf("Error create log...")
			}
		}

		execMethod := ""

		switch typeCommand {
		case "up":
			execMethod = "MigrateApply"
		case "down":
			execMethod = "MigrateRevert"
		default:
			log.Println("command not recognized")
			return
		}

		utils.WriteLog(logFile, utils.InfoLevel, "starting application of migration")

		infrastructure.RunningMigrate(execMethod)

		utils.WriteLog(logFile, utils.InfoLevel, "finish application of migration")

	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.
	migrateCmd.Flags().StringVarP(&typeCommand, "typeCommand", "o", "up", "(up) - aplica as atualizações / (down) - remover alterações no banco")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
