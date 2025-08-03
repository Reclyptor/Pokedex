import type { HTMLAttributes } from "react";
import { twMerge } from "tailwind-merge";
import clsx from "clsx";

export type CardProps = {
} & HTMLAttributes<HTMLDivElement>;

const Card = ({ ...intrinsic }: CardProps) => {
  return (
    // <div { ...intrinsic } className={ twMerge(clsx("bg-surface/95 bg-white box-border min-w-[256px] min-h-[56px]", intrinsic.className)) }>
      <div className="relative w-[256px] h-[56px] box-border border-2 border-green">
        <svg viewBox="0 -55 256 56" version="1.1" width="256" height="56" className="absolute -top-[2px] -left-[2px] -z-[1]" xmlns="http://www.w3.org/2000/svg">
          <path d="m 0,-55 c 0,18.666667 0,37.333333 0,56 33.333333,0 66.666667,0 100,0 8.47721,-3.9276261 13.80479,-12.414045 20.20874,-18.959045 11.2166,-12.178426 20.76087,-26.225719 33.93188,-36.412049 L 156,-55" fill="#b62a0a" stroke="none" id="path1" />
          <path d="m 156,-54.5 c 33.33333,0 66.66667,0 100,0" stroke="#000000" strokeWidth="2" fill="none" id="path3" />
          <path d="m 0,0.5 c 33.333333,0 66.666667,0 100,0" stroke="#000000" strokeWidth="2" fill="none" id="path5" />
          <path d="m 100,0.5 c 8.43916,-3.8274626 13.80253,-12.183403 20.20874,-18.620491 11.21615,-11.963879 20.79497,-25.790649 33.93188,-35.761833 L 156,-54.5" stroke="#000000" strokeWidth="2" fill="none" id="path6" />
        </svg>
      </div>
    // </div>
  );
};

export default Card;