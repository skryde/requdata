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
  <h1>Description</h1>
  <p>Proof of concept to know how much data you can get from a redirect from social media applications.</p>

  <hr/>

  <h3>Method</h3>
  <ul>
    <li>{{ .Method }}</li>
  </ul>

  <h3>Headers</h3>
  <ul>
    {{ range $key, $value := .Headers }}
    <li>{{ $key }}: {{ $value }}</li>
    {{ end }}
  </ul>

  <h3>Query params</h3>
  <ul>
    {{ range $key, $value := .QueryParams }}
    <li>{{ $key }}: {{ $value }}</li>
    {{ end }}
  </ul>

  <h3>Body</h3>
  <p>{{ .Body }}</p>
</body>
</html>`
