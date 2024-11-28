import { useMutation } from '@vue/apollo-composable'
import { UN_FOLLOW_EVENT } from '~/graphql/mutations/unFollow.graphql'
import { FOLLOW_EVENT } from '~/graphql/mutations/follow.graphql'

export function useFollow(myId, organizationId, follows) {
  const { mutate: follow } = useMutation(FOLLOW_EVENT)
  const { mutate: unfollow } = useMutation(UN_FOLLOW_EVENT)

  const toggleFollow = async () => {
    try {
      if (follows.value.length > 0) {
        await unfollow({ followed_organization_id: follows.value[0].followed_organization_id, following_user_id: follows.value[0].following_user_id })
        follows.value = []
      }
      else {
        const { data } = await follow({ followed_organization_id: organizationId, following_user_id: myId })
        follows.value = data.insert_data_follows.returning
      }
    }
    catch (err) {
      console.error(err)
    }
  }

  return { follows, toggleFollow }
}
