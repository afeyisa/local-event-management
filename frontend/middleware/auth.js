import { useRouter } from 'vue-router'
import { CHECK_AUTH_QUERY } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

// let isCheckingAuth = true
export default defineNuxtRouteMiddleware(async (req) => {
  const router = useRouter()
  try {
    const { data } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
    // Check if the response is correct
    if (!data || !data.isAuthenticated) {
      localStorage.setItem('redirect', req.fullPath)
      // Redirect to login if not authenticated
      return router.push('/login')
    }
  }
  catch (error) {
    return router.push('/loadingfallback')
  }
})

