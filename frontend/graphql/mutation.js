import { gql } from 'graphql-tag'

export const LOGIN_MUTATION = gql`
  mutation Login($email: String!, $password: String!) {
    login(email: $email, password: $password) {
    success
  }
  }
`
export const SIGNUP_MUTATION = gql`
mutation Signup($email: String!, $password: String!) {
  signup(email: $email, password: $password) {
    success
  }
}`

export const CREATE_ORGANIZATION = gql`
  mutation CreateOrganization(
    $organization_name: String!, 
    $profile_photo_url: String, 
    $bio: String, 
    $description: String,
  ) {
    insert_data_organizations_one(
      object: {
        organization_name: $organization_name, 
        profile_photo_url: $profile_photo_url, 
        bio: $bio, 
        description: $description, 
      }
    ) {
      organization_id
    }
  }
`

export const CREATE_EVENT = gql`
mutation insertEvent(
  $by_organization_id: uuid!,
  $title: String!,
  $ticket_price: numeric!,
  $total_available_tickets: Int!,
  $event_date: date!,
  $category_id: uuid!,
  $description: String!,
  $venue: String!,
  $is_online: Boolean,
  $thumbnail_image_url: String!,
  $images: [data_images_insert_input!]!,
  $location: data_locations_insert_input!,
  $address: data_address_insert_input!,
  # $tags: [data_event_tags_insert_input!]!
) {
  # Insert into events table
  insert_data_events_one(
    object: {
      by_organization_id: $by_organization_id,
      title: $title,
      ticket_price: $ticket_price,
      total_available_tickets: $total_available_tickets,
      event_date: $event_date,
      category_id: $category_id,
      description: $description,
      venue: $venue,
      is_online: $is_online,
      thumbnail_image_url: $thumbnail_image_url,
      
  # Insert into images table
  images:{data: $images},
  # event_tags{data:}

  # Insert into locations table
  location:{data: $location},

  # Insert into address table
  address:{data: $address},

  # Insert into event_tags table
  # event{data: $tags},
    }
  ) {
    event_id
    title
  }
}`

// insert_data_organizations_one(object: {bio: "t", description: "d", organization_name: "name", profile_photo_url: "", organizes: {data: {organization_id: "i1", organizer_id: "i2"}}}) {
//     bio
//   }
// const UPDATE_ORGANIZATION = gql`
//   mutation UpdateOrganization(
//     $organization_name: String!,
//     $profile_photo_url: String,
//     $bio: String,
//     $description: String
//   ) {
//     update_data_organizations_by_pk(
//       pk_columns: { organization_id: $organization_id },
//       _set: {
//         organization_name: $organization_name,
//         profile_photo_url: $profile_photo_url,
//         bio: $bio,
//         description: $description
//       }
//     ) {
//       organization_id
//       organization_name
//       profile_photo_url
//       bio
//       description
//     }
//   }
// `
