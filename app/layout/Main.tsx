import type { ReactNode } from "react";
import { twMerge } from "tailwind-merge";
import clsx from "clsx";

export type MainProps = {
  sidebar?: ReactNode;
  header?: ReactNode;
  footer?: ReactNode;
  children?: ReactNode;
  className?: string;
};

const Main = (props: MainProps) => {
  return (
    <div className={ twMerge(clsx("flex", props.className)) }>
      <aside className="flex h-full">
        { props.sidebar }
      </aside>
      <div className="flex flex-col flex-grow">
        <header className="w-full">
          { props.header }
        </header>
        <main className="w-full flex-grow overflow-y-auto">
          { props.children }
        </main>
        <footer className="w-full flex-shrink">
          { props.footer }
        </footer>
      </div>
    </div>
  );
};

export default Main;