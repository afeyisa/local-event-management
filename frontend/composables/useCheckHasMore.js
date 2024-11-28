import { apolloClient } from '~/plugins/apollo'
import { usePagination } from '~/stores/pagination'
import { HAS_MORE_NEAR_BY_EVENT } from '~/graphql/querie/getNearyByevents.graphql'
import { HAS_MORE_EVENT } from '~/graphql/querie/browseEvents.graphql'

export const usehasMoreEventsOnLocation = async (whereClause, lt, lg) => {
  const pagination = usePagination()
  const { data } = await apolloClient.query({
    query: HAS_MORE_NEAR_BY_EVENT,
    variables: { where: whereClause, user_lat: lt, user_lon: lg, offset: pagination.offset },
  })
  pagination.has_more = data.data_get_nearby_events_aggregate.has_more.count !== 0
}

export const usehasMoreEventsOnSearch = async (whereClause) => {
  const pagination = usePagination()
  const { data } = await apolloClient.query({
    query: HAS_MORE_EVENT,
    variables: { where: whereClause, offset: pagination.offset },
  })
  pagination.has_more = data.data_events_aggregate.has_more.count !== 0
}
