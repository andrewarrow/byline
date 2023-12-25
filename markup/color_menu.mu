{{ define "color_menu" }}
  {{ $selected := .Selected }}
  {{ $value := .Value }}
  {{ range $i, $item := .Colors }}
    {{ if eq $i $selected }}
      p bg-white text-black
        {{ $item }}-{{ $value }}
    {{ else }}
      p
        {{ $item }}-{{ $value }}
    {{ end }}
  {{ end }}
  {{ end }}
