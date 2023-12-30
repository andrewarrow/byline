div text-lg h-full
  div flex h-full
    div hidden md:block md:w-1/2 font-mono
      div flex flex-col h-full
        div h-full id=editor
        div pl-3 mb-auto id=command bg-purple-900
          commands 
          a href=https://github.com/andrewarrow/byline
            github
    div w-full md:w-1/2 flex-shrink-0 h-full
      div id=preview overflow-y-auto h-full
  div w-96 h-96 p-3 text-black bg-red-700 fixed top-0 left-2/4 id=debug hidden
