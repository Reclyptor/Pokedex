import { Circle as CircleIcon } from "lucide-react";
import { twMerge } from "tailwind-merge";
import clsx from "clsx";

export type SidebarProps = {
  open?: boolean;
  className?: string;
};

const Sidebar = (props: SidebarProps) => {
  return !props.open ? null : (
    <div className={ twMerge(clsx("flex flex-col w-[256px] h-full bg-red", props.className)) }>
      <div className="flex items-center w-full h-[56px] px-6 gap-x-4">
        <div className="w-fit h-fit rounded-full border-3 border-black bg-black/80">
          <div className="w-fit h-fit rounded-full border-2 border-primary bg-black/80">
            <div className="relative w-6 h-6 rounded-[inherit] bg-blue animate-[blink_2s_ease_infinite_0s]">
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
      <div className="relative w-full flex-grow border-r-2 border-r-black box-content">
        <div className="absolute top-1/2 right-0 bottom-0 left-0 bg-[linear-gradient(to_bottom,rgba(0,0,0,0),rgba(0,0,0,.6))] pointer-events-none z-[1]" />
      </div>
    </div>
  );
};

export default Sidebar;