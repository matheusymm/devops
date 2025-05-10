import { useState } from 'react';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Header from '../components/Header';
import useUser from '../hooks/useUser';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { login } = useUser();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await login(email, password);
    } catch (error) {
      console.error('Login failed:', error);
    }
  };

  return (
    <div className="flex flex-col justify-center bg-slate-300 w-screen h-screen">
      <Header />
      <form className='flex flex-col justify-center items-center w-full h-full'>
        <h1 className="text-4xl font-bold mb-4">Login</h1>
        <TextField
          label="E-mail"
          variant="outlined"
          margin="normal"
          className='w-1/4'
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <TextField
          label="Senha"
          type="password"
          variant="outlined"
          margin="normal"
          className='w-1/4'
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <Button variant="contained" color="primary" type="submit" size='large' onClick={handleSubmit}>
          Entrar
        </Button>
      </form>
    </div>
  );
}

export default Login;