import { Blog } from "../../types/blog";
import { Post } from "../../types/post";
import { BlogAuthor } from "../interface/blog-author";
import { BlogPost } from "../interface/blog-post";

type props = {
  blog: Blog;
  posts: Post[];
};

export const BlogPage: React.FC<props> = ({ blog, posts }): JSX.Element => {
  return (
    <>
      <h1>{blog.name}</h1>
      <BlogAuthor author={blog.author} />
      <p>{blog.description}</p>

      <div className="flex flex-col items-center">
        {posts.map((post: Post, index: number) =>
          <div className="w-1/2 my-4" key={index}>
            <BlogPost post={post} />
          </div>
        )}
      </div>
    </>
  );
};
