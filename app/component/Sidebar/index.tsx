import { theme, sidebar } from "~/component/Sidebar/style.css";
import { twMerge } from "tailwind-merge";
import clsx from "clsx";

export type SidebarProps = {
  open?: boolean;
  className?: string;
};

const Sidebar = (props: SidebarProps) => {
  return !props.open ? null : (
    <div className={ twMerge(clsx(theme, sidebar, props.className)) }>
    </div>
  );
};

export default Sidebar;