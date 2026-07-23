export function formatPublishedDate(value: string): string {
  return new Date(value).toLocaleDateString(undefined, { timeZone: 'UTC' })
}
