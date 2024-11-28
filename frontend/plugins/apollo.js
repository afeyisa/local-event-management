import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
} from '@apollo/client/core'
import { provideApolloClient } from '@vue/apollo-composable'
import { graphqlUrl } from '~/composables/graphqlUrl'

const httpLink = createHttpLink({
  uri: graphqlUrl,
  credentials: 'include',
})

export const apolloClient = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache(),
  defaultOptions: {
    query: {
      fetchPolicy: 'no-cache',
    },
    mutate: {
      fetchPolicy: 'no-cache',
    },
  },
})
export default defineNuxtPlugin(() => {
  provideApolloClient(apolloClient)
})
