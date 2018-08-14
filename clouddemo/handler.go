package function

// Handle a serverless request
func Handle(req []byte) string {
	//return fmt.Sprintf("Hello, Go. You said: %s", string(req))
	html := `<div>
				<img src="https://www.openfaas.com/images/openfaas/logo.png" alt="Paris" style="display: block;margin-left: auto;margin-right: auto;width:50%;">
			 </div>
				`
	return html
}
