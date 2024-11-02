import { gql } from 'graphql-tag'

export const LOGIN_MUTATION = gql`
  mutation Login($email: String!, $password: String!) {
    login(email: $email, password: $password) {
    success
  }
  }
`
export const LOGOUT_MUTATION = gql`
  mutation {
    logout {
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

export const UPLOAD_THUMBNAIL_IMAGE = gql`
    mutation uploadImage($image: String!) {
      uploadImage(image: $image) {
        url
      }
    }
  `

export const CREATE_EVENT = gql`
mutation insertEvent(
  $by_organization_id: uuid!,
  $title: String!,
  $ticket_price: numeric,
  $total_available_tickets: Int,
  $event_date: date,
  $category_id: uuid,
  $description: String,
  $venue: String,
  # $is_online: Boolean,
  $thumbnail_image_url: String,
  $images: [data_images_insert_input!]!,  
  $location: data_locations_insert_input!,
  $address: data_address_insert_input!,
  $tags: [data_event_tags_insert_input!]! 
) {
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
      # # is_online: $is_online,
      thumbnail_image_url: $thumbnail_image_url,

      images: {data: $images}, 
      # Insert into locations table
      location: {data: $location},
      # Insert into address table
      address: {data: $address},
      # Insert into event_tags table
      event_tags: {data: $tags},
    }
  ) {
    event_id
    title
  }
}`

export const DELETE_IMAGE = gql`
mutation deleteImage(
  $event_id:uuid!
  $image_url:String
){
  delete_data_images(where: {_and: {event_id: {_eq: $event_id}, image_url: {_eq: $image_url}}}) {
    affected_rows
  }
}
`
export const UPDATE_ORGANIZATION = gql`
  mutation UpdateOrganization(
    $organization_name: String!,
    $profile_photo_url: String,
    $bio: String,
    $description: String
    $organization_id:uuid!
  ) {
    update_data_organizations_by_pk(
      pk_columns: { organization_id: $organization_id },
      _set: {
        organization_name: $organization_name,
        profile_photo_url: $profile_photo_url,
        bio: $bio,
        description: $description
      }
    ) {
      organization_id
    }
  }
`
export const UPDATE_EVENT = gql`
  mutation updateEvent(
    $id: uuid!,
    $category_id: uuid,
    $by_organization_id: uuid!,
    $description: String,
    $event_date: date,
    # $is_online: Boolean,
    $thumbnail_image_url: String,
    $ticket_price: numeric,
    $title: String,
    $total_available_tickets: Int,
    $location: data_locations_set_input!
    $venue: String,
    $tags: [data_event_tags_insert_input!]!,
    $images: [data_images_insert_input!]!
  ) {
    # Update the event details
    update_data_events_by_pk(
      pk_columns: { event_id: $id }
      _set: {
        category_id: $category_id, 
        description: $description, 
        event_date: $event_date, 
        # # is_online: $is_online, 
        thumbnail_image_url: $thumbnail_image_url, 
        ticket_price: $ticket_price, 
        title: $title, 
        total_available_tickets: $total_available_tickets, 
        venue: $venue,
        by_organization_id: $by_organization_id,
      }
    ) {
      event_id
    }
    update_data_locations(where: {located_event_id: {_eq: $id}}, _set: $location) {
    affected_rows
  }
    # Delete existing event tags
    delete_data_event_tags(where: { tagged_event_id: { _eq: $id } }) {
      affected_rows
    }

    # Insert new event tags
    insert_data_event_tags(objects: $tags) {
      affected_rows
    }

    # Insert new images
    insert_data_images(objects: $images) {
      affected_rows
    }
  }
`
export const BOOK_MARK_EVENT = gql`
  mutation bookMarkEvent(
    $book_marked_event_id: uuid!
    $book_marker_user_id: uuid!
  ) {
    insert_data_bookmarks(objects: {book_marked_event_id: $book_marked_event_id, book_marker_user_id: $book_marker_user_id}) {
      returning {
          book_marked_event_id
          book_marker_user_id
        }
    }
  }
`
export const UN_BOOK_MARK_EVENT = gql`
  mutation unBookMarkEvent(
    $book_marker_user_id: uuid!
    $book_marked_event_id: uuid!
  ) { 
    delete_data_bookmarks(where: { _and: {book_marked_event_id: {_eq: $book_marked_event_id}, book_marker_user_id: {_eq: $book_marker_user_id}}}) {
    affected_rows
  }
  }
`

export const FOLLOW_EVENT = gql`
  mutation followEvent(
    $followed_organization_id: uuid!
    $following_user_id: uuid!
  ) {
    insert_data_follows(objects: {followed_organization_id: $followed_organization_id, following_user_id: $following_user_id}) {
    returning {
      followed_organization_id
      following_user_id
    }
    }
  }
`

export const UN_FOLLOW_EVENT = gql`
  mutation unfolloweEVENT(
    $followed_organization_id: uuid!
    $following_user_id: uuid!
  ) { 
    delete_data_follows(where: { _and: {followed_organization_id: {_eq: $followed_organization_id}, following_user_id: {_eq: $following_user_id}}}) {
    affected_rows
  }
  }
`
export const DELETE_EVENT = gql`
mutation deleteEvent(
  $event_id: uuid!
) {
  delete_data_events(where: {event_id: {_eq: $event_id}}) {
    affected_rows
  }
}
`

export const DELETE_ORG = gql`
mutation deleteOrg($organization_id:uuid!) {
  delete_data_organizations(where: {organization_id: {_eq: $organization_id}}) {
    affected_rows
  }
}
`

export const CHECKOUT_TICKET = gql`
mutation checkoutTicket ($event_id:String!,$user_id:String!){
  ticketcheckout(event_id: $event_id, user_id: $user_id) {
    url
  }
}

`
