/** @type {import('tailwindcss').Config} */
export const content = [
  './web/**/*.templ',
  './internal/routes/view/**/*.templ'
];
export const theme = {
  colors: {
    transparent: 'transparent',
    current: 'currentColor',
    'white': '#ffffff',
    'primary': '#253F57',
    'secondary': '#42A4FF',
    'tertiary': '#6589AA',
    't-black': '#242C33',
    'light-bg': '#F7F9FB',
    'light-gray': '#CCCCCC',
    'red': '#E00707',
    'green': '#10A333',
  },
  extend: {
    fontFamily: {
      mono: ['Courier Prime', 'monospace'],
    },
  },
};
export const plugins = [
];
export const corePlugins = { preFlight: true };
