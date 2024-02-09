import { Author } from "./author";
import { Blog } from "./blog";

export type Post = {
  id: number;
  title: string;
  body: string;
  likes: number;
  blog: Blog;
  author: Author;
};
