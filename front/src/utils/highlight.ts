/**
 * Escapes special characters in a string to be used in a regular expression.
 *
 * @param str - The string to escape.
 * @returns The escaped string.
 */
const escapeRegExp = (str: string): string => {
  return str.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
};

/**
 * Highlights search term occurrences in the text by wrapping them in <mark> tags.
 * The query is split by spaces and dots to highlight each part separately.
 *
 * @param text - The text to highlight.
 * @param search - The search terms to highlight.
 * @returns The text with highlighted terms.
 */
export const highlightText = (text: string, search: string): string => {
  if (!search.trim()) {
    return text;
  }
  // Split by spaces and dots, removing empty entries
  const words = search.trim().split(/[\s.]+/).filter(word => word.length > 0);
  // Create regex that matches any of the tokens (case insensitive)
  const regex = new RegExp(`(${words.map(escapeRegExp).join('|')})`, 'gi');
  // Wrap each match in <mark>
  return text.replace(regex, '<mark>$1</mark>');
};
