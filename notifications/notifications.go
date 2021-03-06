package notifications

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

type Level int

func (l Level) String() string {
	switch l {
	case 0:
		return "danger"
	case 1:
		return "warning"
	case 2:
		return "info"
	}
	return ""
}

const (
	DANGER  Level = 0
	WARNING Level = 1
	INFO    Level = 2
)

type NotificationLink struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

type Notification struct {
	Key       string
	Title     string
	Level     Level
	Message   string
	Timestamp time.Time
	Link      NotificationLink
}

func (n Notification) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"title":     n.Title,
		"message":   n.Message,
		"level":     n.Level.String(),
		"timestamp": n.Timestamp.Format(time.RFC3339),
	}
	if n.Link.URL != "" {
		m["link"] = n.Link
	}
	return json.Marshal(m)
}

func List(ctx context.Context) []Notification {
	var (
		res = make([]Notification, 0)
		ch  = make(chan []Notification)
	)
	defer close(ch)
	checkers := []func(chan []Notification){
		//CheckURP, CheckPluginVersion,
		//CheckBalancedLeaders, CheckISRDelta,
	}
	for _, fn := range checkers {
		go fn(ch)
	}
	for i := 0; i < len(checkers); i++ {
		select {
		case notfis := <-ch:
			res = append(res, notfis...)
		case <-ctx.Done():
			fmt.Fprintf(os.Stderr, "[INFO] Notifications timed out: %s\n", ctx.Err())
			break
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Level < res[j].Level
	})
	return res
}
