package cmd

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	url         string
	totalReqs   int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "go-stress-test",
	Short: "Teste de estresse",
	Long: `Essa aplicação realiza testes de estresse para uma url especificada. Exemplo: 
—url=http://google.com —requests=1000 —concurrency=10
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			fmt.Println("Error: --url é obrigatório")
			cmd.Usage()
			os.Exit(1)
		}
		if totalReqs <= 0 {
			fmt.Println("Error: --requests deve ser > 0")
			cmd.Usage()
			os.Exit(1)
		}
		if concurrency <= 0 {
			fmt.Println("Error: --concurrency deve ser > 0")
			cmd.Usage()
			os.Exit(1)
		}

		runStressTest(url, totalReqs, concurrency)
	},
}

func runStressTest(url string, reqs, conc int) {
	start := time.Now()

	fmt.Print("Processando...")

	var wg sync.WaitGroup
	mx := sync.Mutex{}

	var sem = make(chan struct{}, conc)

	var errors int
	var successCount int
	var completed int
	var statusCodes = make(map[int]int)

	for i := 0; i < reqs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}

			resp, err := http.Get(url)

			// lock para evitar race conditions ao incrementar variaveis
			mx.Lock()
			completed++

			if err != nil {
				errors++
			} else {
				statusCodes[resp.StatusCode]++
				if resp.StatusCode == 200 {
					successCount++
				}
				resp.Body.Close()
			}

			mx.Unlock()

			<-sem
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Print("\r\033[K")

	fmt.Println("==== Relatório de Teste de Carga ====")
	fmt.Printf("Tempo total gasto: %v\n", duration)
	fmt.Printf("Total de requests realizados: %d\n", completed)
	fmt.Printf("Requests com status HTTP 200: %d\n", successCount)
	fmt.Println("Distribuição de códigos de status:")
	for k, v := range statusCodes {
		label := fmt.Sprintf("%d", k)
		if k >= 400 {
			label = "erro"
		}
		fmt.Printf("  %s: %d\n", label, v)
	}

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "URL do serviço a ser testado")
	rootCmd.PersistentFlags().IntVarP(&totalReqs, "requests", "r", 0, "Número total de requests")
	rootCmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", 1, "Número de chamadas simultâneas")
}
