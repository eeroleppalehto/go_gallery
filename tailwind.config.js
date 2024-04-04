/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["views/**/*.templ"],
    darkMode: "selector",
    theme: {
        extend: {
            fontFamily: {
                oswald: ["Oswald", "sans-serif"],
                lato: ["Lato", "sans-serif"],
            },
        },
    },
    plugins: [],
};
