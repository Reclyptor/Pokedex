import { Outlet, useNavigate } from "react-router";
import { useEffect, useMemo } from "react";
import Sidebar from "~/component/Sidebar";
import Main from "~/layout/Main";

const Index = () => {
  const navigate = useNavigate();

  useEffect(() => {
    navigate("/pokedex");
  }, []);

  const sidebar = useMemo(() => {
    return (
      <Sidebar open />
    );
  }, []);

  const header = useMemo(() => {
    return (
      <span className="flex items-center justify-center w-full h-[56px] text-neutral font-bold text-center text-lg py-2 bg-red border-b-2 border-b-black">Pokedex</span>
    );
  }, []);

  const footer = useMemo(() => {
    return (
      <span className="flex items-center justify-center w-full text-neutral text-center text-[10px] py-1 leading-[10px] bg-background-dark">This fan site is not affiliated with The Pokémon Company. All Pokémon content and materials are trademarks and copyrights of The Pokémon Company International, Inc.</span>
    );
  }, []);

  return (
    <Main sidebar={ sidebar } header={ header } footer={ footer } className="w-screen h-screen bg-lcd-dark/75">
      <Outlet />
    </Main>
  );
};

export default Index;