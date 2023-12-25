{{ define "menu" }}
  {{ $items := index . "items" }}
  {{ $selected := index . "selected" }}
  {{ range $i, $item := $items }}
    {{ if eq $i $selected }}
      p bg-white text-black
    {{ else }}
      p
    {{ end }}
      {{ $item }}
  {{ end }}
  {{ end }}
