import { Author } from "./author";
import { Blog } from "./blog";

export type Post = {
  id: string;
  title: string;
  body: string;
  likes: number;
  blog: Blog;
  author: Author;
};
