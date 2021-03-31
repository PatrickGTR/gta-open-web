import create from "zustand";
import { persist } from "zustand/middleware";

const useStore = create(
  persist(
    (set) => ({
      loginStatus: false,
      setLoginStatus: (toggle) => set({ loginStatus: toggle }),
    }),
    {
      name: "storage",
      getStorage: () => localStorage,
    },
  ),
);

export default useStore;
