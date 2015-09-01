<form action="/edit/{{.Title}}" method="post">
	<h3>{{.Title}}</h3>
	<h5>内容</h5>
	<textarea name="content">{{.Src}}</textarea>
	<h5>允许脚本</h5>
	<input name="EnableScript" type="checkbox"/>
	<br/>
	<input type="submit"/>
</form>