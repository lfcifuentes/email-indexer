/**
 * Parses a given string by removing HTML tags and replacing certain HTML entities with their corresponding characters.
 *
 * @param body - The string to be parsed.
 * @returns The parsed string with HTML tags removed and certain HTML entities replaced.
 */
export const parseBodyToHTML = (body: string): string => {
  return body
    .replace(/<[^>]+>/g, '')
    .replace(/&nbsp;/g, ' ')
    .replace(/&lt;/g, '<')
    .replace(/&gt;/g, '>')
    .replace(/&amp;/g, '&')
    .replace(/\\n/g, '<br />');
}
