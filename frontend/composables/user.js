import { apolloClient } from '~/plugins/apollo'
import { CHECK_AUTH_QUERY } from '~/graphql/querie/checkAuthentication.graphql'
import { GET_MY_ID } from '~/graphql/querie/getUserId.graphql'

export const isLoggedIn = async () => {
  try {
    const { data } = await apolloClient.query({ query: CHECK_AUTH_QUERY })
    if (data && data.isAuthenticated) {
      return true
    }
    return false
  }
  catch {
    return false
  }
}

export const fetchUseId = async () => {
  const myId = ref(null)

  if (await isLoggedIn()) {
    const { data: userData, loading, error } = await apolloClient.query({ query: GET_MY_ID })
    if (!loading && !error) {
      myId.value = userData.data_users[0].user_id
    }
  }

  return { myId }
}
