import type { HTMLAttributes } from "react";
import Card from "~/component/Card";
import Stat from "~/component/Stat";
import Type, { type TypeProps } from "./Type";
import { type Pokemon } from "~/type/pokemon";
import { twMerge } from "tailwind-merge";
import clsx from "clsx";

export type DexEntryProps = {
  pokemon: Pokemon;
} & HTMLAttributes<HTMLDivElement>;

const DexEntry = (props: DexEntryProps) => {
  return (
    <Card key={ props.pokemon.id } className="flex flex-col items-center w-fit h-fit overflow-hidden">
      <div className="flex items-center justify-between w-full h-fit">
        <div className="flex items-center w-fit h-fit">
          <span className="flex items-center w-fit h-full px-2 text-black font-bold bg-red border-b-2 border-b-black">{ `#${ String(props.pokemon.id).padStart(4, "0") }` }</span>
          <div className="relative flex items-center w-fit h-fit">
            <div className="relative w-0 h-0 border-t-[26px] border-t-black border-r-[26px] border-r-transparent" />
            <div className="absolute w-0 h-0 top-0 left-0 border-t-[24px] border-t-red border-r-[24px] border-r-transparent" />
          </div>
        </div>
        <span className="flex items-center w-fit h-fit text-primary font-bold uppercase px-2">{ props.pokemon.name }</span>
      </div>
      <div className="relative flex items-center justify-center w-full h-fit">
        <div className="absolute top-0 left-0 flex flex-col w-fit h-fit py-1 gap-y-[1px]">
          {
            props.pokemon.types.map((type) => (
              <div className={ twMerge(clsx("flex items-center justify-center rounded-r-lg p-1 border-t-2 border-r-2 border-b-2 border-black", {
                "bg-bug": type === "bug",
                "bg-dark": type === "dark",
                "bg-dragon": type === "dragon",
                "bg-electric": type === "electric",
                "bg-fairy": type === "fairy",
                "bg-fighting": type === "fighting",
                "bg-fire": type === "fire",
                "bg-flying": type === "flying",
                "bg-ghost": type === "ghost",
                "bg-grass": type === "grass",
                "bg-ground": type === "ground",
                "bg-ice": type === "ice",
                "bg-normal": type === "normal",
                "bg-poison": type === "poison",
                "bg-psychic": type === "psychic",
                "bg-rock": type === "rock",
                "bg-steel": type === "steel",
                "bg-water": type === "water",
              })) }>
                <Type key={ type } type={ type as TypeProps["type"] } color="var(--color-surface)" />
              </div>
            ))
          }
        </div>
        <img alt={ props.pokemon.name } src={ props.pokemon.sprite } className="w-fit h-fit object-fill" />
      </div>
      <div className="grid grid-cols-6 w-fit h-fit gap-x-3 gap-y-1 pb-1">
        <div className="flex flex-col items-center">
          <Stat stat="hp" size="16px" />
          <span className="text-xs font-bold">{ props.pokemon.stats.hp }</span>
        </div>
        <div className="flex flex-col items-center">
          <Stat stat="speed" size="16px" />
          <span className="text-xs font-bold">{ props.pokemon.stats.speed }</span>
        </div>
        <div className="flex flex-col items-center">
          <Stat stat="attack" size="16px" />
          <span className="text-xs font-bold">{ props.pokemon.stats.attack }</span>
        </div>
        <div className="flex flex-col items-center">
          <Stat stat="defense" size="16px" />
          <span className="text-xs font-bold">{ props.pokemon.stats.defense }</span>
        </div>
        <div className="flex flex-col items-center">
          <Stat stat="spAttack" size="16px" />
          <span className="text-xs font-bold">{ props.pokemon.stats.spAttack }</span>
        </div>
        <div className="flex flex-col items-center">
          <Stat stat="spDefense" size="16px" />
          <span className="text-xs font-bold">{ props.pokemon.stats.spDefense }</span>
        </div>
      </div>
    </Card>
  );
};

export default DexEntry;
