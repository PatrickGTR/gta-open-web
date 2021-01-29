import create from "zustand";
import { persist } from "zustand/middleware";

const useStore = create(
  persist(
    (set) => ({
      loginStatus: false,
      jwtToken: "",
      adminLevel: 0,
      setJwtToken: (token) => set({ jwtToken: token }),
      setLoginStatus: (toggle) => set({ loginStatus: toggle }),
      setAdminLevel: (level) => set({ adminLevel: level }),
    }),
    {
      name: "zustand-states",
      getStorage: () => localStorage,
    },
  ),
);

export default useStore;
