package bluetooth

import (
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

type Config struct {
	Duration time.Duration
	Debug    bool
}

type Scanner struct {
	config *Config
}

const (
	DEFAULT_DURATION = 10 * time.Second
	DEFAULT_DEBUG    = false
)

func NewScanner(cfg ...*Config) *Scanner {
	var config *Config

	if len(cfg) > 0 && cfg[0] != nil {
		config = cfg[0]
	} else {
		config = &Config{
			Duration: DEFAULT_DURATION,
			Debug:    DEFAULT_DEBUG,
		}
	}

	if config.Duration <= 0 {
		config.Duration = DEFAULT_DURATION
	}

	return &Scanner{config}
}

func (s *Scanner) Start() chan Event {
	if err := adapter.Enable(); err != nil {
		panic(err)
	}

	duration := s.config.Duration

	events := make(chan Event)

	go func() {
		defer close(events)

		if err := adapter.Scan(func(adapter *bluetooth.Adapter, event bluetooth.ScanResult) {
			events <- Event{
				Address:   event.Address.String(),
				Name:      event.LocalName(),
				RSSI:      event.RSSI,
				CreatedAt: time.Now(),
			}
		}); err != nil {
			panic(err)
		}
	}()

	go func() {
		time.Sleep(duration)
		adapter.StopScan()
	}()

	return events
}

type Event struct {
	Address   string    `json:"address"`
	Name      string    `json:"name"`
	RSSI      int16     `json:"rssi"`
	CreatedAt time.Time `json:"created_at"`
}
