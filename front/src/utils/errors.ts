// Define tipos para los mensajes de error
export type ErrorMessage = string | string[];

// Define los posibles tipos de respuesta de error
export type ApiErrorResponse = {
  response?: {
    data?: {
      errors?: ErrorMessage;
      message?: string;
    };
  };
  message?: string;
};

export type ErrorResponse = ApiErrorResponse | string | string[];

// Mapa de mensajes de error comunes
const ERROR_MESSAGES: Record<string, string> = {
  'Network Error': 'Error de conexión con el servidor',
  'Request failed with status code 401': 'No autorizado'
};

/**
 * normaliza los errores de una respuesta de la API.
 * @param err
 * @param defaultText
 * @param returnInArray
 * @returns
 */
export const normalizeErrors = (
  err: ErrorResponse,
  defaultText = '',
  returnInArray = false,
): string | string[] => {
  if (typeof err === 'string') {
    return parseErrors(ERROR_MESSAGES[err] || err, returnInArray);
  }

  if (Array.isArray(err)) {
    return parseErrors(err, returnInArray);
  }

  // Handle ApiErrorResponse
  const message = err?.response?.data?.errors ||
                  err?.response?.data?.message ||
                  err?.message ||
                  defaultText;

  const normalizedMessage = ERROR_MESSAGES[message as string] || message;
  return parseErrors(normalizedMessage, returnInArray);
};

/**
 * parsea los errores en un string o un array de strings.
 * @param errors_array
 * @param returnInArray
 * @param separator
 * @returns
 */
export const parseErrors = (
  errors_array: ErrorMessage,
  returnInArray = false,
  separator = '\n'
): string | string[] => {
  const arrayErrors = typeof errors_array === 'string' ? [errors_array] : errors_array;

  const errors = Object.values(arrayErrors || []).map((e) =>
    Array.isArray(e) ? Object.values(e).join(separator) : `${e}`
  );

  return returnInArray ? errors : errors.join(separator);
};
