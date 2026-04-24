package api

import (
	"context"
	"fmt"
	l "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/config"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/db"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/repository"
	"github.com/aldihorn/pact-cfp-demo/apps/api/internal/service"
	"github.com/pact-foundation/pact-go/v2/log"
	"github.com/pact-foundation/pact-go/v2/provider"
)

func TestToDoProvider(t *testing.T) {

	log.SetLogLevel("INFO")
	go startInstrumentedServer()

	verifier := provider.NewVerifier()

	err := verifier.VerifyProvider(t, provider.VerifyRequest{
		Provider:           "ToDoService-Backend",
		ProviderBaseURL:    fmt.Sprintf("http://127.0.0.1:%d", 8080),
		ProviderBranch:     os.Getenv("VERSION_BRANCH"),
		FailIfNoPactsFound: false,

		BrokerURL:                  os.Getenv("PACT_BROKER_BASE_URL"),
		BrokerToken:                os.Getenv("PACT_BROKER_TOKEN"),
		PublishVerificationResults: true,
		ProviderVersion:            os.Getenv("VERSION_COMMIT"),
	})

	if err != nil {
		t.Log(err)
	}

}

func startInstrumentedServer() {
	cfg := config.Load()

	database, err := db.ConnectFromConfig(cfg)
	if err != nil {
		l.Fatalf("failed to connect database: %v", err)
	}

	repo := repository.NewTodoRepository(database)
	svc := service.NewTodoService(repo)
	router := NewRouter(svc, cfg)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		l.Printf("api listening on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatalf("server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Printf("shutdown error: %v", err)
	}
}
