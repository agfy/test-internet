package speedtest

import (
	"errors"
	"fmt"
	"net/http"
)

// Speedtest is a speedtest client.
type Speedtest struct {
	doer *http.Client
}

// Option is a function that can be passed to New to modify the Client.
type Option func(*Speedtest)

// WithDoer sets the http.Client used to make requests.
func WithDoer(doer *http.Client) Option {
	return func(s *Speedtest) {
		s.doer = doer
	}
}

// New creates a new speedtest client.
func New(opts ...Option) *Speedtest {
	s := &Speedtest{
		doer: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

var defaultClient = New()

func StartTest(servers Servers) (float64, float64, error) {
	for _, s := range servers {
		err := s.PingTest()
		if err != nil {
			fmt.Printf("PingTest failed: " + err.Error())
			continue
		}
		err = testDownload(s)
		if err != nil {
			fmt.Printf("testDownload failed: " + err.Error())
			continue
		}
		err = testUpload(s)
		if err != nil {
			fmt.Printf("testUpload failed: " + err.Error())
			continue
		}
		return s.DLSpeed, s.ULSpeed, nil
	}
	return 0.0, 0.0, errors.New("all servers failed to evaluate speed")
}

func testDownload(server *Server) error {
	err := server.DownloadTest()
	if err != nil {
		return err
	}
	return err
}

func testUpload(server *Server) error {
	err := server.UploadTest()
	if err != nil {
		return err
	}
	return nil
}
