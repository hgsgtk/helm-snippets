{{/* Generate basic labels */}}
{{- define "mychart.labels" }}
  labels:
    generator: helm
    data: {{ now | htmlDate }}
{{- end }}