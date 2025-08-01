import { Circle as CircleIcon } from "lucide-react";
import { twMerge } from "tailwind-merge";
import clsx from "clsx";
import type {ReactNode} from "react";

export type SidebarProps = {
  open?: boolean;
  children?: ReactNode;
  className?: string;
};

const Sidebar = (props: SidebarProps) => {
  return (
    <div className={ twMerge(clsx("relative flex flex-col w-[256px] h-full", props.className)) }>
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
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 -55 256 56" className="absolute top-[54px]">
        {/* Red fill with no stroke */}
        <path d="m 0 -54 L 156 -54 L 256 -54 L 256 -54 L 0 -54 L 0 0 L 100 0 C 108 0 148 -54 156 -54" fill="#B62A0A" stroke="none" />
        {/* Top-left border (from 0,-54 to 156,-54) */}
        <path d="m 0 -54 L 156 -54" stroke="#B62A0A" strokeWidth="2" fill="none" />
        {/* Top-right border (from 156,-54 to 256,-54) */}
        <path d="m 156 -54 L 256 -54" stroke="#000000" strokeWidth="2" fill="none" />
        {/* Left border (from 0,-54 to 0,0) */}
        <path d="m 0 -54 L 0 0" stroke="#B62A0A" strokeWidth="2" fill="none" />
        {/* Bottom-left border (from 0,0 to 100,0) */}
        <path d="m 0 0 L 100 0" stroke="#000000" strokeWidth="2" fill="none" />
        {/* Curved border (from 100,0 to 156,-54) */}
        <path d="m 100 0 C 108 0 148 -54 156 -54" stroke="#000000" strokeWidth="2" fill="none" />
      </svg>
      <div className="w-full flex-grow border-r-2 border-r-black pt-[56px]">
        { props.children }
      </div>
    </div>
  );
};

export default Sidebar;