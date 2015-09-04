<!DOCTYPE html>

<html>
{{template "head.tpl" .}}
<body>
    {{template "header.tpl" .}}
    <h1>history of {{.Title}}</h1>
    <table border="1">
        <thead>
            <tr>
                <td>id</td>
                <td>edit time</td>
                <td>reason</td>
                <td>editor</td>
                <td>action</td>
            </tr>
        </thead>
        <tbody>
            {{range .History}}
            <tr>
                <td>{{.Hid}}</td>
                <td>{{.Update}}</td>
                <td>{{.Reason}}</td>
                <td><a href="/user/{{.Uid}}">{{.Name}}</a></td>
                <td>
                    <a href="/page/{{.Title}}/history/{{.Hid}}">View</a> |
                    <a href="/page/{{.Title}}/history/{{.Hid}}?src">Src</a>
                </td>
            </tr>
            {{end}}
        </tbody>
    </table>
    {{template "feet.tpl" .}}
</body>
</html>
