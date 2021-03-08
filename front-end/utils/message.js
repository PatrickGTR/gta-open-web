import { useToasts } from "react-toast-notifications";

const noAvailableServer = "error fetching data";

const useMessage = () => {
  const { addToast } = useToasts();

  return {
    notifyError: (msg) => addToast(msg, { appearance: "error" }),
    notifySuccess: (msg) => addToast(msg, { appearance: "success" }),
    notifyInfo: (msg) => addToast(msg, { appearance: "info" }),
    notifyWarning: (msg) => addToast(msg, { appearance: "warning" }),
  };
};

export { noAvailableServer, useMessage };
