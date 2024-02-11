import { createLazyFileRoute } from '@tanstack/react-router';

export const Route = createLazyFileRoute('/blog/$blogId/$postId')({
  component: Page
});

function Page() {
  const { blogId, postId } = Route.useParams()

  return (
    <div>
      <h1>Hello /blog/$blogId/$postId! {blogId} {postId}</h1>
    </div>
  )
}
