import { createLazyFileRoute } from '@tanstack/react-router';

export const Route = createLazyFileRoute('/login')({
  component: Page
});

function Page() {
  return (
    <div>Hello /login!</div>
  );
}
