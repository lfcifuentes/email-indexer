export interface Meta {
  current_page: number;
  last_page: number;
  total: number;
}

export interface PageLink {
  label: number | string;
  active?: boolean;
}
