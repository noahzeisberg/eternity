// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
      '@nuxtjs/tailwindcss'
  ],
    $production: {
      nitro: {
          storage: {
              dataStorage: {
                  driver: "redis"
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
