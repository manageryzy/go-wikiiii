<form action="/edit/{{.Title}}" method="post">
	<h3>{{.Title}}</h3>
	<h5>内容</h5>
	<textarea name="content">{{.Src}}</textarea>
	<input type="submit"/>
</form>