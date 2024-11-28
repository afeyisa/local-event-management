import { useMutation } from '@vue/apollo-composable'
import { useToast } from 'vue-toastification/dist/index.mjs'
import { fetchUseId } from './user'
import { CHECKOUT_TICKET } from '~/graphql/mutations/ticketCheckOut.graphql'
import { CHECKOUT_FREE_TICKET } from '~/graphql/mutations/checkOutFreeTicket.graphql'
// add the function that award free ticket or route to ticket payment

export const useTicket = () => {
  const form = ref({
    email: ref(''),
    phone: ref(''),
    lastname: ref(''),
    firstName: ref(''),
  })

  const buyFreeTicket = async (id) => {
    const { myId } = await fetchUseId()
    if (myId.value) {
      try {
        const { mutate } = useMutation(CHECKOUT_FREE_TICKET)
        const { data } = await mutate({
          event_id: id,
          user_id: myId.value,
        })
        if (data) {
          const toast = useToast()
          toast.success(data.ticketcheckoutfree.message)
        }

        else {
          const toast = useToast()
          toast.error('something went wrong try again later!')
        }
      }
      catch (err) {
        const toast = useToast()
        toast.error(err.message)
      }
    }
    else {
      const toast = useToast()
      toast.info('Sign up to buy tickets and enjoy more features')
    }
  }

  const buyTicket = async (id) => {
    const { myId } = await fetchUseId()
    if (myId.value && form.value !== '') {
      try {
        const { mutate } = useMutation(CHECKOUT_TICKET)
        const { data } = await mutate({
          event_id: id,
          user_id: myId.value,
          email: form.value.email,
          first_name: form.value.firstName,
          last_name: form.value.lastname,
          phone_number: form.value.phone,
        })
        if (data) {
          window.location.href = data.ticketcheckoutpaid.url
        }
        else {
          const toast = useToast()
          toast.error('something went wrong try again later!')
        }
      }
      catch (err) {
        console.log(err)
        const toast = useToast()
        toast.error('something went wrong try again later!')
      }
    }
    else {
      const toast = useToast()
      toast.info('Sign up to buy tickets and enjoy more features')
    }
  }
  return { form, buyTicket, buyFreeTicket }
}
