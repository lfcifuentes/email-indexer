/**
 * Escapes special characters in a string to be used in a regular expression.
 *
 * @param str - The string to escape.
 * @returns The escaped string.
 */
const escapeRegExp = (str: string): string => {
  // Escapa caracteres especiales para usarlos en una expresión regular.
  return str.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
};
/**
 * Highlights occurrences of the search terms in the given text by wrapping them in <mark> tags.
 *
 * @param text - The text in which to highlight search terms.
 * @param search - The search terms to highlight, separated by spaces.
 * @returns The text with highlighted search terms.
 */
export const highlightText = (text: string, search: string): string => {
  if (!search.trim()) {
    return text;
  }
  // Separa por espacios y elimina entradas vacías.
  const words = search.trim().split(/\s+/).filter(word => word.length > 0);
  // Crea un regex que busque cualquiera de las palabras (modo insensible).
  const regex = new RegExp(`(${words.map(escapeRegExp).join('|')})`, 'gi');
  // Envuelve cada palabra encontrada en <mark>
  return text.replace(regex, '<mark>$1</mark>');
};
