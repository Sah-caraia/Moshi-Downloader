package main

import (
	"bufio"
	"fmt"
	"moshi/internal/commands"
	"moshi/internal/texts"
	"moshi/internal/utils"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	utils.Clear() // clear the terminal
	prompt := promptui.Select{ // show the options
		Label: texts.PortugueseDefault.SelectAnOption,
		Items: []string{ "mp3", "mp4" },
		Templates: &promptui.SelectTemplates{
			Label: "{{ . | blue }}",
			Selected: "{{ . | blue | bold }}",
			Active: "{{ . | green | bold }}",
			Inactive: "{{ . | white }}",
			Help: " ",
		},
	}
	_, filetype, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		return
	}


	urls := []string{} // slice of URLs 


	var mediaDownloadCount int32 = 0 
	urlCount := 0
	for {
		utils.Clear()
		fmt.Println("🔗 URLs: ", urlCount)
		fmt.Println("🌷 ", texts.PortugueseDefault.SendYourUrl)
		fmt.Println("💡 ", texts.PortugueseDefault.InputHint, "\n")
		fmt.Print("\033[36m>>> \033[0m")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if !strings.HasPrefix(strings.ToLower(input), "https:") {
			if input == "END" {
				break
			} else {
				fmt.Println(texts.PortugueseErrors.InvalidUrl)
				time.Sleep(time.Second * 2)
				continue
			}
		} else {
			urls = append(urls, input)
			urlCount += 1
			continue
		}
	}


	var wg sync.WaitGroup 
	errorChan := make(chan error)
	for _, url := range urls { // run goroutines
		wg.Add(1)
		go func(u string){
			defer wg.Done()
			err := commands.DownloadMedia(u, filetype)
			errorChan <- err
			atomic.AddInt32(&mediaDownloadCount, 1)
		}(url)
	}

	go func(){
		wg.Wait()
		close(errorChan)
	}()
	for err :=  range errorChan { // to show errors
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		} else {
			fmt.Println(texts.PortugueseSuccess.MediaInstalled)
		}
	}
	if mediaDownloadCount >= 1 {
		fmt.Println(texts.PortugueseSuccess.AllSuccess)
	}
}
