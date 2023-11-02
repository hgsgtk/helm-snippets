{{/* Generate basic labels */}}
{{- define "mychart.labels" }}
  labels:
    generator: helm
    data: {{ now | htmlDate }}
    chart: {{ .Chart.Name }}
    version: {{ .Chart.Version }}
{{- end }}