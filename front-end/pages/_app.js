import "../styles/navbar.css";
import "../styles/style.css";
import { ToastProvider } from "react-toast-notifications";

function MyApp({ Component, pageProps }) {
  return (
    <ToastProvider autoDismiss={true} placement={"bottom-right"}>
      <Component {...pageProps} />;
    </ToastProvider>
  );
}
export default MyApp;
