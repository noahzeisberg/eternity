// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    modules: ["@nuxt/image", "@nuxt/ui"],
    app: {
        head: {
            title: "Eternity Web - Active Development Preview",
            meta: [{ name: "referrer", content: "no-referrer" }]
        },
    },
    ssr: false,
    nitro: {
      storage: {
          database: {
              driver: "fs",
              base: "database"
          }
      }
    },
    routeRules: {
        "/api/**": {
            cors: true
        }
    }
})