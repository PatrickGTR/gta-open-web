import create from "zustand";

const useStore = create((set) => ({
  loginStatus: false,
  setLoginStatus: (toggle) => set({ loginStatus: toggle }),
}));

export default useStore;
