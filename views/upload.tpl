<!DOCTYPE html>

<html>
{{template "head.tpl" .}}
<body>
{{template "header.tpl" .}}
<form action="/upload" method="post" enctype="multipart/form-data">
    name:<input name="name"/><br/>
    file:<input type="file" name="file"/><br/>
    <input type="submit"/>
</form>
{{template "feet.tpl" .}}
</body>
</html>