export const content = [
    './pages/**/*.{vue,ts,tsx}',
    './components/**/*.{vue,ts,tsx}'
];
export const theme = {
    extend: {
        colors: {
            zinc: {
                925: "#111113"
            }
        }
    }
};
export const plugins = [
    require('@headlessui/tailwindcss'),
    require('@headlessui/tailwindcss')({prefix: 'ui'})
];