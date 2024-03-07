package main

import (
	sitemap "github.com/Ishankhan21/go-site-map-generator/sitemap"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
    A link to another page
    <span> some span  </span>
  </a>
  <a href="/page-two">A link to a second page</a>
  <a href="/page-two">A link to a second page</a>

</body>
</html>
`

func main() {
	sitemap.SiteMap()
}
