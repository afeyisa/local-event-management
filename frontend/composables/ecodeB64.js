export const encodeB64 = (file) => {
  if (!file) return Promise.resolve(null)

  return new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.readAsDataURL(file)

    reader.onload = () => resolve(reader.result)
    reader.onerror = () => reject(new Error('Failed to read the file.'))
  })
}
