import { Link } from "@tanstack/react-router";
import { ButtonLink } from "./button";

type Element = {
  name: string;
  to: string;
}

export const Navigator = (): JSX.Element => {
  const elements: Element[] = [
    {
      name: "home",
      to: "/"
    },
    {
      name: "about",
      to: "/about"
    },
    {
      name: "blog",
      to: "/blog"
    }
  ];

  return (
    <header className="bg-white">
      <div className="mx-auto max-w-screen-xl px-4 sm:px-6 lg:px-8">
        <div className="flex h-16 items-center justify-between">
          <div className="hidden md:block">
            <nav aria-label="Global">
              <ul className="flex items-center gap-6 text-sm">
                {elements.map((element: Element, index: number) =>
                  <li key={index}>
                    <Link
                      className="text-gray-500 transition hover:text-gray-500/75 [&.active]:font-bold"
                      to={element.to}
                    >
                      {element.name}
                    </Link>
                  </li>
                )}
              </ul>
            </nav>
          </div>

          <div className="flex items-center gap-4">
            <div className="sm:flex sm:gap-4">
              <ButtonLink
                text="login"
                to="/login"
                style="standard"
              />

              <div className="hidden sm:flex">
                <ButtonLink
                  text="register"
                  to="/register"
                  style="alternate"
                />
              </div>
            </div>

            <div className="block md:hidden">
              <button className="rounded bg-gray-100 p-2 text-gray-600 transition hover:text-gray-600/75">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  strokeWidth="2"
                >
                  <path strokeLinecap="round" strokeLinejoin="round" d="M4 6h16M4 12h16M4 18h16" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </header>
  )
};
