module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
     "./public/index.html",
  ],
  theme: {
    extend: {
      colors: {
        'bookstore-primary': '#2d5a4d',
        'bookstore-secondary': '#5fe9bc',
      },
      fontFamily: {
        'sans': ['Prompt', 'sans-serif'],
      }
    },
  },
  plugins: [],
}