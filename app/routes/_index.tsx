import { Outlet, useLocation, useNavigate } from "react-router";
import {useEffect, useMemo, useState} from "react";
import Sidebar from "~/component/Sidebar";
import Main from "~/layout/Main";
import Cover from "~/component/Cover";

const Index = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const [coverOpen, setCoverOpen] = useState<boolean>(location.pathname !== "/");
  const [sidebarOpen, setSidebarOpen] = useState<boolean>(location.pathname !== "/");
  const [ready, setReady] = useState<boolean>(location.pathname !== "/");

  const sidebar = useMemo(() => {
    return (
      <Sidebar open={ sidebarOpen } onReady={ () => setReady(true) } className="" />
    );
  }, [sidebarOpen]);

  const header = useMemo(() => {
    return (
      <span className="flex items-center justify-center w-full h-[56px] text-black font-extrabold text-center text-lg py-2 bg-red border-b-2 border-b-black" />
    );
  }, []);

  const footer = useMemo(() => {
    return (
      <span className="flex items-center justify-center w-full text-black text-center text-[10px] py-1 leading-[10px] bg-red border-t-2 border-t-black">This fan site is not affiliated with The Pokémon Company. All Pokémon content and materials are trademarks and copyrights of The Pokémon Company International, Inc.</span>
    );
  }, []);

  useEffect(() => {
    setSidebarOpen(coverOpen);
  }, [coverOpen]);

  useEffect(() => {
    if (location.pathname === "/" && coverOpen && sidebarOpen && ready) {
      navigate("/pokedex", { replace: true });
    }
  }, [location, navigate, coverOpen, sidebarOpen, ready]);

  return (
    <Main sidebar={ sidebar } header={ header } footer={ footer } className="w-screen h-screen">
      <Cover open={ coverOpen } onChange={ setCoverOpen } className="w-full h-full">
        { location.pathname !== "/" && <Outlet /> }
      </Cover>
    </Main>
  );
};

export default Index;