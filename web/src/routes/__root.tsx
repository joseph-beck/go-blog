import {
  createRootRoute,
  Outlet
} from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import { Navigator } from '../components/interface/navigator';

export const Route = createRootRoute({
  component: Page
});

function Page() {
  return (
    <>
      <Navigator />
      <hr />
      <Outlet />
      <TanStackRouterDevtools />
    </>
  );
}
