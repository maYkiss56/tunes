export type Review = {
  id: number;
  user: User;
  song: Track;
  body: string;
  is_like: boolean;
  is_valid: boolean;
  updated_at: string;
};

export type Track = {
  id: number;
  title: string;
  full_title: string;
  image_url: string;
  release_date: string;
  like_count: number;
  dislike_count: number;
  rating: number;
  genre: Genre;
  artist: Artist;
  album?: Album;
};

export type Artist = {
  id: number;
  nickname: string;
  bio: string;
  country: string;
};

export type Album = {
  id: number;
  title: string;
  image_url: string;
  artist: Artist;
};

export type Genre = {
  id: number;
  title: string;
  image_url: string;
};

export type User = {
  id: number;
  email: string;
  username: string;
  avatar_url?: string;
  is_banned: boolean;
  role_id: number;
};

export type Reviewer = {
  id: number;
  username: string;
  avatar_url?: string;
  review_count: number;
};
