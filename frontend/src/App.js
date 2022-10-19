import './App.css';
import { GoogleLogin } from 'react-google-login';
import { gapi } from 'gapi-script';
import {useEffect} from 'react';

function App() {
  const clientId = "321059495197-k1teta6i8jmk2685s8h08jf3hkab1spr.apps.googleusercontent.com"

  useEffect(() => {
    const initClient = () => {
          gapi.auth2.init({
          clientId: clientId,
          scope: ''
        });
      };
      gapi.load('client:auth2', initClient);
  });

    const onSuccess = (res) => {
        console.log('success:', res);
    };
    const onFailure = (err) => {
        console.log('failed:', err);
    };
    return (
       <GoogleLogin
          clientId={clientId}
          buttonText="Sign in with Google"
          onSuccess={onSuccess}
          onFailure={onFailure}
          cookiePolicy={'single_host_origin'}
          isSignedIn={true}
      />
  );
}

export default App;
