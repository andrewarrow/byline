div p-3 text-4xl pb-64
  div py-9 text-center
    welcome
  div flex justify-center space-x-9
    div
      button id=login border rounded bg-blue-600 text-white py-2 px-2
        Login
    div
      button id=register whitespace-nowrap border rounded bg-blue-600 text-white py-2 px-2
        Sign Up
  form py-9 id=login-form method=POST hidden
    div flex justify-center
      div space-y-9
        div 
          input w-64 type=text id=login-username autofocus=true required=true placeholder=username
        div
          input w-64 type=password id=login-password required=true placeholder=password
        div
          input w-64 type=submit value=Go border rounded bg-blue-600 text-white py-2 px-2
        div text-2xl text-center
          a href=/ underline
            Forgot Password
  form py-9 id=register-form method=POST hidden
    div flex justify-center
      div space-y-9
        div 
          input w-64 type=text id=register-username autofocus=true required=true placeholder=username
        div
          input w-64 type=password id=register-password required=true placeholder=password
        div
          input w-64 type=submit value=Go border rounded bg-blue-600 text-white py-2 px-2
