export const graphqlUrl = () => {
  const config = useRuntimeConfig()
  return config.public.graphqlUrl
}
