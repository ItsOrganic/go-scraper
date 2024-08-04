<h1>An amazon web scraper that scrapes amazon products</h1>
1. Clone the repository: 
<pre><code> git clone https://github.com/ItsOrganic/go-scraper.git
</code></pre>
2. Go in the go-scraper directory:
<pre><code> cd go-scraper
</code></pre>
3. Install the dependencies:
<pre><code> go mod tidy
</code></pre>
4. Build the project:
<pre><code> go build .
</code></pre>
5. Run the project using the flag "-product=name-of-the-product"
<pre><code> ./go-scraper -product=bags
</code></pre>
6. Open localhost:3000 on your browser to see the json result.
