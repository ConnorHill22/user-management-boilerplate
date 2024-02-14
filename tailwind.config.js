/** @type {import('tailwindcss').Config} */
export const content = [
  './web/**/*.templ',
];
export const theme = {
  extend: {
    fontFamily: {
      mono: ['Courier Prime', 'monospace'],
    },
  },
};
export const plugins = [
];
export const corePlugins = { preFlight: true };
