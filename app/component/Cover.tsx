import type { ReactNode } from "react";
import { twMerge } from "tailwind-merge";
import clsx from "clsx";

export type CoverProps = {
  open?: boolean;
  children?: ReactNode;
  className?: string;
};

const Cover = (props: CoverProps) => {
  return !props.open ? props.children : (
    <div className={ twMerge(clsx("flex flex-col justify-end bg-red", props.className)) }>
      <div className="relative m-12">
        <div className="absolute w-0 h-0 top-[-4px] left-[-2px] border-t-[54px] border-t-transparent border-l-[66px] border-l-black  border-b-[54px] border-b-transparent" />
        <div className="relative w-0 h-0 border-t-[50px] border-t-transparent border-l-[60px] border-l-yellow border-b-[50px] border-b-transparent animate-[blink_2s_ease_infinite_0s]" />
      </div>
      <div className="flex items-end justify-center h-[calc(50%-98px)] p-28">
        <div className="w-full h-[32px] border-2 border-black rounded-3xl" />
      </div>
    </div>
  );
};

export default Cover;
