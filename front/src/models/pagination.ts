/**
 * Pagination model
 */

// Meta data for pagination
export interface Meta {
  current_page: number;
  last_page: number;
  total: number;
}

// Pagination links
export interface PageLink {
  index: number | string;
  label?: number | string;
  active?: boolean;
}
