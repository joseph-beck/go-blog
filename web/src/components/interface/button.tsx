import { Link } from "@tanstack/react-router";

type props = {
  text: string;
  to?: string;
  onClick?: () => void;
  style: "standard"  | "alternate";
};

export const Button: React.FC<props> = ({ text, onClick, style }): JSX.Element => {
  let className = "";

  if (style === "standard") {
    className = "block bg-yellow-300 px-5 py-3 text-center text-xs font-bold uppercase text-gray-900 transition hover:bg-yellow-400";
  } else {
    className = "block bg-orange-300 px-5 py-3 text-center text-xs font-bold uppercase text-gray-900 transition hover:bg-orange-400";
  }

  return (
    <button
      onClick={onClick}
      className={className}
    >
      {text}
    </button>
  );
};

export const ButtonLink: React.FC<props> = ({ text, to, style }): JSX.Element => {
  let className = "";

  if (style === "standard") {
    className = "block bg-yellow-300 px-5 py-3 text-center text-xs font-bold uppercase text-gray-900 transition hover:bg-yellow-400";
  } else {
    className = "block bg-orange-300 px-5 py-3 text-center text-xs font-bold uppercase text-gray-900 transition hover:bg-orange-400";
  }

  return (
    <Link
      to={to}
      className={className}
    >
      {text}
    </Link>
  );
};
