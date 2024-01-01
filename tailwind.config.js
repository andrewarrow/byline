/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['views/*.html','extra.html'],
  theme: {
    extend: {
      colors: {
        'lime': '#8fbc8f',
        "grape-100": "#b3aefb",
        "grape-200": "#998ff8",
        "grape-300": "#7e73f6",
        "grape-400": "#4f46e5",
        "grape-500": "#4f46e5",
        "grape-600": "#3e378c",
        "grape-700": "#302c6e",
        "grape-800": "#241e57",
        "grape-900": "#18123e",
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
