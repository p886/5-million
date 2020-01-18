package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/p886/5-million-steps/calculator"
	"github.com/p886/5-million-steps/representer"
	"github.com/spf13/cobra"
)

// ServerCmd starts an HTTP server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Exposes an endpoint for the steps calculation",
	Long:  `Call / with the URL parameter 'steps' to get the current statistics`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		listenAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))
		log.Printf("Listening on %s", listenAddr)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			steps, ok := r.URL.Query()["steps"]
			if !ok || len(steps[0]) < 1 {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Missing URL parameter 'steps'"))
				log.Println("Missing URL parameter 'steps'")
				return
			}
			currentSteps, err := strconv.Atoi(steps[0])
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				log.Printf("Error converting steps URL param: %v", err)
				return
			}
			result := calculator.Calculate(currentSteps)
			jsonRepresentation, err := representer.AsJSON(result)
			if err != nil {
				log.Printf("Error representing JSON: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(jsonRepresentation))
		})
		http.ListenAndServe(listenAddr, nil)
	},
}
