import { BROWSEEVENTS, PUBLIC_BROWSEEVENTS } from '~/graphql/querie/browseEvents.graphql'
import { apolloClient } from '~/plugins/apollo'
import { fetchUseId } from '~/composables/user'
import { usePagination } from '~/stores/pagination'
import { useEventStore } from '~/stores/events'
import { PUB_GET_NEAR_BY_EVENTS, GET_NEAR_BY_EVENTS } from '~/graphql/querie/getNearyByevents.graphql'

export const useOnSearch = async (whereClause) => {
  const { myId } = await fetchUseId()
  const eventStore = useEventStore()
  const pagination = usePagination()

  const eQuery = myId?.value ? BROWSEEVENTS : PUBLIC_BROWSEEVENTS
  try {
    const { data } = await apolloClient.query({
      query: eQuery,
      variables: { where: whereClause, limit: pagination.limit, offset: pagination.offset },
    })
    eventStore.setEvents(data.data_events)
  }
  catch (err) {
    console.log(err)
    /** */
  }
}

export const defaultEventHomeEvents = async () => {
  useOnSearch({})
}

export const useOnsearchLocation = async (whereClause, lt, lg) => {
  const pagination = usePagination()

  if (pagination.has_more) {
    const { myId } = await fetchUseId()
    const where = whereClause ? whereClause : {}
    const eventStore = useEventStore()
    const eQuery = myId?.value ? GET_NEAR_BY_EVENTS : PUB_GET_NEAR_BY_EVENTS
    try {
      const { data } = await apolloClient.query({
        query: eQuery,
        variables: { where: where, user_lat: lt, user_lon: lg, limit: pagination.limit, offset: pagination.offset },
      })
      eventStore.setEvents(data.data_get_nearby_events)
    }
    catch (err) {
      console.log(err)
    /** */
    }
  }
}
