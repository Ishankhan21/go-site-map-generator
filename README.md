# Golang Site Map Generator

A simple and efficient site map generator written in Go. This tool generates sitemaps in XML format for websites, helping in SEO and website navigation.

## Features

- Generates sitemaps in XML format.
- Supports Breath-First Search (BFS) algorithm for crawling websites.
- Filters out links with different domains to ensure sitemap relevance.
- Efficiently parses and constructs URLs from the crawled pages.

## Installation

1. Ensure you have Go installed on your system. If not, you can download it from [here](https://golang.org/dl/).
2. Clone the repository:

### Example Output

Here is an example of what the output might look like:
```
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://gophercises.com/demos/cyoa/mark-bates</loc>
  </url>
  <url>
    <loc>https://gophercises.com</loc>
  </url>
  <url>
    <loc>https://gophercises.com/demos/cyoa/home</loc>
  </url>
  <url>
    <loc>https://gophercises.com/demos/cyoa/debate</loc>
  </url>
  <url>
    <loc>https://gophercises.com/demos/cyoa/denver</loc>
  </url>
  <url>
    <loc>https://gophercises.com/demos/cyoa/sean-kelly</loc>
  </url>
  <url>
    <loc>https://gophercises.com/demos/cyoa/</loc>
  </url>
  <url>
    <loc>https://gophercises.com/</loc>
  </url>
  <url>
    <loc>https://gophercises.com/demos/cyoa/new-york</loc>
  </url>
</urlset>
```