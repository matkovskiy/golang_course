package pub_sub

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Subscriber chan string

type Publisher struct {
	Subscribers map[Subscriber]struct{}
}

func (p *Publisher) AddSubscriber(sub Subscriber) {
	p.Subscribers[sub] = struct{}{}
}

func (p *Publisher) RemoveSubscriber(sub Subscriber) {
	delete(p.Subscribers, sub)
	close(sub)
}

func (p *Publisher) Notify(message string) {
	for sub := range p.Subscribers {
		sub <- message
	}
}

func DetectChangesINDirectory(p *Publisher, directory string) {
	for {
		err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				p.Notify(fmt.Sprintf("Change in : %s", path))
			}
			return nil
		})
		if err != nil {
			fmt.Println("error:", err)
		}
		time.Sleep(5 * time.Second)
	}
}
