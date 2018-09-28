package function

// Handle a serverless request
func Handle(req []byte) string {
	color := true
	if color {
		html := `
		<body style="display:table;position:absolute;height:100%;background-color:#B0D0D3;width:100%;">
			<div style="display: table-cell;vertical-align: middle;">
				<div style="margin-left: auto;margin-right: auto;">
					<img src="https://www.openfaas.com/images/openfaas/openfaas.png" style="display:block;margin:auto">
					<img src="https://www.openfaas.com/images/openfaas/logo.png" style="display:block;margin:auto">
					<img src="https://irp-cdn.multiscreensite.com/0e5c5b66/dms3rep/multi/tablet/animated-cloud-right-to-left-350x211-350x211.png" style="z-index:5;height:250px;position:relative;top:-70px;left:-10px;display:block;margin:auto">
				</div>
			</div>
		</body>
		`
		return html
	}
	html := `
		<body style="display:table;position:absolute;height:100%;background-color:#276ddd;width:100%;">
			<div style="display: table-cell;vertical-align: middle;">
				<div style="margin-left: auto;margin-right: auto;"> 
					<img src="https://www.openfaas.com/images/openfaas/openfaas.png" style="display:block;margin:auto">
					<img src="https://www.openfaas.com/images/openfaas/logo.png" style="display:block;margin:auto">
					<img src="https://irp-cdn.multiscreensite.com/0e5c5b66/dms3rep/multi/tablet/animated-cloud-right-to-left-350x211-350x211.png" style="z-index:5;height:250px;position:relative;top:-70px;left:-10px;display:block;margin:auto">
				</div>
			</div>
		</body>
		`
	return html
}
