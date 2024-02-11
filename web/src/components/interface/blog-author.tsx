import { Author } from "../../types/author";

type props = {
  author: Author;
};

export const BlogAuthor: React.FC<props> = ({ author }): JSX.Element => {
  return (
    <>
      <p>{author.username}, {author.name}</p>
      {author.picture && <p>{author.picture}</p>}
    </>
  );
};
