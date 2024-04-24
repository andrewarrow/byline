html attr-1 data-theme=dark
  head
    {{ $build := index . "build" }}
    link rel=stylesheet type=text/css href=/assets/css/tail.min.css?id!{{$build}}
    script src=https://cdn.tailwindcss.com
    script src=/assets/javascript/wasm_exec.js?id!{{$build}}
    script
      function $(id) { return document.getElementById(id); }
    title
      {{ index . "title" }}
    {{ index . "viewport" }}
  body
    div id=flash bg-red-600 text-white text-center fixed top-0 left-0 w-full
      {{ index . "flash" }}
    div h-full font-poppins text-base
      {{ index . "content" }}
    {{ index . "wasm" }}
