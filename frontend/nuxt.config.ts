import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({

  modules: [
    '@vee-validate/nuxt',
    '@nuxt/eslint',
    '@pinia/nuxt',
  ],
  devtools: { enabled: false },
  css: ['~/assets/css/main.css'],

  runtimeConfig: {
    infoUrl: process.env.NUXT_IPINFO,
    geoapify: process.env.NUXT_GEOAPIFY_URL,
    geoapify_key: process.env.NUXT_GEOAPIFY_KEY,
    ipinfo_token: process.env.NUXT_TOKEN_IP,
    public: {
      graphqlUrl: process.env.NUXT_GRAPHQL_URL,
    },
  },
  build: {
    transpile: ['vue-toastification'],
  },
  compatibilityDate: '2024-04-03',

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  eslint: {
    config: {
      stylistic: true,
    },
  },
  veeValidate: {
    autoImports: true,
    componentNames: {
      Form: 'VeeForm',
      Field: 'VeeField',
      FieldArray: 'VeeFieldArray',
      ErrorMessage: 'VeeErrorMessage',
    },
  },
})
