import { CHECK_AUTH_QUERY } from '~/graphql/queries'
import { apolloClient } from '~/plugins/apollo'

// let isCheckingAuth = true
export default defineNuxtRouteMiddleware(async (req) => {
  // const router = useRouter()
  console.log(req.fullPath)
  // try {
  const { data } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
  // Check if the response is correct
  if (data && data.isAuthenticated) {
    return navigateTo('/dasboard')
  }
  // return router.push('/loadingfallback')
  // }
  // catch (error) {
  //   return navigateTo('/loadingfallback')
  // }
})
