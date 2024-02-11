import { createLazyFileRoute } from '@tanstack/react-router';

export const Route = createLazyFileRoute('/register')({
  component: Page,
});

function Page() {
  return (
    <div>Hello /register!</div>
  );
}
