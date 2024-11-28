export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const query = getQuery(event)
  const address = query.address
  if (!address) {
    throw createError({
      statusCode: 400,
      statusMessage: 'Address is required',
    })
  }

  const apiKey = config.geoapify_key
  const url = config.geoapify + `=${encodeURIComponent(address)}&apiKey=${apiKey}`
  console.log(url)
  try {
    const response = await $fetch(url)
    return response
  }
  catch (error) {
    throw createError({
      statusCode: 500,
      statusMessage: 'Failed to fetch geocoding data',
      data: error,
    })
  }
})
