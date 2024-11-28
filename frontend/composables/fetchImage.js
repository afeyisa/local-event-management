import { GET_IMAGE } from '~/graphql/querie/getImage.graphql'
import { apolloClient } from '~/plugins/apollo'

export const fetchBase64Image = async (link) => {
  if (link) {
    try {
      const { data } = await apolloClient.query({
        query: GET_IMAGE,
        variables: { link },
      })
      return data.getImage.image
    }
    catch {
      return null
    }
  }
}
