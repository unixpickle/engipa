package main

import (
	"errors"
	"fmt"
	"github.com/unixpickle/engipa/fetcher"
	"github.com/unixpickle/engipa/wordlist"
	"os"
	"sync"
)

const NumRoutines = 16

func main() {
	if err := ErrMain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ErrMain() error {
	if len(os.Args) != 3 {
		return errors.New("Usage: fetch-lines <input.txt> <output.txt>")
	}
	lines, err := wordlist.ReadLines(os.Args[1])
	if err != nil {
		return err
	}
	m := FindIPAWords(lines)
	return wordlist.Mapping(m).WriteFile(os.Args[2])
}

func FindIPAWords(words []string) map[string]string {
	res := map[string]string{}
	numErrors := 0
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	ch := make(chan string)
	// Launch background goroutines
	for i := 0; i < NumRoutines; i++ {
		wg.Add(1)
		go func() {
			for {
				word, ok := <-ch
				if !ok {
					break
				}
				ipa, err := fetcher.GuessIPA(word)
				lock.Lock()
				if err != nil {
					numErrors++
				} else {
					res[word] = ipa
				}
				lock.Unlock()
			}
			wg.Done()
		}()
	}
	// Push words to the queue
	for i, word := range words {
		ch <- word
		if i%500 == 0 {
			lock.Lock()
			fmt.Println("Up to", i, "with", numErrors, "errors")
			lock.Unlock()
		}
	}
	close(ch)
	wg.Wait()
	fmt.Println("Done with", numErrors, "errors")
	return res
}
