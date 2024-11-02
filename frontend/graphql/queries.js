// hel

import { gql } from 'graphql-tag'

export const CHECK_AUTH_QUERY = gql`
  query {
    isAuthenticated
  }
`
export const GET_MY_ID = gql`
  query{data_users {
    user_id
  }
  }
`

export const GET_ORGANIZATIONS = gql`
  query orgs($where:data_organizations_bool_exp!){
    data_organizations(where:$where) {
      organization_name
      bio
      created_at
      description
      organization_id
      profile_photo_url
      followers
      follows {
        following_user_id
        followed_organization_id
      }
    }
  }
`

export const PUBLIC_GET_ORGANIZATIONS = gql`
  query orgs($where:data_organizations_bool_exp!){
    data_organizations(where:$where) {
      organization_name
      bio
      created_at
      description
      organization_id
      profile_photo_url
      followers
    }
  }
`
// when creating event or for event form
export const GET_ORGANIZATIONS_MY_ORG = gql`
  query getMyorgs($where:data_organizations_bool_exp!){
    data_organizations(where:$where)
    {
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

export const GET_MY_EVENTS = gql`
  query myEvents($where: data_events_bool_exp!){
    data_events(where: $where, order_by: {created_at: desc}) {
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
      event_id
      title
      ticket_price
      total_available_tickets
      event_date
      by_organization_id
      category_id
      category {
        category_name
      }
      description
      venue
      thumbnail_image_url
      images {
        image_url
      }
      address {
        street_name
        city_name
        region_name
      }

      event_tags {
      tag_word_id
      }
      bookmarks {
        book_marked_event_id
        book_marker_user_id
      }
      location{
        latitude
        longitude
      }
    }
  }
`

export const PUBLIC_GET_EVENT_DETAILS = gql`
  query GetEventDetails($id: uuid!) {
    data_events_by_pk(event_id: $id) {
      event_id
      title
      ticket_price
      total_available_tickets
      event_date
      by_organization_id
      category_id
      category {
        category_name
      }
      description
      venue
      thumbnail_image_url
      images {
        image_url
      }
      address {
        street_name
        city_name
        region_name
      }

      event_tags {
      tag_word_id
      }
      location{
        latitude
        longitude
      }
    }
  }
`

export const BROWSEEVENTS = gql`
 query SearchEventsByTitle($where: data_events_bool_exp!) {
    data_events(where: $where) {
      event_id
      title
      description
      ticket_price
      event_date
      thumbnail_image_url
      total_available_tickets
      address {
        street_name
        city_name
        region_name
      }
      location{
        latitude
        longitude
      }
      category{
      category_name
      }

    bookmarks {
      book_marker_user_id
      book_marked_event_id
    }

    organization {
      created_at
      followers
      organization_id
      organization_name
      follows {
        following_user_id
        followed_organization_id
      }
    }

    }
  }
`

export const PUBLIC_BROWSEEVENTS = gql`
 query SearchEventsByTitle($where: data_events_bool_exp!) {
    data_events(where: $where) {
      event_id
      title
      description
      ticket_price
      event_date
      thumbnail_image_url
      total_available_tickets
      address {
        street_name
        city_name
        region_name
      }
      location{
        latitude
        longitude
      }
      category{
      category_name
      }
    organization {
      created_at
      followers
      organization_id
      organization_name
    }

    }
  }

`
export const ORG_TOTAL_EVENTS = gql`
query GetEventCountByOrganization($organizationId: uuid!) {
  data_events_aggregate(where: { by_organization_id: { _eq: $organizationId } }) {
    aggregate {
      count
    }
  }
}

`
