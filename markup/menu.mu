{{ define "menu" }}
  {{ $selected := .Selected }}
  {{ range $i, $item := .Items }}
    {{ if eq $i $selected }}
      p bg-white text-black
        {{ $item }}
    {{ else }}
      p
        {{ $item }}
    {{ end }}
  {{ end }}
  {{ end }}
