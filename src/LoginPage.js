import React, { useState } from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { login } from './api';

import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function LoginPage(props) {
  const [name, setName] = useState('');
  const [loading, setLoading] = useState(false);

  const doLogin = async () => {
    setLoading(true);
    try {
      await login(name);
      props.loginSuccess(name);
    } catch (err) {
      console.error('Failed to log in. ', err);
      toast.error(`Failed to log in. ${err}`);
      setLoading(false);
    }
  }

  return (
    <div style={{ height: '100%', width:  '100%' }}>
      <div style={{
        width: '400px',
        height: '300px',
        backgroundColor: '#C8C8C8',
        position: 'absolute',
        left:0, right:0,
        top:0, bottom:0,
        margin: 'auto',
        maxWidth: '100%',
        maxHeight: '100%',
        overflow: 'auto',
        borderRadius: '4px',
        boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.4), 0 6px 20px 0 rgba(0, 0, 0, 0.3)',
      }}>
        <div style={{
          marginLeft: '30px',
          marginTop: '30px',
          width: '340px',
          height: '240px',
        }}>
          <TextField
            fullWidth
            id="outlined-required"
            label="Username"
            margin="normal"
            variant="filled"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <TextField
            fullWidth
            id="outlined-password-input"
            label="Password"
            type="password"
            margin="normal"
            variant="filled"
          />
          <br />
          <br />
          <Button
            fullWidth
            id="login"
            variant="contained"
            color="primary"
            disabled={loading}
            style={{ height: '56px' }}
            onClick={() => doLogin()}
          >
            Sign In
          </Button>
        </div>
      </div>
      <ToastContainer />
    </div>
  )
}

export default LoginPage;
