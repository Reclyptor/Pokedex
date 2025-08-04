export type Pokemon = {
  id: number;
  name: string;
  default: boolean;
  stats: {
    hp: number;
    attack: number;
    defense: number;
    spAttack: number;
    spDefense: number;
    speed: number;
  },
  types: string[],
  sprite: string;
};