import type { HTMLAttributes } from "react";
import { twMerge } from "tailwind-merge";
import clsx from "clsx";

export type CardProps = {
} & HTMLAttributes<HTMLDivElement>;

const Card = ({ ...intrinsic }: CardProps) => {
  return (
    <div { ...intrinsic } className={ twMerge(clsx("bg-surface text-primary min-w-[56px] min-h-[56px] border-2 border-black rounded-xl", intrinsic.className)) } />
  );
};

export default Card;