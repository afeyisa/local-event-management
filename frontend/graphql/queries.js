// hel

import { gql } from 'graphql-tag'

export const CHECK_AUTH_QUERY = gql`
  query {
    isAuthenticated
  }
`
