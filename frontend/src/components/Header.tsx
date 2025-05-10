import { useNavigate } from 'react-router';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import useUser from '../hooks/useUser';

const Header = () => {
  const navigate = useNavigate();
  const { logout } = useUser();
  const token = localStorage.getItem('token');

  const onClickLogin = () => {
    navigate('/login');
  }

  const onClickSignup = () => {
    navigate('/signup');
  }

  const onClickHistory = () => {
    navigate('/moods');
  }

  const onClickLogout = () => {
    logout();
  }

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }} fontWeight={700}>
            Mood Tracker
          </Typography>
          {token ? (
            <>
              <Button color="inherit" onClick={onClickHistory}>Histórico</Button>
              <Button color="inherit" onClick={onClickLogout}>Logout</Button>
            </>
          ) : (
            <>
              <Button color="inherit" onClick={onClickLogin}>Login</Button>
              <Button color="inherit" onClick={onClickSignup}>Cadastre-se</Button>
            </>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default Header;