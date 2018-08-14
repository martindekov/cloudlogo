package function

// Handle a serverless request
func Handle(req []byte) string {
	//return fmt.Sprintf("Hello, Go. You said: %s", string(req))
	html := `
			<body style="background-color:#B0D0D3">
				<div style="display: block;margin-left: auto;margin-right: auto;width:50%;">
					<img src="https://www.openfaas.com/images/openfaas/logo.png" alt="Paris">
					<img src="https://irp-cdn.multiscreensite.com/0e5c5b66/dms3rep/multi/tablet/animated-cloud-right-to-left-350x211-350x211.png" style="z-index=3;">
				</div>
			</body>
			`
	return html
}
