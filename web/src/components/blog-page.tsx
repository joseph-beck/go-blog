import { Blog } from "../types/blog";
import { Post } from "../types/post";
import { BlogAuthor } from "./blog-author";
import { BlogPost } from "./blog-post";

type props = {
  blog: Blog;
  posts: Post[];
};

export const BlogPage: React.FC<props> = ({ blog, posts }): JSX.Element => {
  return (
    <>
      <h1>{blog.name}</h1>
      <BlogAuthor author={blog.author}/>
      <p>{blog.description}</p>

      {posts.map((post: Post, index: number) =>
        <div className="my-4" key={index}>
          <BlogPost post={post}/>
        </div>
      )}
    </>
  );
};

