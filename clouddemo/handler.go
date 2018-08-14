package function

// Handle a serverless request
func Handle(req []byte) string {
	//return fmt.Sprintf("Hello, Go. You said: %s", string(req))
	html := `<img src="https://www.openfaas.com/images/openfaas/logo.png" alt="Paris" class="center">`
	return html
}
