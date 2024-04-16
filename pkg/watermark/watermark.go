package watermark

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/velotiotech/watermark-service/internal"

	"github.com/go-kit/kit/log"
	"github.com/lithammer/shortuuid/v3"
)

type waterMarkService struct {}

func NewService() Service { return &waterMarkService{} }

func (w *waterMarkService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	doc := internal.Document{
		Content: "book",
		Title: "Harry Potter and Half Blood Prince",
		Author: "J.K Rowling", 
		Topic: "Fiction and Magic",
	}
	return []internal.Document{doc}, nil
}

func (w *waterMarkService) Status(_ context.Context, ticketID string) (internal.Status, error) {
	// this method returns internal status of watermark
	return internal.InProgress(), nil
}

func (w *waterMarkService) Watermark(_ context.Context, ticketID string, mark string) (int, error) {
	// this method returns internal status with OK if request has finished successful
	return http.StatusOK, nil
}

func (w *waterMarkService) AddDocument(_ context.Context, doc *internal.Document) (string, error) {
	// this method adds a new document and returns it ticketId 
	newTicketID := shortuuid.New()
	return newTicketID, nil
}

func (w *waterMarkService) ServiceStatus(_ context.Context) (int, error) {
	// this method is created to check the service status
	logger.Log("Checking the Service health...")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}