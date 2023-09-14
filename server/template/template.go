package template

func HtmlTemplate(htmlContent string) (page string) {
	page = HtmlTop() + htmlContent + HtmlBottom()
	return
}

func HtmlTop() (topPage string) {
	topPage = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Dynamic Go</title>
    <link rel="stylesheet" type="text/css" href="/static/main.css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body>
    <header class="hd-cont">
        <div>
            <h1><a href="/" title="Home" class="no-underline">Dynamic Go Server</a></h1>
        </div>
        <nav class="nv-cont">
            <div class="flx-rw">
            <a href="/login" class="no-underline mr-rgt">
                <button class="btn w120">Log In</button>
            </a>
            <a href="/register" class="no-underline">
                <button class="btn">Register</button>
            </a>
        </div>
        </nav>
    </header>
    <main class="mn-sect">
        <article class="center-container">`
	return
}

func HtmlBottom() (bottomPage string) {
	bottomPage = `</article>
    </main>
    <footer class="ft-cont">
        <div>
        	<span class="cl-gr">Dynamic Go Server</span>
    	</div>
    </footer>
</body>
</html>`
	return
}
