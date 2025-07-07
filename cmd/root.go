package cmd

import (
	"fmt"
	"os"

	"payment/internal/server"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "payment",
	Short: "payment is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gopayment.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewDevelopment()
		sugarLogger := logger.Sugar()
		router := gin.Default()
		srv := server.InitServer(9090, router, sugarLogger)

		err := srv.StartServer()
		if err != nil {
			fmt.Println("Failed Initserver %+v", err)
		}
		_ = logger.Sync()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
