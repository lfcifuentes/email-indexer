import type { AxiosInstance } from "axios";
import axios from "axios";

/**
 * Trims the trailing slash from a URL if it exists.
 *
 * @param url - The URL to be trimmed.
 * @returns The trimmed URL.
 */
const trimUrl = (url: string): string => {
  if (url.substr(-1) === "/") {
    url = url.substr(0, url.length - 1);
  }
  return url;
}

/**
 * Options for configuring the API instance.
 */
interface ApiOptions {
  /**
   * Custom headers to be included in the API requests.
   */
  customHeaders?: any;
}

/**
 * Creates an Axios instance configured with the specified options.
 *
 * @param options - Configuration options for the API instance.
 * @returns An Axios instance configured with the specified options.
 */
export const getApi = (options: ApiOptions = {}): AxiosInstance => {
  const {
    customHeaders = {},
  } = options;

  let baseURL = import.meta.env.VITE_BASE_URL || "";

  baseURL = trimUrl(baseURL);

  const instance = axios.create({
    baseURL: baseURL,
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json",
      ...customHeaders,
    }
  });

  return instance;
}
