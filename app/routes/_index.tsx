import {Outlet, useNavigate} from "react-router";
import {useEffect, useMemo} from "react";
import Main from "~/layout/Main";

const Index = () => {
  const navigate = useNavigate();

  useEffect(() => {
    navigate("/pokedex");
  }, []);

  const header = useMemo(() => {
    return (
      <span className="text-center text-[10px] py-1 w-fit leading-[10px]">PokemonChamber</span>
    );
  }, []);

  const footer = useMemo(() => {
    return (
      <span className="text-center text-[10px] py-1 w-fit leading-[10px]">PokemonChamber is a fan site unaffiliated with The Pokémon Company. All Pokémon content and materials are trademarks and copyrights of The Pokémon Company International, Inc.</span>
    );
  }, []);

  return (
    <Main header={ header } footer={ footer } className="w-screen h-screen [&>header]:border-b [&>header]:border-b-border [&>footer]:border-t [&>footer]:border-t-border">
      <Outlet />
    </Main>
  );
};

export default Index;