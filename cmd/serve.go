package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/ethan-rng/zero/internals/daemon"
	"github.com/ethan-rng/zero/internals/web"
	"github.com/spf13/cobra"
)



var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "launches basic gui on port 6781 and launches the zero daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		headless, _ := cmd.Flags().GetBool("headless")
		dev, _ := cmd.Flags().GetBool("dev")

		if headless && dev {
			return fmt.Errorf("cannot use headless and dev mode together")
		}

		// 1. Start the GUI in a goroutine so it doesn't block the daemon
		if !headless {
			go func() {
				files, err := fs.Sub(web.Content, "out")
				if err != nil {
					log.Fatal(err)
				}
				http.Handle("/", http.FileServer(http.FS(files)))
				fmt.Println("GUI running at http://localhost:6781")
				if err := http.ListenAndServe(":6781", nil); err != nil {
					log.Fatal(err)
				}
			}()
		}

		// 2. Logic for the daemon
		locally, _ := cmd.Flags().GetBool("local")
		
		srv := daemon.NewServer(":8080", locally)
		if err := srv.Start(); err != nil {
			return fmt.Errorf("daemon server failed: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolP("headless", "H", false, "run in headless mode (no gui)")
	serveCmd.Flags().BoolP("local", "L", false, "run locally")
	serveCmd.Flags().BoolP("dev", "D", false, "run in dev mode")
}
