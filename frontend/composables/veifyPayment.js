// where i can write a query to verify the payment

import { VERIFY_PAYMENT } from '~/graphql/querie/verifyPayment.graphql'
import { apolloClient } from '~/plugins/apollo'

export const verifyPayment = async (tx_ref) => {
  const status = ref(null)
  const error = ref(null)
  const loading = ref(null)
  try {
    const { data, loading: ld } = await apolloClient.query({
      query: VERIFY_PAYMENT,
      variables: { tx_ref: tx_ref },
    })
    status.value = data.verifyPayment.status
    loading.value = ld
    return { status, error, loading }
  }
  catch (err) {
    error.value = err
  }

  return { status, error, loading }
}
