import { createLazyFileRoute } from '@tanstack/react-router'
import * as React from 'react'
import { BlogPage } from '../components/BlogPage';
import { Post } from '../types/post';
import { Blog } from '../types/blog';

export const Route = createLazyFileRoute('/blog')({
  component: Page
})

type Response = {
  blog: Blog;
  posts: Post[];
};
function Page() {
  const [data, setData] = React.useState<Response>();

  const fetchData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/api/v1/blogs/1`);
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const jsonData = await response.json();
      setData(jsonData);
      console.log(data)
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  React.useEffect(() => {
    fetchData();
  }, []);

  return (
    <div className="">
      {data && <BlogPage blog={data.blog} posts={data.posts}/>}
    </div>
  );
}
