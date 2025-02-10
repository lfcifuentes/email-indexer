/**
 * Escapa caracteres especiales en una cadena para ser usada en una expresión regular.
 *
 * @param str - La cadena a escapar.
 * @returns La cadena escapada.
 */
const escapeRegExp = (str: string): string => {
  return str.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
};

/**
 * Resalta las ocurrencias de los términos de búsqueda en el texto envolviéndolas en etiquetas <mark>.
 * Ahora se separa el query por espacios y puntos, para resaltar cada parte.
 *
 * @param text - El texto en el que resaltar.
 * @param search - Los términos de búsqueda a resaltar.
 * @returns El texto con los términos resaltados.
 */
export const highlightText = (text: string, search: string): string => {
  if (!search.trim()) {
    return text;
  }
  // Separa por espacios y puntos, eliminando entradas vacías.
  const words = search.trim().split(/[\s.]+/).filter(word => word.length > 0);
  // Crea el regex que busca cualquiera de los tokens (modo insensible).
  const regex = new RegExp(`(${words.map(escapeRegExp).join('|')})`, 'gi');
  // Envuelve cada coincidencia en <mark>
  return text.replace(regex, '<mark>$1</mark>');
};
