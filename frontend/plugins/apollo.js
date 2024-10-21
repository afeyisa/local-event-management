import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
} from '@apollo/client/core'
import { provideApolloClient } from '@vue/apollo-composable'

// Create an HTTP link to your GraphQL endpoint
const httpLink = createHttpLink({
  uri: 'http://localhost:8080/v1/graphql', // Your GraphQL endpoint
  credentials: 'include', // Include cookies in requests
})

// Create an Apollo Client instance
export const apolloClient = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache(),
})
// export default apolloClient
export default defineNuxtPlugin(() => {
  provideApolloClient(apolloClient)
})
