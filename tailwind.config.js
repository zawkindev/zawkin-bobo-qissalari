/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.html"],
  theme: {
    extend: {
      colors: {
        black: {
          50: "#252525",
          100: "#2e2e2e",
          300: "#1e1e1e",
          600: "#111111",
        },
      },
    },
  },
  plugins: [],
};
