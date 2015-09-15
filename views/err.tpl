<!DOCTYPE html>

<html>
{{template "head.tpl" .}}
<body>
    {{template "header.tpl" .}}
    <h1>失败！<a href="javascript:history.back()">返回</a></h1>
    <div>错误原因：{{.ERROR}}</div>
    {{template "feet.tpl" .}}
</body>
</html>