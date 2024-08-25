package logging

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jsdbroughton/speckle-go/pkg/speckle/core"
	"io"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var (
	track          = TrackDefault
	hostApp        = HostAppDefault
	hostAppVersion = HostAppVersionDefault
	metricsTracker *MetricsTracker
	once           sync.Once
	metricsLogger  = log.New(log.Writer(), "METRICS: ", log.Ldate|log.Ltime)
)

type MetricsTracker struct {
	lastUser       string
	lastServer     string
	analyticsUrl   string
	analyticsToken string
	platform       string
	queue          chan map[string]interface{}
	httpClient     *http.Client
}

func Disable() {
	track = false
}

func Enable() {
	track = true
}

func SetHostApp(app string, version string) {
	hostApp = app
	if version != "" {
		hostAppVersion = version
	}
}

func Track(action string, account *core.Account, customProps map[string]interface{}) {
	if !track {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			metricsLogger.Printf("Recovered in Track: %v", r)
		}
	}()

	InitializeTracker(account)

	eventParams := map[string]interface{}{
		"event": action,
		"properties": map[string]interface{}{
			"distinct_id":    metricsTracker.lastUser,
			"server_id":      metricsTracker.lastServer,
			"token":          metricsTracker.analyticsToken,
			"hostApp":        hostApp,
			"hostAppVersion": hostAppVersion,
			"$os":            metricsTracker.platform,
			"type":           "action",
		},
	}

	if customProps != nil {
		for k, v := range customProps {
			eventParams["properties"].(map[string]interface{})[k] = v
		}
	}

	select {
	case metricsTracker.queue <- eventParams:
		// Successfully queued the event
	default:
		metricsLogger.Printf("Metrics queue is full, dropping event: %s", action)
	}
}

func InitializeTracker(account *core.Account) {
	once.Do(func() {
		metricsTracker = &MetricsTracker{
			platform:       Platforms[runtime.GOOS],
			analyticsUrl:   AnalyticsUrl,
			analyticsToken: AnalyticsToken,
			queue:          make(chan map[string]interface{}, 1000),
		}
		go metricsTracker.processQueue()
	})

	if account != nil && account.UserInfo.Email != "" {
		metricsTracker.lastUser = account.UserInfo.Email
	}
	if account != nil && account.ServerInfo.URL != "" {
		metricsTracker.lastServer = account.ServerInfo.URL
	}
}

func (m *MetricsTracker) processQueue() {
	for eventParams := range m.queue {
		err := m.send(eventParams)
		if err != nil {
			metricsLogger.Printf("Error sending to Mixpanel: %v", err)
		}
	}
}

func (m *MetricsTracker) send(eventParams map[string]interface{}) error {
	jsonData, err := json.Marshal(eventParams)
	if err != nil {
		return fmt.Errorf("error marshaling event data: %w", err)
	}

	encodedData := base64.StdEncoding.EncodeToString(jsonData)
	body := []byte(fmt.Sprintf("data=%s", encodedData))

	for i := 0; i < MaxRetries; i++ {
		err := func() error {
			req, err := http.NewRequest("POST", AnalyticsUrl, bytes.NewBuffer(body))
			if err != nil {
				return fmt.Errorf("error creating request: %w", err)
			}

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			resp, err := m.httpClient.Do(req)
			if err != nil {
				return fmt.Errorf("error sending request: %w", err)
			}
			defer func() {
				if closeErr := resp.Body.Close(); closeErr != nil {
					metricsLogger.Printf("Error closing response body: %v", closeErr)
				}
			}()

			if resp.StatusCode != http.StatusOK {
				bodyBytes, _ := io.ReadAll(resp.Body)
				return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
			}

			return nil
		}()

		if err == nil {
			return nil // Successfully sent to Mixpanel
		}

		if i == MaxRetries-1 {
			return fmt.Errorf("failed to send to Mixpanel after %d retries: %w", MaxRetries, err)
		}

		metricsLogger.Printf("Retrying Mixpanel send (attempt %d/%d): %v", i+1, MaxRetries, err)
		time.Sleep(RetryDelay * time.Duration(i+1)) // Exponential backoff
	}

	return fmt.Errorf("failed to send to Mixpanel after %d retries", MaxRetries)
}
