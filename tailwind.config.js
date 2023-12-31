/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['views/*.html',],
  theme: {
    extend: {
      colors: {
        'lime': '#8FBC8F',
        'squeeze': '#5423e7',
      },
      fontFamily: {
        poppins: ['Poppins'],
        montserrat: ['Montserrat'],
        oxygen: ['Oxygen Mono'],
      },
    },
  },
  plugins: [],
}
