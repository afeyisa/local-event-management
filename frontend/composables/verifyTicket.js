import { GET_TICKET_INFO } from '~/graphql/querie/getTicketInfo.graphql'
import { apolloClient } from '~/plugins/apollo'

export const getTicketInfo = async (id) => {
  const ticket = ref(null)
  const error = ref(null)
  const loading = ref(null)
  try {
    const { data, loading: ld } = await apolloClient.query({
      query: GET_TICKET_INFO,
      variables: { ticket_id: id },
    })
    ticket.value = data.data_tickets[0]
    loading.value = ld
    return { ticket, error, loading }
  }
  catch (err) {
    console.log(err)
    error.value = err
  }

  return { ticket, error, loading }
}
