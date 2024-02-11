import * as React from 'react';
import { createLazyFileRoute } from '@tanstack/react-router';
import { PostPage } from '../components/page/post-page';
import { Post } from '../types/post';

export const Route = createLazyFileRoute('/blog/$blogId/$postId')({
  component: Page
});

function Page() {
  const { blogId, postId } = Route.useParams()
  const [data, setData] = React.useState<Post>();

  const fetchData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/api/v1/posts/${blogId}/${postId}`);
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
    <div>
      <h1>Hello /blog/$blogId/$postId! {blogId} {postId}</h1>
      {data && <PostPage post={data}/>}
    </div>
  )
}
