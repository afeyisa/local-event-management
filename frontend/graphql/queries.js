// hel

import { gql } from 'graphql-tag'

export const CHECK_AUTH_QUERY = gql`
  query {
    isAuthenticated
  }
`

export const GET_ORGANIZATIONS = gql`
  query {
    data_organizations {
      organization_name
      bio
      created_at
      description
      organization_id
      profile_photo_url
    }
  }
`
