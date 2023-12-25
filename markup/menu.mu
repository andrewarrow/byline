{{ define "menu" }}
  {{ $items := index . "items" }}
  {{ $selected := index . "selected" }}
  {{ range $i, $item := $items }}
    {{ if eq $i $selected }}
      p bg-white text-black
        {{ $item }}
    {{ else }}
      p
        {{ $item }}
    {{ end }}
  {{ end }}
  {{ end }}
