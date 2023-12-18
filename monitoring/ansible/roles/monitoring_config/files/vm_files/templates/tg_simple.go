{{ define "__tg_fire_text_alert_list" }}{{ range . }}
🆘 🆘 🆘
<b>Firing</b>
{{.Annotations.summary | safeHtml}}
{{.Annotations.description | safeHtml}}
{{ if gt (len .Annotations.dashboard) 0 }}<a href="{{ .Annotations.dashboard }}">Панель мониторинга</a>{{ end }}
=======================================

{{ end }}{{ end }}

{{ define "__tg_resolve_text_alert_list" }}{{ range . }}
✅ ✅ ✅
<b>Resolved</b>
{{.Annotations.summary | safeHtml}}
{{ if gt (len .Annotations.dashboard) 0 }}<a href="{{ .Annotations.dashboard }}">Панель мониторинга</a>{{ end }}
=======================================

{{ end }}{{ end }}

{{ define "tg.simple.message" }}{{ if gt (len .Alerts.Firing) 0 }}{{ template "__tg_fire_text_alert_list" .Alerts.Firing }}
{{ end }}{{ if gt (len .Alerts.Resolved) 0 }}{{ template "__tg_resolve_text_alert_list" .Alerts.Resolved }}
{{ end }}{{ end }}