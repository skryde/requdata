package main

var htmlString = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">

  <title>Request Data</title>
  <meta name="description" content="Proof of concept to know how much data you can get from a redirect from social media websites.">
  <meta name="author" content="SitePoint">
</head>

<body>
  <h1>Method</h1>
  <ul>
    <li>{{ .Method }}</li>
  </ul>

  <h1>Headers</h1>
  <ul>
    {{ range $key, $value := .Headers }}
    <li>{{ $key }}: {{ $value }}</li>
    {{ end }}
  </ul>

  <h1>Body</h1>
  <p>{{ .Body }}</p>
</body>
</html>`
