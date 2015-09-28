<!DOCTYPE html>

<html>
{{template "head.tpl" .}}
<body>
    {{template "header.tpl" .}}
    <h1>history of {{.Title}}</h1>
    <table border="1">
        <thead>
            <tr>
                <td>file name</td>
                <td>edit time</td>
                <td>editor</td>
                <td>action</td>
            </tr>
        </thead>
        <tbody>
            {{range .History}}
            <tr>
                <td>{{.FileName}}</td>
                <td>{{.Update}}</td>
                <td><a href="/user/{{.Uid}}">{{.Name}}</a></td>
                <td>
                    <a href="/file/history/get/{{.Fhid}}">View</a> |
                    <a href="/file/history/use/{{.Fhid}}">Use as current file</a>
                </td>
            </tr>
            {{end}}
        </tbody>
    </table>
    {{template "feet.tpl" .}}
</body>
</html>
