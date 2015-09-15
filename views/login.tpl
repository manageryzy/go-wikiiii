<!DOCTYPE html>

<html>
{{template "head.tpl" .}}
<body>
    {{template "header.tpl" .}}
	<form action="/login">
		username: <input name="username"/>
		password: <input name="password" type="password"/>
		<input type="submit"/>
	</form>
	{{template "feet.tpl" .}}
</body>
</html>