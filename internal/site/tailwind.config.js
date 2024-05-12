/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["internal/site/render/view/**/**/*.templ"],
    theme: {
        extend: {},
    },
    corePlugins: {
        preflight: true,
    }
}