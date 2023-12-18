{{ define "__subject" }}[{{ .Status | toUpper }}{{ if eq .Status "firing" }}:{{ .Alerts.Firing | len }}{{ if gt (.Alerts.Resolved | len) 0 }}, RESOLVED:{{ .Alerts.Resolved | len }}{{ end }}{{ end }}] {{ .GroupLabels.SortedPairs.Values | join " " }} {{ if gt (len .CommonLabels) (len .GroupLabels) }}{{ end }}{{ if eq .Status "firing" }} [{{(index .Alerts.Firing 0).Labels.severity}}]{{ end }}{{ end }}

{{ define "__fire_text_alert_list" }}{{ range . }}
<b style="color: red;">Firing [{{.Labels.severity}}]</b>
{{.Annotations.summary | safeHtml}}
{{.Annotations.description | safeHtml}}
{{ if gt (len .Annotations.dashboard) 0 }}<a href="{{ .Annotations.dashboard }}">Панель мониторинга</a>{{ end }}
=======================================

{{ end }}{{ end }}

{{ define "__resolve_text_alert_list" }}{{ range . }}
<b style="color: green;">Resolved</b>
{{.Annotations.summary | safeHtml}}
{{ if gt (len .Annotations.dashboard) 0 }}<a href="{{ .Annotations.dashboard }}">Панель мониторинга</a>{{ end }}
=======================================

{{ end }}{{ end }}

{{ define "simple.title" }}{{ template "__subject" . }}{{ end }}

{{ define "simple.message" }}{{ if gt (len .Alerts.Firing) 0 }}{{ template "__fire_text_alert_list" .Alerts.Firing }}
{{ end }}{{ if gt (len .Alerts.Resolved) 0 }}{{ template "__resolve_text_alert_list" .Alerts.Resolved }}
{{ end }}{{ end }}