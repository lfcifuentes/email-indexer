import {toast as toastify, type ToastPosition, type ToastTheme } from 'vue3-toastify';
import { normalizeErrors, type ErrorMessage } from './errors';

// Define interfaces para mejorar la tipificación
interface ToastOptions {
  autoClose: number;
  theme: ToastTheme;
  position: ToastPosition;
}

// Define una interfaz para los datos de los mensajes
interface MessageData {
  data: string | string[];
  time: number;
}

// Define opciones por defecto para los mensajes
export const defaultToastOptions: ToastOptions = {
  autoClose: 3000,
  theme: 'dark',
  position: 'bottom-right',
};

/**
 * Define un objeto con métodos para mostrar mensajes de error, advertencia y éxito.
 * @type {{success(*): void, warning(*): void, getMessageData(*): {data: string, time: number}}}
 */
export const toastr = {
  getMessageData(message: ErrorMessage): MessageData {
    const normalizedMessage = Array.isArray(message) ? message : [message];

    const data = normalizeErrors(
      normalizedMessage,
      'Ocurrió un error inesperado',
      false,
    );

    return {
      data,
      time: data.length < 3 ? 3000 : (1500 * data.length)
    }
  },

  success(message: ErrorMessage) {
    const { data, time } = this.getMessageData(message);
    toastify.success(data, {
      ...defaultToastOptions,
      autoClose: time
    });
  },

  warning(message: ErrorMessage) {
    const { data, time } = this.getMessageData(message);
    toastify.warning(data, {
      ...defaultToastOptions,
      autoClose: time
    });
  },

  error(message: ErrorMessage) {
    const { data, time } = this.getMessageData(message);
    toastify.error(data, {
      ...defaultToastOptions,
      autoClose: time
    });
  }
}
