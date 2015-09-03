<!DOCTYPE html>

<html>
{{template "head.tpl" .}}
<body>
{{template "header.tpl" .}}
<div class="main">
    {{.Page}}
    <div class="wiki-pages">
        {{range .Pages}}
        <li><a href="/page/{{.Title}}">{{.Title}}</a></li>
        {{end}}
    </div>
</div>

<div class="wiki-categories">
    {{range .Category}}
    <li><a href="/page/{{.Category}}/category">{{.Category}}</a></li>
    {{end}}
</div>

{{template "feet.tpl" .}}
</body>
</html>
