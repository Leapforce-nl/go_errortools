package errortools

import (
	"time"

	"github.com/getsentry/sentry-go"
)

const SentryDSNTest string = "https://da7ff9970b3f4fb6b7f84ffeee423f87@o326694.ingest.sentry.io/5510091"

// InitSentry initializes logging to sentry
//
func InitSentry(dsn string, isLive bool) {
	if !isLive {
		// log to sentry test project
		dsn = SentryDSNTest
	}
	// We need to use the sync transport (which is not the default),
	// otherwise if you use "log.Fatal()" the program will exit before the
	// error is sent to Sentry.
	sentrySyncTransport := sentry.NewHTTPSyncTransport()
	sentrySyncTransport.Timeout = time.Second * 3
	err := sentry.Init(sentry.ClientOptions{
		AttachStacktrace: true,
		Dsn:              dsn,
		Transport:        sentrySyncTransport,
	})
	if err != nil {
		panic(err)
	}
}
