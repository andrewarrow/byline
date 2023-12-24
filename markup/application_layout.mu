html attr-1
  head
    {{ $build := index . "build" }}
    link rel=stylesheet type=text/css href=/assets/css/tail.min.css?id!{{$build}}
    {{ if index . "USE_LIVE_TEMPLATES" }}
      script src=https://cdn.tailwindcss.com
    {{ end }}
    script src=/assets/javascript/wasm_exec.js?id!{{$build}}
    script
      function $(id) { return document.getElementById(id); }
    title
      {{ index . "title" }}
    {{ index . "viewport" }}
  body
    div id=flash bg-red-600 text-white text-center fixed top-0 left-0 w-full
      {{ index . "flash" }}
    div bg-squeeze text-gray-300 overflow-y-auto pl-3 pr-3 min-h-screen font-poppins text-base
      {{ index . "content" }}
    {{ index . "wasm" }}
