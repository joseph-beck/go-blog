import { Post } from "../types/post"

type props = {
  post: Post;
}

export const BlogPost: React.FC<props> = ({ post }) => {
  return (
    <>
      {JSON.stringify(post)}
    </>
  );
};
