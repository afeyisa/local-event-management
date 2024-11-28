import { PAYMENT_AMOUNT } from '~/graphql/querie/getPaymentAmout.graphql'
import { apolloClient } from '~/plugins/apollo'

export const getpaymentAmount = async (id) => {
  const paymentAmount = ref(null)
  const error = ref(null)
  const loading = ref(null)
  try {
    const { data, loading: ld } = await apolloClient.query({
      query: PAYMENT_AMOUNT,
      variables: { tx_ref: id },
    })
    paymentAmount.value = data.data_payments[0]
    loading.value = ld
    return { paymentAmount, error, loading }
  }
  catch (err) {
    console.log(err)
    error.value = err
  }

  return { paymentAmount, error, loading }
}
