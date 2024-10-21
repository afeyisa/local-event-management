import { useRouter } from 'vue-router'

import { CHECK_AUTH_QUERY } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

export default defineNuxtRouteMiddleware(async () => {
  const router = useRouter()
  try {
    const { data } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
    // Check if the response is correct
    if (data && data.isAuthenticated) {
      return router.push('/')
    }
  }
  catch (error) {
    // If there is an error, redirect to login
  }
})
