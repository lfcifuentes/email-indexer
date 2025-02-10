export interface Meta {
  current_page: number;
  last_page: number;
  total: number;
}

export interface PageLink {
  index: number | string;
  label?: number | string;
  active?: boolean;
}
