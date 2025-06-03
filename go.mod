module TryRTMP

go 1.24.3

// go mod edit -replace github.com/diamondburned/gotk4=github.com/rswilli/gotk4@generator_refactor
// did I have a go install here?
// go get github.com/go-gst/go-gst@generated_bindings
// go mod tidy

replace github.com/diamondburned/gotk4 => github.com/rswilli/gotk4 v0.0.0-20250519074043-31f51c6f5601

require github.com/go-gst/go-gst v1.4.1-0.20250519092828-e2ae3be3b208

require (
	github.com/diamondburned/gotk4 v0.3.1 // indirect
	golang.org/x/sync v0.8.0 // indirect
)
