{{ define "color_menu" }}
  {{ $selected := .Selected }}
  {{ range $i, $item := .Colors }}
    {{ if eq $i $selected }}
      p bg-white text-black
        {{ $item }}
    {{ else }}
      p
        {{ $item }}
    {{ end }}
  {{ end }}
  {{ end }}
