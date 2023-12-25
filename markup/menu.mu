{{ define "menu" }}
  {{ range $i, $item := . }}
    p
      {{ $item }}
  {{ end }}
  {{ end }}
