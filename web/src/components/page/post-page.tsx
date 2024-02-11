import { Post } from "../../types/post";
import { BlogPost } from "../interface/blog-post";

type props = {
  post: Post;
}

export const PostPage: React.FC<props> = ({ post }): JSX.Element => {
  return (
    <div className="flex flex-col items-center">
      <div className="w-1/2 my-4">
        <BlogPost post={post} />
      </div>
    </div>
  );
};
