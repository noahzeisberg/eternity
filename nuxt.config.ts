// https://nuxt.com/docs/api/configuration/nuxt-config
// @ts-ignore
export default defineNuxtConfig({
    devtools: {
        enabled: false
    },
    modules: ['@nuxtjs/tailwindcss', "@nuxt/image"],
    app: {
        head: {
            title: "Unity Chores - Preview",
            link: [{ rel: 'icon', type: 'image/ico', href: 'favicon.ico' }]
        },
    },
    ssr: false,
    nitro: {
      storage: {
          dataStorage: {
              driver: "fs",
              base: ".cache"
          }
      }
    }
})