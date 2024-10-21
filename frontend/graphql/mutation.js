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
