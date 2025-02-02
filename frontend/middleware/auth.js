import { useRouter } from 'vue-router'
import { CHECK_AUTH_QUERY } from '~/graphql/querie/checkAuthentication.graphql'
import { apolloClient } from '~/plugins/apollo'

export default defineNuxtRouteMiddleware(async (req) => {
  const router = useRouter()
  try {
    const { data } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
    if (!data || !data.isAuthenticated) {
      localStorage.setItem('redirect', req.fullPath)
      return navigateTo('/login')
    }
  }
  catch (error) {
    return router.push('/loadingfallback')
  }
})
