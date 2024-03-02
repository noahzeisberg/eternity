export const content = [
  './pages/**/*.{vue,ts,tsx}',
  './components/**/*.{vue,ts,tsx}'
];
export const plugins = [
    require('@headlessui/tailwindcss'),
    require('@headlessui/tailwindcss')({prefix: 'ui'})
];