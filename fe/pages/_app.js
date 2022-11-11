import '../styles/globals.css'
import { GlobalContextProvider } from '/contexts/globalContext';

function MyApp({ Component, pageProps }) {
  return (
    <>
    <GlobalContextProvider>
      <Component {...pageProps} />
    </GlobalContextProvider>
  </>
  );
}

export default MyApp
