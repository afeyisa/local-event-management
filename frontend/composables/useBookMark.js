import { useMutation } from '@vue/apollo-composable'
import { useToast } from 'vue-toastification/dist/index.mjs'
import { BOOK_MARK_EVENT } from '~/graphql/mutations/bookMark.graphql'
import { UN_BOOK_MARK_EVENT } from '~/graphql/mutations/unBookMark.graphql'
import { fetchUseId } from '~/composables/user'

export const useBookMark = async (eBookmarks, event_id) => {
  const { mutate: unBook } = useMutation(UN_BOOK_MARK_EVENT)
  const { mutate: book } = useMutation(BOOK_MARK_EVENT)
  const { myId } = await fetchUseId()
  const bookmarks = ref(eBookmarks)

  const booker = async () => {
    if (bookmarks.value.length > 0 && myId.value) {
      try {
        await unBook({ book_marker_user_id: bookmarks.value[0].book_marker_user_id, book_marked_event_id: bookmarks.value[0].book_marked_event_id })
        bookmarks.value = []
      }
      catch {
      /** */
        const toast = useToast()
        toast.error('something went wrong try again!')
      }
    }
    else if (myId.value) {
      try {
        const { data } = await book({ book_marked_event_id: event_id, book_marker_user_id: myId.value })
        bookmarks.value = data.insert_data_bookmarks.returning
      }
      catch {
        const toast = useToast()
        toast.error('something went wrong try again!')
        /** */
      }
    }
    else {
      const toast = useToast()
      toast.info('Sign up to bookmark events and enjoy more features!')
    }
  }

  return { booker, bookmarks }
}
