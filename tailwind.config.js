/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.html"],
  theme: {
    extend: {
      colors: {
        "primary-red": "#FF0307",
        "primary-text": "#000908",
        "primary-bg": "#FBFFFE",
        "primary-blue": "#91A6FF",
      },
    },
  },
  plugins: [],
};
