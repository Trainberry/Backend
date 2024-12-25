package services

import(
	"server/internal/cache"
	"time"
	"github.com/sirupsen/logrus"
	"net/http"
	"fmt"
	"sync"
)

func init() {
	go scheduler()
}


func scheduler() {
	var wg sync.WaitGroup
	for {
		time.Sleep(time.Second * 30)
		logrus.Info("Refreshing cache")

		for _, k := range cache.Map.Keys() {
			item, ok := cache.Map.Get(k)
			if !ok {
				continue
			}

			wg.Add(1)

			go func() {
				// Create request
				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/ping", item.IP), nil)
				if err != nil {
					logrus.Errorf("Could not create request to ping train %s : %v", k, err)
					wg.Done()
					return
				}

				// Send request
				res, err := http.DefaultClient.Do(req)
				if err != nil {
					logrus.Errorf("Unable to send request to ping train %s : %v", k, err)
					wg.Done()
					return
				}

				if res.StatusCode != 204 {
					logrus.Errorf("Expected status code 204 on ping for train %s but got %d", k, res.StatusCode)
					wg.Done()
					return
				}

				// If ok, update LastContact
				item.LastContact = time.Now()
				cache.Map.Set(k, item)
				logrus.Infof("Train %s have been seen!", k)
				wg.Done()
				return
			}()


		}

		wg.Wait()

	}
}