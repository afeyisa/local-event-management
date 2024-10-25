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
export const GET_ORGANIZATIONS_FOR_EVENT = gql`
  query{
    data_organizations{
      organization_id
      organization_name
    }
  }`
export const GET_CATAGORIES = gql`
  query{
    data_categories {
      category_id
      category_name
  }
}
`

export const GET_TAGS = gql`
  query{
    data_tags {
    tag_id
    tag_word
  }
  }
`

export const GET_EVENT_SAMPLE = gql`
  query{
    data_events {
      event_id
      title
      description
      event_date
      venue
      thumbnail_image_url
  }
}
`
export const GET_EVENT_DETAILS = gql`
  query GetEventDetails($id: uuid!) {
    data_events_by_pk(event_id: $id) {
      title
      ticket_price
      total_available_tickets
      event_date
      category {
        category_name
      }
      description
      venue
      is_online
      thumbnail_image_url
      images {
        image_url
      }
      address {
        street_name
        city_name
        region_name
      }
    }
  }
`