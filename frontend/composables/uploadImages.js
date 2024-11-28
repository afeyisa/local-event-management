import { useToast } from 'vue-toastification/dist/index.mjs'
import { useMutation } from '@vue/apollo-composable'

import { SAVEIMAGES } from '~/graphql/mutations/saveImages.graphql'

export const uploadImages = async (t, o) => {
  const toast = useToast()

  if (!t && (!o || !o.length > 0)) {
    return null
  }
  try {
    const { mutate } = useMutation(SAVEIMAGES)
    const { data } = await mutate(
      { thumbnail: t, otherImages: o },
    )
    const er = data.saveImages.errors
    if (er && er.length > 0) {
      er.forEach((element) => {
        toast.error(element)
      })
    }

    return {
      thumbnailImageUrl: data.saveImages.thumbnail_image_url,
      otherImageUrls: data.saveImages.other_images_urls,
    }
  }
  catch {
    toast.error('some thing went wrong on saving image')
    return null
  }
}
