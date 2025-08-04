import { type LoaderFunction, type LoaderFunctionArgs, useLoaderData } from "react-router";
import type { Pokemon } from "~/type/Pokemon";
import { getPokedex } from "~/service/pokedex.server";
import DexEntry from "~/component/DexEntry";

type LoaderData = {
  pokemon: Pokemon[];
};

const getPokemon = async (): Promise<Pokemon[]> => {
  return getPokedex();
};

export const loader: LoaderFunction = async (_: LoaderFunctionArgs): Promise<LoaderData> => {
  return { pokemon: await getPokemon() };
};

const Pokedex = () => {
  const { pokemon } = useLoaderData() as LoaderData;

  return (
    <div className="w-full h-full">
      <div className="flex flex-wrap gap-1 p-1 w-full h-fit">
        { pokemon.map((pokemon) => <DexEntry key={ pokemon.id } pokemon={ pokemon }/>) }
      </div>
    </div>
  );
};

export default Pokedex;