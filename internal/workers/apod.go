package workers

import (
	"NasaEnjoyer/apod"
	"context"
	"log"
	"time"
)

type APODWorker struct {
	service  *apod.Service
	interval time.Duration
}

func NewAPODWorker(service *apod.Service, interval time.Duration) *APODWorker {
	return &APODWorker{
		service:  service,
		interval: interval,
	}
}

func (w *APODWorker) Start(ctx context.Context, timeoutContext time.Duration) {
	for {
		log.Println("Starting APOD fetch job")
		jobCtx, cancel := context.WithTimeout(ctx, timeoutContext)
		err := w.service.FetchAndStoreAPOD(jobCtx)
		cancel()
		if err != nil {
			log.Printf("Error fetching and storing APOD: %v", err)
		} else {
			log.Println("Successfully fetched and stored APOD")
		}

		select {
		case <-time.After(w.interval):
			// Continue to the next iteration
		case <-ctx.Done():
			log.Println("Stopping APOD worker")
			return
		}
	}
}
