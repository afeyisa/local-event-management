import { fetchUseId } from './user'
import { useBookMark } from '~/composables/useBookMark'
import { apolloClient } from '~/plugins/apollo'
import { GET_EVENT_DETAILS, PUBLIC_GET_EVENT_DETAILS } from '~/graphql/querie/getEventDetails.graphql'

export const fetcheEventDetails = async (eventId) => {
  const event = ref(null)
  const bookmarks = ref([])

  const loading = ref(null)
  const error = ref(null)
  try {
    const { myId } = await fetchUseId()

    const eQUery = myId.value ? GET_EVENT_DETAILS : PUBLIC_GET_EVENT_DETAILS
    const { data, loading: ld, error: er } = await apolloClient.query({ query: eQUery, variables: { id: eventId } })
    loading.value = ld
    error.value = er
    event.value = data.data_events_by_pk

    if (myId.value) {
      bookmarks.value = event.value.bookmarks
    }

    const updatedImages = [{ image_url: event.value.thumbnail_image_url }, ...(event.value.images || [])]
    event.value = { ...event.value, images: updatedImages }

    const { booker } = await useBookMark(bookmarks, eventId)
    return { bookmarks, booker, loading, error, event }
  }
  catch (error) {
    console.log(error)
    return { bookmarks, loading, error, event }
  }
}
