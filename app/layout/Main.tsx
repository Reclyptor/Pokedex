import type { ReactNode } from "react";
import { twMerge } from "tailwind-merge";
import clsx from "clsx";

export type MainProps = {
  header?: ReactNode;
  footer?: ReactNode;
  children?: ReactNode;
  className?: string;
};

const Main = (props: MainProps) => {
  return (
    <div className={ twMerge(clsx("flex flex-col [&>main]:overflow-y-auto", props.className)) }>
      <header className="flex items-center justify-center w-full flex-shrink">
        { props.header }
      </header>
      <main className="flex items-center justify-center w-full flex-grow">
        { props.children }
      </main>
      <footer className="flex items-center justify-center w-full flex-shrink">
        { props.footer }
      </footer>
    </div>
  );
};

export default Main;