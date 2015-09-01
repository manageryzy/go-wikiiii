{{template "head.tpl" .}}

<div class="main">
    {{.Page}}
</div>

<div class="wiki-categories">
    {{range .Category}}
        <li><a href="/page/{{.Category}}/category">{{.Category}}</a></li>
    {{end}}
</div>

{{template "feet.tpl"}}
