export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const userIp
      = getRequestHeader(event, 'x-forwarded-for')?.split(',')[0]
      || event.node.req.socket.remoteAddress

  if (!userIp) {
    return createError({
      statusCode: 400,
      statusMessage: 'Unable to determine IP address',
    })
  }
  console.log(userIp)
  const apiUrl = `${config.infoUrl}/${userIp}/json?token=${config.ipinfo_token}`

  try {
    const response = await $fetch(apiUrl)
    return response
  }
  catch {
    return createError({
      statusCode: 500,
      statusMessage: 'Failed to fetch IP information',
    })
  }
})
