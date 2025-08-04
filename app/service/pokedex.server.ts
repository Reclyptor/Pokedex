import type { Pokemon } from "~/type/Pokemon";
import Pokedex from 'pokedex-promise-v2';

export const getPokedex = async (): Promise<Pokemon[]> => {
  const pokedex = new Pokedex();
  const { results } = await pokedex.getPokemonsList();
  const entries = results.slice(0, 1);
  const pokemon = await Promise.all(entries.map((entry) => pokedex.getPokemonByName(entry.name)));
  return pokemon.map((pokemon) => ({
    id: pokemon.id,
    name: pokemon.name,
    default: pokemon.is_default,
    stats: {
      hp: pokemon.stats.find(stat => stat.stat.name === 'hp')?.base_stat || 0,
      attack: pokemon.stats.find(stat => stat.stat.name === 'attack')?.base_stat || 0,
      defense: pokemon.stats.find(stat => stat.stat.name === 'defense')?.base_stat || 0,
      spAttack: pokemon.stats.find(stat => stat.stat.name === 'special-attack')?.base_stat || 0,
      spDefense: pokemon.stats.find(stat => stat.stat.name === 'special-defense')?.base_stat || 0,
      speed: pokemon.stats.find(stat => stat.stat.name === 'speed')?.base_stat || 0,
    },
    types: pokemon.types.map((type) => type.type.name),
    sprite: pokemon.sprites.front_default || "",
  }));
};