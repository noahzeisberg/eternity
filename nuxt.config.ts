// https://nuxt.com/docs/api/configuration/nuxt-config
// @ts-ignore
export default defineNuxtConfig({
    devtools: {
        enabled: true
    },
    modules: [
      '@nuxtjs/tailwindcss'
    ],
    app: {
        baseURL: "/unity/"
    },
    ssr: false,
    $production: {
      nitro: {
          storage: {
              dataStorage: {
                  driver: "memory"
              }
          }
      }
    },
    nitro: {
      storage: {
          dataStorage: {
              driver: "fs",
              base: ".cache"
          }
      }
    }
})
