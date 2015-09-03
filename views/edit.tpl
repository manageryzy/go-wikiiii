<form action="/edit/{{.Title}}" method="post">
	<h3>{{.Title}}</h3>
	<h5>内容</h5>
	<textarea name="content">{{.Src}}</textarea>
	<br/>
	编辑原因:<input name="reason"/>
	<br/>
	<b>允许脚本</b>
	<input name="EnableScript" type="checkbox"/>
	<br/>
	<input type="submit"/>
</form>