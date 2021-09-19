package main

import (
	"github.com/gagliardetto/solana-go"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	desiredPrefix := "sol"  //don't put more than 3 unless you want to wait.... ALOT

	start := time.Now()
	var ops uint64

	var wg sync.WaitGroup

	for {
		for i := 0; i < 150; i++ {
			wg.Add(1)

			go func() {
				for c := 0; c < 40000; c++ {

					// Create a new account:
					account := solana.NewWallet()

					if strings.HasPrefix(account.PublicKey().String(), desiredPrefix) {
						log.Println("-------------------------")
						log.Printf("public key: %s [private key: %s]", account.PublicKey().String(), account.PrivateKey)
					}

					atomic.AddUint64(&ops, 1)
				}
				wg.Done()
			}()
		}
		wg.Wait()
		elapsed := time.Since(start)
		log.Printf("searched of %d wallets in %s", ops, elapsed)
	}

	//cancel the script with CTRL+C
}
