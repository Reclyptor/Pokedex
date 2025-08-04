import { twMerge } from "tailwind-merge";
import { type ReactNode, useEffect, useState } from "react";
import clsx from "clsx";

export type SidebarProps = {
  open?: boolean;
  onReady?(): void;
  children?: ReactNode;
  className?: string;
};

const Sidebar = (props: SidebarProps) => {
  const [open, setOpen] = useState<boolean>(!!props.open);
  const [animating, setAnimating] = useState<boolean>(false);

  useEffect(() => {
    if (!!props.open && !open) {
      setAnimating(true);
      setTimeout(() => {
        setAnimating(false);
        setOpen(true);
        props.onReady?.();
      }, 500); // Match this to animation duration
    }
  }, [props.open]);

  return (
    <div className={ twMerge(clsx("relative flex flex-col w-[256px] h-full bg-surface", props.className)) }>
      <div className="flex items-center w-full h-[54px] px-6 gap-x-4 bg-red">
        <div className="w-fit h-fit rounded-full border-3 border-black bg-black/80">
          <div className="w-fit h-fit rounded-full border-2 border-primary bg-black/80">
            <div className="relative w-8 h-8 rounded-[inherit] bg-blue animate-[blink_2s_ease_infinite_0s]">
              <div className="absolute top-1/8 left-1/8 w-3 h-3 rounded-full bg-white/25" />
            </div>
          </div>
        </div>
        <div className="flex h-full pt-3 gap-x-[2px]">
          <div className="w-fit h-fit rounded-full border-2 border-black bg-black/80">
            <div className="relative w-3 h-3 rounded-[inherit] bg-red animate-[blink_2s_ease_infinite_0s]">
              <div className="absolute top-1/8 left-1/8 w-1 h-1 rounded-full bg-white/25" />
            </div>
          </div>
          <div className="w-fit h-fit rounded-full border-2 border-black bg-black/80">
            <div className="relative w-3 h-3 rounded-[inherit] bg-yellow animate-[blink_2s_ease_infinite_.33s]">
              <div className="absolute top-1/8 left-1/8 w-1 h-1 rounded-full bg-white/25" />
            </div>
          </div>
          <div className="w-fit h-fit rounded-full border-2 border-black bg-black/80">
            <div className="relative w-3 h-3 rounded-[inherit] bg-green animate-[blink_2s_ease_infinite_.66s]">
              <div className="absolute top-1/8 left-1/8 w-1 h-1 rounded-full bg-white/25" />
            </div>
          </div>
        </div>
      </div>
      <svg viewBox="0 -55 256 56" className="absolute top-[54px] z-[1]" version="1.1" width="256" height="56" xmlns="http://www.w3.org/2000/svg">
        <path d="m 0,-55 c 0,18.666667 0,37.333333 0,56 33.333333,0 66.666667,0 100,0 8.47721,-3.9276261 13.80479,-12.414045 20.20874,-18.959045 11.2166,-12.178426 20.76087,-26.225719 33.93188,-36.412049 L 156,-55" fill="#b62a0a" stroke="none" id="path1" />
        <path d="m 156,-54.5 c 33.33333,0 66.66667,0 100,0" stroke="#000000" strokeWidth="2" fill="none" id="path3" />
        <path d="m 0,0.5 c 33.333333,0 66.666667,0 100,0" stroke="#000000" strokeWidth="2" fill="none" id="path5" />
        <path d="m 100,0.5 c 7.5,0 48.5,-55 56,-55" stroke="#000000" strokeWidth="2" fill="none" id="path6" />
      </svg>
      <div className="w-full flex-grow border-r-2 border-r-black">
        { !open ? <div className={ clsx("w-full h-full pt-[56px] bg-red transition-transform duration-500 ease-in-out", { "-translate-x-full": animating } ) } /> :
          <div className="w-full h-full pt-[56px]">
            { props.children }
          </div>
        }
      </div>
    </div>
  );
};

export default Sidebar;